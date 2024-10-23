package handlers

import (
	"github/skrebnevf/comedy_belgrade_bot/pkg/database"
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

func CommandHandlers(b *telebot.Bot, db *supabase.Client) {

	b.Handle("/start", func(c telebot.Context) error {
		if err := database.WriteUser(c, db); err != nil {
			log.Printf("cannot write data from /addme: %v", err)
		}
		return c.Send(Start)
	})

	b.Handle("/events", func(c telebot.Context) error {
		events, err := database.GetEvents(db)
		if err != nil {
			log.Printf("cannot get events from '/events': %v", err)
		}
		return c.Send(events)
	})

	b.Handle("/orgy", func(c telebot.Context) error {
		WaitingForAdminMessage[c.Message().Sender.ID] = true
		return c.Send(OrgyMsg)
	})

	b.Handle("/addme", func(c telebot.Context) error {
		if err := database.WriteUser(c, db); err != nil {
			log.Printf("cannot write data from /addme: %v", err)
		}
		WaitingForMessage[c.Message().Sender.ID] = true
		return c.Send(AddMeFormMsg)
	})

	b.Handle("/human", func(c telebot.Context) error {
		AwaitingForward = true
		OriginalUserID = c.Sender().ID
		return c.Send(ReplyToHumanMsg)
	})

	b.Handle("/lenochka", func(c telebot.Context) error {
		file, err := os.Open(Output)
		if err != nil {
			log.Printf("cannot open file: %v", err)
			return c.Send(CannotOpenFileErrMsg)
		}
		defer file.Close()

		doc := &telebot.Document{
			File:     telebot.FromReader(file),
			FileName: Output,
		}

		if err := c.Send(doc); err != nil {
			log.Printf("cannot sent file: %v", err)
			return c.Send(EmptyFileErrMsg)
		} else {
			return c.Send(SentFileMsg)
		}
	})

	b.Handle("/ochko", func(c telebot.Context) error {
		file, err := os.OpenFile(Output, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("cannot open file: %v", err)
			return c.Send(c.Message().Sender, CannotOpenFileErrMsg)
		}
		defer file.Close()

		if _, err := file.WriteString(""); err != nil {
			log.Printf("cannot write file: %v", err)
			return c.Send(c.Message().Sender, CannotClearFileMsg)
		}

		return c.Send("Записи удалены мой повелитель")
	})
}

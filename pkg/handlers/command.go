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
		database.WriteUser(c, db)
		return c.Send(Start)
	})

	b.Handle("/events", func(c telebot.Context) error {
		events := database.GetEvents(db)
		return c.Send(events)
	})

	b.Handle("/orgy", func(c telebot.Context) error {
		WaitingForAdminMessage[c.Message().Sender.ID] = true
		return c.Send(OrgyMsg)
	})

	b.Handle("/addme", func(c telebot.Context) error {
		database.WriteUser(c, db)
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
			b.Send(c.Message().Sender, CannotOpenFileErrMsg)
			log.Printf("cannot open file: %v", err)
		}
		defer file.Close()

		doc := &telebot.Document{
			File:     telebot.FromReader(file),
			FileName: Output,
		}

		if _, err := b.Send(c.Message().Sender, doc); err != nil {
			log.Printf("cannot sent file: %v", err)
			b.Send(c.Message().Sender, EmptyFileErrMsg)
		} else {
			b.Send(c.Message().Sender, SentFileMsg)
		}

		return c.Send(RazumMsg)
	})

	b.Handle("/ochko", func(c telebot.Context) error {
		file, err := os.OpenFile(Output, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			b.Send(c.Message().Sender, CannotOpenFileErrMsg)
			log.Printf("cannot open file: %v", err)
		}
		defer file.Close()

		if _, err := file.WriteString(""); err != nil {
			c.Send(c.Message().Sender, CannotClearFileMsg)
			log.Printf("cannot write file: %v", err)
		}

		return c.Send("Записи удалены мой повелитель")
	})
}

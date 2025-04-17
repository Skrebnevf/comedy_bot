package handlers

import (
	"fmt"
	"github/skrebnevf/comedy_belgrade_bot/pkg/database"
	"log"
	"os"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

func CommandHandlers(b *telebot.Bot, db *supabase.Client, botUrl string) {
	b.Handle("/start", func(c telebot.Context) error {
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false
		AwaitingForward[c.Message().Sender.ID] = false
		if c.Message().Chat.ID == ChatID {
			return nil
		}
		if err := database.WriteUser(c, db); err != nil {
			log.Printf("cannot write data from /start: %v", err)
		}
		if err := database.AddCommandCounter(c, db); err != nil {
			log.Printf("cannot write command and set counter from /start: %v", err)
		}
		return c.Send(Start)
	})

	b.Handle("/events", func(c telebot.Context) error {
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false
		AwaitingForward[c.Message().Sender.ID] = false
		if c.Message().Chat.ID == ChatID {
			return nil
		}
		events, err := database.GetEvents(db)
		if err != nil {
			log.Printf("cannot get events from '/events': %v", err)
		}
		if err := database.AddCommandCounter(c, db); err != nil {
			log.Printf("cannot write command and set counter from /addme: %v", err)
		}
		return c.Send(events)
	})

	b.Handle("/orgy", func(c telebot.Context) error {
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false
		AwaitingForward[c.Message().Sender.ID] = false
		WaitingForAdminMessage[c.Message().Sender.ID] = true
		return c.Send(OrgyMsg)
	})

	b.Handle("/addme", func(c telebot.Context) error {
		if c.Message().Chat.ID == ChatID {
			return nil
		}

		WaitingForCancel[c.Message().Sender.ID] = false
		AwaitingForward[c.Message().Sender.ID] = false
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = true
		if err := database.WriteUser(c, db); err != nil {
			log.Printf("cannot write data from /addme: %v", err)
		}
		if err := database.AddCommandCounter(c, db); err != nil {
			log.Printf("cannot write command and set counter from /addme: %v", err)
		}
		return c.Send(AddMeFormMsg)
	})

	b.Handle("/cancel", func(c telebot.Context) error {
		if c.Message().Chat.ID == ChatID {
			return nil
		}

		AwaitingForward[c.Message().Sender.ID] = false
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = true
		if err := database.AddCommandCounter(c, db); err != nil {
			log.Printf("cannot write command and set counter from /cancel: %v", err)
		}
		return c.Send(CancelMeMsg)
	})

	b.Handle("/human", func(c telebot.Context) error {
		if c.Message().Chat.ID == ChatID {
			return nil
		}

		AwaitingForward[c.Message().Sender.ID] = true
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false
		OriginalUserID = c.Sender().ID
		if err := database.AddCommandCounter(c, db); err != nil {
			log.Printf("cannot write command and set counter from /addme: %v", err)
		}
		return c.Send(ReplyToHumanMsg)
	})

	b.Handle("/lenochka", func(c telebot.Context) error {
		reservation, err := database.GetReservations(db)
		cancelation, err := database.GetCancelReservations(db)
		if err != nil {
			log.Printf("cannot get reservations lis, error: %v", err)
		}

		AwaitingForward[c.Message().Sender.ID] = false
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false

		file, err := os.OpenFile(Output, os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			log.Printf("cannot open file: %v", err)
			WaitingForMessage[c.Message().Sender.ID] = false
			c.Send(CannotWriteFileMsg)
		}
		defer file.Close()

		if _, err := file.WriteString(""); err != nil {
			log.Printf("cannot write file: %v", err)
			WaitingForMessage[c.Message().Sender.ID] = false
			c.Send(CannotWriteFileMsg)
		}

		if _, err := file.WriteString(fmt.Sprintf("Резервации: %s \n", reservation+"\n")); err != nil {
			log.Printf("cannot write file: %v", err)
			WaitingForMessage[c.Message().Sender.ID] = false
			c.Send(CannotWriteFileMsg)
		}

		if _, err := file.WriteString(fmt.Sprintf("Отмены: %s \n", cancelation)); err != nil {
			log.Printf("cannot write file: %v", err)
			WaitingForMessage[c.Message().Sender.ID] = false
			c.Send(CannotWriteFileMsg)
		}

		output, err := os.Open(Output)
		if err != nil {
			log.Printf("cannot open file: %v", err)
			return c.Send(CannotOpenFileErrMsg)
		}
		defer output.Close()

		doc := &telebot.Document{
			File:     telebot.FromReader(output),
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
		AwaitingForward[c.Message().Sender.ID] = false
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false
		database.AddReservations(c, db, "")
		database.CancelReservation(c, db, "")
		return c.Send("Записи удалены мой повелитель")
	})

	b.Handle("/ebanko", func(c telebot.Context) error {
		if c.Message().Chat.ID == ChatID {
			return nil
		}
		WaitingForAdminMessage[c.Message().Sender.ID] = false
		WaitingForMessage[c.Message().Sender.ID] = false
		WaitingForCancel[c.Message().Sender.ID] = false
		AwaitingForward[c.Message().Sender.ID] = false
		AwaitingSpamMessage[c.Message().Sender.ID] = true
		return c.Send("В следующем сообщении будет анонс мероприятия")
	})
}

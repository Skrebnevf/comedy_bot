package handlers

import (
	"log"
	"os"

	"gopkg.in/telebot.v4"
)

func CommandHandlers(b *telebot.Bot) {
	b.Handle("/start", func(c telebot.Context) error {
		return c.Send(Start)
	})

	b.Handle("/addme", func(c telebot.Context) error {
		WaitingForMessage[c.Message().Sender.ID] = true
		return c.Send(AddMeFormMsg)
	})

	b.Handle("/human", func(c telebot.Context) error {
		AwaitingForward = true
		OriginalUserID = c.Sender().ID // Сохраняем ID пользователя
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

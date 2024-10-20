package handlers

import (
	"log"
	"os"
	"strings"

	"gopkg.in/telebot.v4"
)

func TextHandler(b *telebot.Bot) {
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			if strings.Contains(c.Message().Text, strings.TrimSpace(AdminHelper)) {
				return c.Send(AdminCommandMsg)
			}
			return nil
		}

		if AwaitingForward {
			msg := c.Message()
			log.Println(c.Message().Sender.Username + " Спросил - " + msg.Text)

			var err error
			ForwardedMsg, err = b.Forward(&telebot.Chat{ID: ChatID}, msg)
			if err != nil {
				return err
			}

			AwaitingForward = false
			return c.Send(ReplyedToHumanMsg)
		}

		if WaitingForMessage[c.Message().Sender.ID] {
			text := strings.TrimPrefix(c.Message().Text, "/addme")
			text = strings.TrimSpace(text)

			log.Println(c.Message().Sender.Username + " Записался - " + text)

			file, err := os.OpenFile(Output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				c.Send(CannotWriteFileMsg)
				log.Fatalf("cannot open file: %v", err)
			}
			defer file.Close()

			if _, err := file.WriteString(text + "\n"); err != nil {
				c.Send(CannotWriteFileMsg)
				log.Fatalf("cannot write file: %v", err)
			}

			WaitingForMessage[c.Message().Sender.ID] = false

			return c.Send(AddMeCompleteMsg)
		}

		return c.Send(BaseMsg)
	})
}

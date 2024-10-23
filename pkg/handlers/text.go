package handlers

import (
	"github/skrebnevf/comedy_belgrade_bot/pkg/database"
	"log"
	"os"
	"strings"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

func TextHandler(b *telebot.Bot, db *supabase.Client) {
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
				log.Printf("cannot forwared message: %v", err)
				AwaitingForward = false
				return c.Send(CannotForvaredMsg)
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
				log.Printf("cannot open file: %v", err)
				WaitingForMessage[c.Message().Sender.ID] = false
				c.Send(CannotWriteFileMsg)
			}
			defer file.Close()

			if _, err := file.WriteString(text + "\n"); err != nil {
				log.Printf("cannot write file: %v", err)
				WaitingForMessage[c.Message().Sender.ID] = false
				c.Send(CannotWriteFileMsg)
			}

			WaitingForMessage[c.Message().Sender.ID] = false

			return c.Send(AddMeCompleteMsg)
		}

		if WaitingForAdminMessage[c.Message().Sender.ID] {
			text := strings.TrimPrefix(c.Message().Text, "/ordgy")
			text = strings.TrimSpace(text)

			err := database.AddEvent(c, db, text)
			if err != nil {
				log.Printf("cannot add event: %v", err)
				WaitingForMessage[c.Message().Sender.ID] = false
				return c.Send(CannotAddEventMsg)
			}

			WaitingForMessage[c.Message().Sender.ID] = false

			return c.Send("Объявление для набора в оргию записано")
		}

		return c.Send(BaseMsg)
	})
}

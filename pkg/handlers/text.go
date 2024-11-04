package handlers

import (
	"github/skrebnevf/comedy_belgrade_bot/pkg/database"
	"log"
	"strings"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

func TextHandler(b *telebot.Bot, db *supabase.Client) {
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if strings.Contains(c.Message().Text, AdminHelper) {
			return c.Send(AdminCommandMsg)
		}

		if AwaitingForward {
			msg := c.Message()
			log.Println(c.Message().Sender.Username + " Asked - " + msg.Text)

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
			reservations, err := database.GetReservations(db)
			if err != nil {
				log.Println("cannot get list of reservations, error: %v", err)
			}

			text := strings.TrimPrefix(c.Message().Text, "/addme")
			text = strings.TrimSpace(text)

			log.Println(c.Message().Sender.Username + " Add reservations - " + text)

			msg := reservations + "\n" + text

			if err := database.AddReservations(c, db, msg); err != nil {
				log.Println("cannot write new reservations, error: %v", err)
			}

			WaitingForMessage[c.Message().Sender.ID] = false
			return c.Send(AddMeCompleteMsg)
		}

		if WaitingForAdminMessage[c.Message().Sender.ID] {
			text := strings.TrimPrefix(c.Message().Text, "/orgy")
			text = strings.TrimSpace(text)

			err := database.AddEvent(c, db, text)
			if err != nil {
				log.Printf("cannot add event: %v", err)
				WaitingForAdminMessage[c.Message().Sender.ID] = false
				return c.Send(CannotAddEventMsg)
			}

			WaitingForAdminMessage[c.Message().Sender.ID] = false
			return c.Send("Объявление для набора в оргию записано")
		}

		return c.Send(BaseMsg)
	})
}

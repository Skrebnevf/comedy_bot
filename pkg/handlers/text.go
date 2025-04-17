package handlers

import (
	"github/skrebnevf/comedy_belgrade_bot/pkg/database"
	"log"
	"strings"
	"time"

	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

func TextHandler(b *telebot.Bot, db *supabase.Client, botUrl string) {
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if err := database.WriteMessageLog(c, db); err != nil {
			log.Printf("cannot write message log: %v", err)

		}

		if strings.Contains(c.Message().Text, AdminHelper) {
			return c.Send(AdminCommandMsg)
		}

		if AwaitingForward[c.Message().Sender.ID] {
			msg := c.Message()
			log.Println(c.Message().Sender.Username + " Asked - " + msg.Text)

			var err error
			ForwardedMsg, err = b.Forward(&telebot.Chat{ID: ChatID}, msg)
			if err != nil {
				log.Printf("cannot forwared message: %v", err)
				AwaitingForward[c.Message().Sender.ID] = false
				return c.Send(CannotForvaredMsg)
			}

			AwaitingForward[c.Message().Sender.ID] = false
			return c.Send(ReplyedToHumanMsg)
		}

		if WaitingForMessage[c.Message().Sender.ID] {
			reservations, err := database.GetReservations(db)
			if err != nil {
				log.Printf("cannot get list of reservations, error: %v", err)
			}

			text := strings.TrimPrefix(c.Message().Text, "/addme")
			text = strings.TrimSpace(text)

			log.Println(c.Message().Sender.Username + " Add reservations - " + text)

			msg := reservations + "\n" + text

			if err := database.AddReservations(c, db, msg); err != nil {
				log.Printf("cannot write new reservations, error: %v", err)
			}

			WaitingForMessage[c.Message().Sender.ID] = false
			return c.Send(AddMeCompleteMsg)
		}

		if WaitingForCancel[c.Message().Sender.ID] {
			cancel, err := database.GetCancelReservations(db)
			if err != nil {
				log.Printf("cannot get list of cancel reservation, error: %v", err)
			}

			text := strings.TrimPrefix(c.Message().Text, "/cancel")
			text = strings.TrimSpace(text)

			log.Println(c.Message().Sender.Username + " cancel reserved - " + text)

			msg := cancel + "\n" + text

			if err := database.CancelReservation(c, db, msg); err != nil {
				log.Printf("cannot write cancelation, error: %v", err)
			}

			WaitingForCancel[c.Message().Sender.ID] = false
			return c.Send(CancelReservationMsg)
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

		if AwaitingSpamMessage[c.Message().Sender.ID] {
			msg := c.Message().Text
			AwaitingSpamMessage[c.Message().Sender.ID] = false

			users, err := database.GetUserIDs(db)
			if err != nil {
				log.Println(err)
				return c.Reply("Sorry DB have error")
			}

			for _, user := range users {
				if c.Message().Sender.ID == user.ID {
					continue
				}

				_, err = c.Bot().Send(&telebot.Chat{ID: user.ID}, msg)
				if err != nil {
					log.Println(err)
					return c.Reply("Ошибка в отправке сообщения для %d, %v", user.ID, err)
				}

				time.Sleep(100 * time.Millisecond)
			}

			return c.Send("Заспамили")
		}

		return c.Send(BaseMsg)
	})
}

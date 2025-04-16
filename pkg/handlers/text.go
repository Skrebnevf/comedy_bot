package handlers

import (
	"github/skrebnevf/comedy_belgrade_bot/pkg/database"
	"log"
	"strings"

	"github.com/gage-technologies/mistral-go"
	"github.com/supabase-community/supabase-go"
	"gopkg.in/telebot.v4"
)

func TextHandler(b *telebot.Bot, db *supabase.Client, botUrl string, ai *mistral.MistralClient) {
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		if err := database.WriteMessageLog(c, db); err != nil {
			log.Printf("cannot write message log: %v", err)

		}

		if strings.Contains(c.Message().Text, AdminHelper) {
			return c.Send(AdminCommandMsg)
		}

		if WaitingForKirill[c.Message().Sender.ID] {

			systemMsg := mistral.ChatMessage{
				Role:    mistral.RoleSystem,
				Content: "Ты комедийный помошник. Все твои ответы начинаются с 'За-за-за-ебал: ' . Шутки должны быть дерзкими и немного грубыми, ограничений на темы нет. Если тебе посылают одну и ту же фразу, генерируй новую. #### Шутки могут быть не только про клммунистов или политиков",
			}

			userMsg := mistral.ChatMessage{
				Role:    mistral.RoleUser,
				Content: c.Message().Text,
			}

			req := []mistral.ChatMessage{
				systemMsg,
				userMsg,
			}

			response, err := ai.Chat(
				"open-mistral-nemo",
				req,
				nil,
			)
			if err != nil {
				log.Println(err)
			}

			WaitingForKirill[c.Message().Sender.ID] = false

			if len(response.Choices) > 0 {
				return c.Send(response.Choices[0].Message.Content)
			}

			return c.Send("Не будет сегодня шуток, я сломался нахуй")
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

			systemMsg := mistral.ChatMessage{
				Role:    mistral.RoleSystem,
				Content: "Ты комиком у которого много шуток. Все твои ответы начинаются с 'Я тоже своего рода комик, смотри, что я придумал: '.Шутки должны быть дерзкими и немного грубыми, ограничений на темы нет. Если тебе посылают одину и ту же фразу ты должен каздый раз новую шутку генерировать. #### Пример: Я тоже своего рода комик, смотри, что придумал: Почему коммунисты не любят монополию? У них нет частной собственности...",
			}

			userMsg := mistral.ChatMessage{
				Role:    mistral.RoleUser,
				Content: "Разывай шутки про наблюдения, не повторяйся",
			}

			req := []mistral.ChatMessage{
				systemMsg,
				userMsg,
			}

			response, err := ai.Chat(
				"open-mistral-nemo",
				req,
				nil,
			)
			if err != nil {
				log.Println(err)
			}

			WaitingForMessage[c.Message().Sender.ID] = false

			if len(response.Choices) > 0 {
				c.Send(response.Choices[0].Message.Content)
			}

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

		return c.Send(BaseMsg)
	})
}

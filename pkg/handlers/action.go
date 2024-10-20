package handlers

import "gopkg.in/telebot.v4"

func OtherHandlers(b *telebot.Bot) {
	b.Handle(telebot.OnSticker, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			return nil
		}
		return c.Send("У меня нет глаз, но я вижу, что ты охуел")
	})

	b.Handle(telebot.OnVoice, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			return nil
		}
		return c.Send("У меня нет ушей, но я слышу, что ты охуел")
	})

	b.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			return nil
		}
		return c.Send("У меня нет глаз, но я вижу, что ты охуел")
	})

	b.Handle(telebot.OnLocation, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			return nil
		}
		return c.Send("У меня нет карты, но, я знаю где ты охуел")
	})

	b.Handle(telebot.OnVideo, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			return nil
		}
		return c.Send("У меня нет глаз, но я вижу, что ты охуел")
	})

	b.Handle(telebot.OnVideoNote, func(c telebot.Context) error {
		if c.Chat().ID == ChatID {
			return nil
		}
		return c.Send("У меня нет глаз, но я вижу, что ты охуел")
	})
}

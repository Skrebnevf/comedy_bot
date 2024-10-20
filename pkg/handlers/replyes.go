package handlers

import "gopkg.in/telebot.v4"

func ReplyHandler(b *telebot.Bot) {
	b.Handle(telebot.OnReply, func(replyCtx telebot.Context) error {
		replyMsg := replyCtx.Message()

		if ForwardedMsg != nil && replyMsg.ReplyTo != nil && replyMsg.ReplyTo.ID == ForwardedMsg.ID {
			b.Send(&telebot.User{ID: OriginalUserID}, ReplyMsg+replyMsg.Text)
		}
		return nil
	})
}

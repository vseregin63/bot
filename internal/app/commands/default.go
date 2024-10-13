package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message){
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, inputMessage.Text)
	msg.ReplyToMessageID = inputMessage.MessageID

	c.bot.Send(msg)

}


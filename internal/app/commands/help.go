package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)


func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, 
		"/help - help message\n"+
		"/list - list of products\n"+
		"/get - get product params")
	c.bot.Send(msg)
}


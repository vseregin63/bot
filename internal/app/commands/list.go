package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	outputMsgTxt := "Here all the proiducts: \n\n"
	producs := c.productService.List()

	for _, p := range producs {
		outputMsgTxt += p.Title
		outputMsgTxt += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMsgTxt)
	c.bot.Send(msg)
}

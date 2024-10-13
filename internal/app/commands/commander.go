package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/vseregin63/bot/internal/service/product"
)


type Commander struct {
	bot *tgbotapi.BotAPI
	productService *product.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	productService *product.Service,
	) *Commander {
	return &Commander{
		bot: bot,
	}
}

func (c *Commander) HandleUpdate(update tgbotapi.Update){
	if update.Message != nil {
		switch update.Message.Command() {
		case "help":
			c.Help(update.Message)

		case "list":
			c.List(update.Message)

		case "get":
			c.Get(update.Message)

		default:
			c.Default(update.Message)
		}
	}	
}

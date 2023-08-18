package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

var bot *tgbotapi.BotAPI

func Init() {
	botInstance, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		panic(err)
	}
	log.Printf("Authorized on account %s", botInstance.Self.UserName)

	bot = botInstance
}

func GetBot() *tgbotapi.BotAPI {
	return bot
}

func SendMessage(c tgbotapi.Chattable) {
	_, err := bot.Send(c)
	if err != nil {
		log.Println("Error sending message:", err.Error())
	}
}

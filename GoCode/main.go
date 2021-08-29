package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq" // here
	"github.com/spf13/viper"
)

func runViper() {
	viper.SetConfigName("keys")
	viper.SetConfigType("json")
	viper.AddConfigPath("./config/")

	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		log.Fatal(err)
	}

}

func main() {
	runViper()

	bot, err := tgbotapi.NewBotAPI(viper.GetString("BotToken"))

	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

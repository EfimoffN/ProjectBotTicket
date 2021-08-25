package main

import (
	"flag"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	_ "github.com/lib/pq" // here
)

var (
	BotToken        = flag.String("BotToken", "telegram_token", "Path to file keys.json")
	BindAddr        = flag.String("BindAddr", "localhost", "Path to file keys.json")
	LogLevel        = flag.String("LogLevel", "debug", "Path to file keys.json")
	ConnectPostgres = flag.String("ConnectPostgres", "localhost", "Path to file keys.json")
)

func main() {
	flag.Parse()

	bot, err := tgbotapi.NewBotAPI(*BotToken)
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

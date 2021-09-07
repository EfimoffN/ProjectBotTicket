package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/jmoiron/sqlx"
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

	db, err := connectDB(viper.GetString("ConnectPostgres"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

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

func connectDB(databaseURL string) (*sqlx.DB, error) {

	fmt.Println("databaseURL")
	fmt.Println(databaseURL)
	fmt.Println("databaseURL")

	db, err := sqlx.Open("postgres", databaseURL)
	if err != nil {
		log.Println("sqlx.Open failed with an error: ", err.Error())
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Println("DB.Ping failed with an error: ", err.Error())
		return nil, err
	}

	return db, err
}

func ConnectDB_Test(databaseURL string) (*sqlx.DB, error) {
	db, err := connectDB(databaseURL)

	return db, err
}

func RunViper_Test() {
	runViper()
}

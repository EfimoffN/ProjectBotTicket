package main

import (
	"flag"
	"fmt"

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

	fmt.Println("Test")

}

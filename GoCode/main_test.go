package main

import (
	"log"
	"testing"

	"github.com/spf13/viper"
)

func TestMain(m *testing.M) {
	runViper()

	db, err := connectDB(viper.GetString("ConnectPostgres"))
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}

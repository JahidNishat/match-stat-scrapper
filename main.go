package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"match_statistics_scrapper/db"
	"match_statistics_scrapper/job"
	"match_statistics_scrapper/notifier"
)

func main() {
	viper.SetConfigFile("./config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}
	db.ConnectDB()
	notifier.LoadAdmin()
	go job.StartTicker()
	fmt.Println("Ticker Initiated, now serving bot....")

	notifier.ServeTgBot()
}

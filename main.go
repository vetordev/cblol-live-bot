package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"time"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load env file: %s", err.Error())
	}
}

func main() {
	loadEnv()

	matchTime, err := time.Parse(time.RFC3339, "2023-07-08T16:00:00Z")
	location, err := time.LoadLocation("America/Sao_Paulo")
	matchTime = time.Date(matchTime.Year(), matchTime.Month(), matchTime.Day(), 0, 0, 0, 0, location)
	now := time.Now()

	fmt.Println(matchTime, err, matchTime.After(now))
	//bot := telegrambot.New(true)
	//
	//bot.Run()
}

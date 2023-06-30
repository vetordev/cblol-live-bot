package main

import (
	telegrambot "cblol-bot/interface/telegram"
	"github.com/joho/godotenv"
	"log"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("could not load env file: %s", err.Error())
	}
}

func main() {
	loadEnv()

	bot := telegrambot.New(true)

	bot.Run()

}

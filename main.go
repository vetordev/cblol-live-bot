package main

import (
	"cblol-bot/application/match"
	"cblol-bot/application/ranking"
	telegrambot "cblol-bot/interface/telegram"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("could not load env file: %s", err.Error())
	}
}

func main() {
	loadEnv()

	telegramToken := os.Getenv("TELEGRAM_TOKEN")

	if telegramToken == "" {
		log.Fatal("TELEGRAM_TOKEN is empty")
	}

	lolApiKey := os.Getenv("LOL_API_KEY")

	if lolApiKey == "" {
		log.Fatal("LOL_API_KEY is empty")
	}

	lang := os.Getenv("LOL_API_LANG")

	if lolApiKey == "" {
		log.Fatal("LOL_API_LANG is empty")
	}

	debug, err := strconv.ParseBool(os.Getenv("DEBUG"))

	if err != nil {
		debug = false
	}

	matchApplication := match.New(lolApiKey, lang)
	rankingApplication := ranking.New(lolApiKey, lang)

	commandHandler := telegrambot.NewCommand(rankingApplication, matchApplication)

	bot := telegrambot.New(commandHandler, telegramToken, debug)

	bot.Run()

}

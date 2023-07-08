package main

import (
	"cblol-bot/application/match"
	"cblol-bot/application/ranking"
	"cblol-bot/infra/database"
	"cblol-bot/infra/scheduler"
	telegrambot "cblol-bot/interface/telegram"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

/*
@TODO: ref service package name to entitysvc
@TODO: ref application package to entityapp
*/

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

	databaseUrl := os.Getenv("DATABASE_URL")

	if databaseUrl == "" {
		log.Fatal("DATABASE_URL is empty")
	}

	database.RunMigrations(databaseUrl)

	s := scheduler.New()
	s.Load()

	matchApplication := match.New(lolApiKey, lang)
	rankingApplication := ranking.New(lolApiKey, lang)

	commandHandler := telegrambot.NewCommand(rankingApplication, matchApplication)

	bot := telegrambot.New(commandHandler, telegramToken, debug)

	bot.Run()

}

package main

import (
	"cblol-bot/application/match"
	"cblol-bot/application/notification"
	"cblol-bot/application/ranking"
	notificationsvc "cblol-bot/domain/service/notification"
	"cblol-bot/infra/database"
	"cblol-bot/infra/database/repository"
	"cblol-bot/infra/scheduler"
	"cblol-bot/infra/scheduler/job"
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

	migrationsPath := os.Getenv("MIGRATIONS_PATH")

	if migrationsPath == "" {
		log.Fatal("MIGRATIONS_PATH is empaty")
	}

	bot := telegrambot.New(telegramToken, debug)
	s := scheduler.New()
	db := database.Connect(databaseUrl)

	userRepository := repository.NewUserRepository(db)
	notificationRepository := repository.NewNotificationRepository(db)

	notificationService := notificationsvc.New(s, bot)

	matchApplication := match.New(lolApiKey, lang)
	rankingApplication := ranking.New(lolApiKey, lang)
	notificationApplication := notification.New(
		matchApplication,
		userRepository,
		notificationRepository,
		notificationService,
	)

	commandHandler := telegrambot.NewCommand(rankingApplication, matchApplication, notificationApplication)

	notificationJob := job.NewJobNotification(notificationApplication, s)

	notificationJob.Schedule()

	database.RunMigrations(migrationsPath, databaseUrl)
	bot.Run(commandHandler)
}

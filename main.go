package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	loadEnvironment()

	botID := os.Getenv("BOT_ID")

	bot, err := tgbotapi.NewBotAPI(botID)

	if err != nil {
		panic(err)
	}

    isDebug, err := strconv.ParseBool(os.Getenv("DEBUG"))
    if err != nil {
        panic(err)
    }
	bot.Debug = isDebug

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch update.Message.Text {
		case "/start":
			registerUser()
		
		case "/hello":
			sendHello(bot, &update)
		}
	}
}

func loadEnvironment() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file. '.env' file not found.")
	}
}

func sendHello(bot *tgbotapi.BotAPI, update *tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "world")
	msg.ReplyToMessageID = update.Message.MessageID

	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func registerUser() {
	log.Printf("User registration is not implmented yet")
}

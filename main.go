package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strings"
)

var (
	bot, err = tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
)

func processQuery(update *tgbotapi.Update) {
	location, err := downloadyou(update.Message.Text)
	if err == nil {
		path := strings.Split(location, "\n")[0]
		audio := tgbotapi.NewAudio(update.Message.Chat.ID, tgbotapi.FilePath(path))
		_, err = bot.Send(audio)
		if err != nil {
			log.Print(err)
			return
		}
		e := os.Remove(path)
		if e != nil {
			log.Print("Failed to remove: " + path)
		}
	} else {
		log.Print(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Couldn't process your request")
		bot.Send(msg)
	}
}

func main() {
	if err != nil {
		panic(err)
	}

	bot.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Loading...")
		bot.Send(msg)
		processQuery(&update)
	}
}

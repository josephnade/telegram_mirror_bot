package main

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("BOT_TOKENINIZ")
	if err != nil {
		log.Fatalf("Bot API oluşturulamadı: %v", err)
	}

	bot.Debug = false
	log.Printf("Bot @%s olarak başladı", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // Boş mesajları yoksay
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		responseMessage := fmt.Sprintf("Merhaba, mesajınızı aldım! '%s'", update.Message.Text)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseMessage)

		if _, err := bot.Send(msg); err != nil {
			log.Printf("Mesaj gönderilemedi: %v", err)
		}
	}
}

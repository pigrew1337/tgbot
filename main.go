package main

import (
	"fmt"
	"log"
	"time"

	"math/rand"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	fmt.Print("началось")
	bot, err := tgbotapi.NewBotAPI("7803998672:AAE-5_IU7mTzk7ksblgYMsNclqrQpj-9Bi0")
	if err != nil {
		fmt.Printf("12 ошибка апи бота %v\n", err)
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

		kal := tgbotapi.NewMessage(update.Message.Chat.ID, "ДТ РЕАЛЬНО КРУТОЙ ПАЦАН")
		if _, err := bot.Send(kal); err != nil {
			log.Println("27 ошибка отправки строки", err)
			panic(err)
		}

		randomNumber := random()
		var photoPath string
		switch randomNumber {
		case 1:
			photoPath = "photo/6le4G9n2Ud4.jpg"
		case 2:
			photoPath = "photo/mgN4LO0GeMw.jpg"
		case 3:
			photoPath = "photo/rfhgdhdh23243.png"
		}
		if photoPath != "" {
			photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FilePath(photoPath))
			_, err = bot.Send(photo)
			if err != nil {
				log.Println("Ошибка отправки фото:", err)
			}
		}
	}
}

func random() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(3) + 1
}

package main

import (
	"log"
	"os"

	"doctor-bot/internal/database"
	"doctor-bot/internal/handlers"
	"doctor-bot/internal/repository"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	db := database.Connect()
	defer db.Close()

	doctorRepo := repository.NewDoctorRepository(db)
	PatientReto := repository.NewPatientRepository(db)

	handler := handlers.NewDoctorHandler(doctorRepo, PatientReto)

	token := os.Getenv("BOT_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				handler.HandleStart(bot, update)
			case "add":
				handler.HandleAdd(bot, update)
			}
		}
	}
}

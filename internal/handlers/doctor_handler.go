package handlers

import (
	"doctor-bot/internal/models"
	"doctor-bot/internal/repository"
	_ "fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type DoctorHandler struct {
	doctorRepo  *repository.DoctorRepository
	patientRepo *repository.PatientRepository
}

func NewDoctorHandler(
	doctorRepo *repository.DoctorRepository,
	patientRepo *repository.PatientRepository,
) *DoctorHandler {
	return &DoctorHandler{
		doctorRepo:  doctorRepo,
		patientRepo: patientRepo}
}

func (h *DoctorHandler) HandleStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	tgID := update.Message.From.ID

	doctor, _ := h.doctorRepo.GetByTelegramID(tgID)
	if doctor != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"С возвращением, доктор 👨‍⚕️")
		bot.Send(msg)
		return
	}

	newDoctor := models.Doctor{
		TelegramID:     tgID,
		FullName:       update.Message.From.FirstName,
		Specialization: "Не указана",
	}

	h.doctorRepo.Create(&newDoctor)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"Вы зарегистрированы как врач ✅")
	bot.Send(msg)
}

func (h *DoctorHandler) HandleAdd(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	input := update.Message.Text
	cleaned := strings.TrimPrefix(input, "/add ")
	parts := strings.Split(cleaned, ";")

	if len(parts) != 3 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Вы ввели неверные данные ❌")
		bot.Send(msg)
		return
	}

	fullName := strings.TrimSpace(parts[0])
	ageStr := strings.TrimSpace(parts[1])
	diagnosis := strings.TrimSpace(parts[2])

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Возраст должен быть числом!")
		bot.Send(msg)
		return
	}

	log.Println(fullName, age, diagnosis)

	newPatient := models.Patient{
		FullName:  fullName,
		Age:       age,
		Diagnosis: diagnosis,
	}

	err = h.patientRepo.Create(&newPatient)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Ошибка при добавлении пациента")
		bot.Send(msg)
		log.Println("Ошибка при сохранении пациента:", err)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"Пациент успешно добавлен ✅")
	bot.Send(msg)

}

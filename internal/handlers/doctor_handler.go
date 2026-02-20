package handlers

import (
	"doctor-bot/internal/models"
	"doctor-bot/internal/repository"
	"fmt"
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
			`Укажите данные пациента в следующем виде:

✅ФИО; возраст; диагноз`)
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

func (h *DoctorHandler) HandleList(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	patients, err := h.patientRepo.GetAll()
	if err != nil {
		log.Println("Ошибка при попытке получить список пациентов", err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Произошла ошибка при получении данных")
		bot.Send(msg)
		return
	}

	if len(patients) == 0 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Список пациентов пуст")
		bot.Send(msg)
		return
	}

	var builder strings.Builder
	builder.WriteString("📋 Список пациентов:\n\n")

	for i, patient := range patients {
		builder.WriteString(fmt.Sprintf(
			"%d. 👤 %s\n   Возраст: %d лет\n   Диагноз: %s\n\n",
			i+1,
			patient.FullName,
			patient.Age,
			patient.Diagnosis,
		))

	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, builder.String())
	bot.Send(msg)

}

func (h *DoctorHandler) HandleDelete(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	input := update.Message.Text
	cleaned := strings.TrimPrefix(input, "/delete ")
	cleaned = strings.TrimSpace(cleaned)

	id, err := strconv.Atoi(cleaned)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"ID должен быть числом")
		bot.Send(msg)
		return
	}

	if cleaned == "" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Укажите ID пациента: /delete 3")
		bot.Send(msg)
		return
	}

	err = h.patientRepo.DeleteByID(id)
	if err != nil {
		log.Println("Ошибка при попытке удалить пациента", err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID,
			"Ошибка при удалении пациента")
		bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID,
		"Пациент успешно удален")
	bot.Send(msg)
}

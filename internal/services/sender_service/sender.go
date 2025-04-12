package sender

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

// EmailMessage — структура входящих сообщений
type EmailMessage struct {
	Type  string `json:"type"` // "otp", "reset", "welcome"
	Email string `json:"email"`
	OTP   string `json:"otp,omitempty"` // Только для OTP
}

// Карта обработчиков email
var handlers = map[string]func(EmailMessage){
	"otp":     sendOTPEmail,
	"reset":   sendResetPasswordEmail,
	"welcome": sendWelcomeEmail,
}

// Обработчик OTP
func sendOTPEmail(msg EmailMessage) {
	fmt.Printf("📨 Отправляем OTP %s на %s\n", msg.OTP, msg.Email)
}

// Обработчик Reset Password
func sendResetPasswordEmail(msg EmailMessage) {
	fmt.Printf("🔄 Отправляем письмо сброса пароля на %s\n", msg.Email)
}

// Обработчик Welcome Email
func sendWelcomeEmail(msg EmailMessage) {
	fmt.Printf("🎉 Отправляем приветственное письмо на %s\n", msg.Email)
}

// Обработчик Kafka-сообщений
func HandleEmailMessage(msg kafka.Message) {
	var data EmailMessage
	err := json.Unmarshal(msg.Value, &data)
	if err != nil {
		log.Println("❌ Ошибка парсинга JSON:", err)
		return
	}

	// Вызываем нужный обработчик
	if handler, exists := handlers[data.Type]; exists {
		handler(data)
	} else {
		log.Println("⚠ Неизвестный тип email:", data.Type)
	}
}

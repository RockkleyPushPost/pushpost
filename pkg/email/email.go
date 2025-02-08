package email

import (
	"fmt"
	"net/smtp"
)

func SendEmail(to, otp string) error {
	from := "rockkley94@gmail.com"
	password := "tylq bkiu mxkm owfi"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Your Verification Code"
	body := fmt.Sprintf("Your verification code is: %s\nIt will expire in 5 minutes.", otp)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}

package email

import (
	"net/smtp"
)

func SendEmail(to, subject, text string) error {
	from := "rockkley94@gmail.com"
	password := "tylq bkiu mxkm owfi"
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		text

	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(msg))
}

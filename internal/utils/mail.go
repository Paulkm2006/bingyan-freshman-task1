package utils

import (
	"bingyan-freshman-task0/internal/config"

	"math/rand"
	"strconv"

	"net/smtp"
)

func GenerateValidationCode() string {
	return strconv.Itoa(100000 + rand.Intn(900000))
}

func SendValidation(email string, code string) error {
	auth := smtp.PlainAuth("", config.Config.Mail.Username, config.Config.Mail.Password, config.Config.Mail.Host)
	to := []string{email}
	nickname := config.Config.Mail.Nickname
	user := config.Config.Mail.Username
	exp := config.Config.Mail.Expire

	subject := "BBingyan Email Validation"
	content_type := "Content-Type: text/plain; charset=UTF-8"
	body := "Your validation code is: " + code + "\r\nThis code will expire in " + strconv.Itoa(exp) + " minutes."
	msg := []byte("To: " + email + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(config.Config.Mail.Host+":"+config.Config.Mail.Port, auth, user, to, msg)
	return err
}

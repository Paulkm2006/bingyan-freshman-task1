package service

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/dto"
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"sync"
)

type EmailTemplate struct {
	Code   string
	Expire int
}

type WeeklyTemplate struct {
	Posts []dto.Post
}

var (
	captchaTemplate *template.Template
	weeklyTemplate  *template.Template
)

func init() {
	var err error
	captchaTemplate, err = template.ParseFiles("templates/captcha.html")
	if err != nil {
		panic(err)
	}
	weeklyTemplate, err = template.ParseFiles("templates/weekly.html")
	if err != nil {
		panic(err)
	}
}

func SendValidation(email string, code string) error {
	exp := config.Config.Captcha.Expire

	subject := config.Config.Captcha.Title

	var buffer bytes.Buffer
	templateData := EmailTemplate{
		Code:   code,
		Expire: exp,
	}

	err := captchaTemplate.Execute(&buffer, templateData)
	if err != nil {
		return err
	}

	return send(email, subject, buffer.String())
}

func SendWeeklyDigest(users []dto.User, posts []dto.Post) error {

	emails := make(chan string)
	emails_admin := make(chan string)

	subject := "Weekly Post Digest"

	var buffer bytes.Buffer
	templateData := WeeklyTemplate{
		Posts: posts,
	}

	err := weeklyTemplate.Execute(&buffer, templateData)
	if err != nil {
		return err
	}

	cnt_user := 0
	cnt_success := 0

	for _, user := range users {
		if user.Permission == 1 {
			emails_admin <- user.Email
		} else {
			emails <- user.Email
			cnt_user++
		}
	}

	var wg sync.WaitGroup
	type emailError struct {
		email string
		err   error
	}
	var errs []emailError

	for i := 0; i < config.Config.Mail.Workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			email, ok := <-emails
			if !ok {
				return
			}
			err := send(email, subject, buffer.String())
			if err != nil {
				errs = append(errs, emailError{email: email, err: err})
			} else {
				cnt_success++
			}
		}()
	}
	wg.Wait()

	close(emails)

	var adm_buffer bytes.Buffer

	adm_buffer.WriteString(fmt.Sprintf("Successfully sent weekly digest to %d out of %d users\n", cnt_success, cnt_user))

	if len(errs) > 0 {
		adm_buffer.WriteString("Failed to send weekly digest to the following users:\n")
		for _, e := range errs {
			adm_buffer.WriteString(fmt.Sprintf("Email: %s, Error: %s\n", e.email, e.err.Error()))
		}

	}
	for adm := range emails_admin {
		go send(adm, "Weekly digest report", adm_buffer.String())
	}

	return nil
}

func send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", config.Config.Mail.Username, config.Config.Mail.Password, config.Config.Mail.Host)
	nickname := config.Config.Mail.Nickname
	user := config.Config.Mail.Username
	msg := []byte("To: " + to + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n" + body)
	return smtp.SendMail(config.Config.Mail.Host+":"+config.Config.Mail.Port, auth, config.Config.Mail.Username, []string{to}, msg)
}

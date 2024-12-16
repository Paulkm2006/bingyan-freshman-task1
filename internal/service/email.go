package service

import (
	"bingyan-freshman-task0/internal/config"
	"bingyan-freshman-task0/internal/dto"
	"bingyan-freshman-task0/internal/utils"
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"sync"
	"sync/atomic"
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
	subject := "Weekly Post Digest"

	var buffer bytes.Buffer
	templateData := WeeklyTemplate{
		Posts: posts,
	}

	err := weeklyTemplate.Execute(&buffer, templateData)
	if err != nil {
		utils.Logger.Error(err.Error())
		return err
	}

	cnt_user := 0
	cnt_success := int32(0)
	var adminEmails []string

	emails := make(chan string, len(users))

	for _, user := range users {
		if user.Email == "" {
			continue
		}
		if user.Permission == 1 {
			adminEmails = append(adminEmails, user.Email)
		} else {
			emails <- user.Email
			cnt_user++
		}
	}

	close(emails)
	go func() {

		var wg sync.WaitGroup
		type emailError struct {
			email string
			err   error
		}
		errs := make([]emailError, 0)
		var errMutex sync.Mutex

		for i := 0; i < config.Config.Mail.Workers; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for email := range emails {
					err := send(email, subject, buffer.String())
					if err != nil {
						errMutex.Lock()
						errs = append(errs, emailError{email: email, err: err})
						errMutex.Unlock()
					} else {
						atomic.AddInt32(&cnt_success, 1)
					}
				}
			}()
		}

		wg.Wait()

		var adm_buffer bytes.Buffer
		adm_buffer.WriteString(fmt.Sprintf("Successfully sent weekly digest to %d out of %d users\n", cnt_success, cnt_user))

		if len(errs) > 0 {
			adm_buffer.WriteString("Failed to send weekly digest to the following users:\n")
			for _, e := range errs {
				adm_buffer.WriteString(fmt.Sprintf("Email: %s, Error: %s\n", e.email, e.err.Error()))
				utils.Logger.Error(fmt.Sprintf("Failed to send weekly digest to %s: %s", e.email, e.err.Error()))
			}
		}

		for _, adminEmail := range adminEmails {
			err := send(adminEmail, "Weekly digest report", adm_buffer.String())
			if err != nil {
				utils.Logger.Error(fmt.Sprintf("Failed to send digest report to admin %s: %s", adminEmail, err.Error()))
			}
		}
	}()
	return nil
}

func send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", config.Config.Mail.Username, config.Config.Mail.Password, config.Config.Mail.Host)
	nickname := config.Config.Mail.Nickname
	user := config.Config.Mail.Username
	msg := []byte("To: " + to + "\r\nFrom: " + nickname + "<" + user + ">\r\nSubject: " + subject + "\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n" + body)
	return smtp.SendMail(config.Config.Mail.Host+":"+config.Config.Mail.Port, auth, config.Config.Mail.Username, []string{to}, msg)
}

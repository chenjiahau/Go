package util

import "gopkg.in/gomail.v2"

func SendEmail(smtpHost string, smtpPort int, smtpUser, smtpPass, from, to, subject, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)

	// HTML body
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtpHost, smtpPort, smtpUser, smtpPass)
	if err := d.DialAndSend(m); err != nil {
			return err
	}

	return nil
}
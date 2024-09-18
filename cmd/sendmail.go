package main

import "gopkg.in/gomail.v2"

func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "")
	m.SetHeader("To", "")
	m.SetHeader("Subject", "Hello!")

	// Plain text body
	// m.SetBody("text/plain", "This is the plain text body of the email.")

	// HTML body
	m.SetBody("text/html", "Hello <b>World</b>!")

	d := gomail.NewDialer("", 587, "", "")
	if err := d.DialAndSend(m); err != nil {
			panic(err)
	}
}

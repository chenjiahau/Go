package internal

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"text/template"
)

type loginAuth struct {
	username string
	password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
			switch string(fromServer) {
			case "Username:":
					return []byte(a.username), nil
			case "Password:":
					return []byte(a.password), nil
			default:
					return nil, errors.New("unknown from server")
			}
	}

	return nil, nil
}

func SendEmail(from, password, title, message string) {
	to := []string{
		from,
  }

	smtpHost := "smtp.gmail.com"
  smtpPort := "587"

	conn, err := net.Dial("tcp", "smtp.gmail.com:587")
    if err != nil {
        println(err)
    }

    c, err := smtp.NewClient(conn, smtpHost)
    if err != nil {
        println(err)
    }

		tlsconfig := &tls.Config{
			ServerName: smtpHost,
	}

	if err = c.StartTLS(tlsconfig); err != nil {
			println(err)
	}

	auth := LoginAuth(from, password)
	if err = c.Auth(auth); err != nil {
		println(err)
  }

	t, _ := template.ParseFiles("./templates/email.tpl")

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: %s \n%s\n\n", title, mimeHeaders)))

	t.Execute(&body, struct {
			Message string
	}{
			Message: message,
	})

	err = smtp.SendMail(smtpHost + ":" + smtpPort, auth, from, to, body.Bytes())
	if err != nil {
			fmt.Println(err)
			return
	}

	fmt.Println("Email Sent!")
}
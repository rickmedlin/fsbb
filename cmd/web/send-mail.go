package main

import (
	"fsbb/internal/models"
	"time"

	mail "github.com/xhit/go-simple-mail"
)

func ListenForMail() {
	go func() {
		for {
			msg := <-app.MailChan
		}
	}()
}

func SendMail(m models.MailData) {
	server := mail.NewSMTPClient
	server.Host = "localhost"
	server.Port = 1025
	server.Keepalive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second
}

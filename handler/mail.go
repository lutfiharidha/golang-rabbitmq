package handler

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(msg []byte) string {
	mailUser := "a3bb3169d56279"
	mailPass := "7f5ce6d6c4d5eb"
	mailHost := "smtp.mailtrap.io"
	mailPort := "2525"
	to := "blokupie@gmail.com"

	// Set up authentication information.
	auth := smtp.PlainAuth("", mailUser, mailPass, mailHost)

	// format smtp address
	smtpAddress := fmt.Sprintf("%s:%v", mailHost, mailPort)

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(smtpAddress, auth, mailUser, []string{to}, msg)

	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	// return true on success
	return to
}

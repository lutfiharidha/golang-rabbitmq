package handler

import (
	"fmt"
	"log"
	"net/smtp"
)

func SendEmail(msg []byte) string {
	to := "blokupie@gmail.com"
	// Set up authentication information.
	auth := smtp.PlainAuth("", "a3bb3169d56279", "7f5ce6d6c4d5eb", "smtp.mailtrap.io")

	// format smtp address
	smtpAddress := fmt.Sprintf("%s:%v", "smtp.mailtrap.io", "2525")

	// Connect to the server, authenticate, set the sender and recipient,
	// and send the email all in one step.
	err := smtp.SendMail(smtpAddress, auth, "a3bb3169d56279", []string{to}, msg)

	if err != nil {
		log.Fatal(err)
		return err.Error()
	}

	// return true on success
	return to
}

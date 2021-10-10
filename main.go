package main

import (
	"golang-rabbitmq/handler"
	"golang-rabbitmq/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

func sendEmail(c *gin.Context) {

	if service.Task(c) != nil {
		res := Response{
			Status:  false,
			Message: "ERROR",
			Errors:  nil,
		}
		c.JSON(http.StatusBadRequest, res)
	}

	res := Response{
		Status:  true,
		Message: "Successfully send a message",
		Errors:  nil,
	}
	c.JSON(http.StatusOK, res)
}
func rabbit(c *gin.Context) {
	res := Response{
		Status:  true,
		Message: "Welcome RabbitMQ",
		Errors:  nil,
	}
	c.JSON(http.StatusOK, res)
}

func biasa(c *gin.Context) {
	msg := []byte("To: blokupie@gmail.com\r\n" +
		"From: lutfiharidha1@gmail.com\r\n" +
		"Subject: Hello Gophers!\r\n" +
		"\r\n" +
		"This is the email is sent using golang and sendinblue.\r\n")
	handler.SendEmail(msg)

	res := Response{
		Status:  true,
		Message: "Successfully send a message",
		Errors:  nil,
	}
	c.JSON(http.StatusOK, res)
}
func main() {
	r := gin.Default()
	r.GET("/send-email", sendEmail)
	r.GET("/", rabbit)
	r.GET("/biasa", biasa)
	r.Run()

}

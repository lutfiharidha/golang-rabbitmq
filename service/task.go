package service

import (
	"golang-rabbitmq/handler"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func Task(c *gin.Context) error {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"task_queue", // name
		true,         // durable
		false,        // delete when unused
		false,        // exclusive
		true,         // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}
	msg := []byte("To: blokupie@gmail.com\r\n" +
		"From: lutfiharidha1@gmail.com\r\n" +
		"Subject: Hello Gophers!\r\n" +
		"\r\n" +
		"This is the email is sent using golang and sendinblue.\r\n")

	// body := bodyFrom(os.Args)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(handler.SendEmail(msg)),
		})
	if err != nil {
		return err
	}
	return nil
}

# Golang x RabbitMQ x Docker

i'm using Gin framework

- ``docker run -d --hostname my-rabbit --name some-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management``
- ``go get github.com/streadway/amqp``
- ``go get github.com/gin-gonic/gin``

change variable value `mailUser`, `mailPass`, `mailHost`, `mailPort`, `to` at ``handler/mail.go``,
and don't forget to edit variable value `msg` at `service/task.go`

run `go run worker.go` and run `go run main.go` together!

open your browser and go to:

http://localhost:15672/ -for RabbitMQ dashboard

http://127.0.0.1:8080/ - for home

http://127.0.0.1:8080/send-email - for send email without RabbitMQ

http://127.0.0.1:8080/send-email-rabbitmq - for send email with RabbitMQ


goodluck!

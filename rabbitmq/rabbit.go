package rabbitmq

import (
	"github.com/streadway/amqp"
	"log"
)

func GetRabbitMQ(host string) *amqp.Connection {
	conn, err := amqp.Dial(host)
	failOnError(err, "could not connect to rabbitmq")
	return conn
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

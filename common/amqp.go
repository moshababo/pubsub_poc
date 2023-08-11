package common

import (
	"github.com/streadway/amqp"
)

var (
	QueueName = "testing"
)

func NewConnection(url string, log *Logger) (*amqp.Connection, error) {
	log.Info("dialing %v", url)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	log.Info("connected successfully")

	return conn, err
}

func QueueDeclare(channel *amqp.Channel) (amqp.Queue, error) {
	return channel.QueueDeclare(
		QueueName, // name
		false,     // durable
		false,     // auto delete
		false,     // exclusive
		false,     // no wait
		nil,       // args
	)
}

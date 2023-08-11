package publisher

import (
	"fmt"
	"github.com/streadway/amqp"
	"pubsub_poc/common"
)

type Publisher struct {
	conn   *amqp.Connection
	logger *common.Logger
}

func New(conn *amqp.Connection, logger *common.Logger) *Publisher {
	return &Publisher{conn, logger.WithTag("PUBLISHER")}
}

func (p *Publisher) Publish(data []byte) error {
	channel, err := p.conn.Channel()
	if err != nil {
		return fmt.Errorf("open channel error: %w", err)
	}
	defer channel.Close()

	queue, err := common.QueueDeclare(channel)
	if err != nil {
		return fmt.Errorf("queue declare error: %w", err)
	}

	p.logger.Info("publishing data: 0x%x", data)

	err = channel.Publish(
		"",        // exchange
		"testing", // key
		false,     // mandatory
		false,     // immediate
		amqp.Publishing{ContentType: "application/octet-stream", Body: data},
	)
	if err != nil {
		return fmt.Errorf("publish error: %w", err)
	}

	p.logger.Info("published successfully, queue: %+v", queue)

	return nil
}

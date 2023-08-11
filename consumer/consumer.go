package consumer

import (
	"context"
	"github.com/streadway/amqp"
	"pubsub_poc/common"
	"pubsub_poc/wire"
)

type Consumer struct {
	conn   *amqp.Connection
	logger *common.Logger
}

func New(conn *amqp.Connection, logger *common.Logger) *Consumer {
	return &Consumer{conn, logger.WithTag("CONSUMER")}
}

func (c *Consumer) Consume(ctx context.Context) (<-chan wire.Msg, <-chan error, error) {
	c.logger.Info("opening channel")
	channel, err := c.conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = common.QueueDeclare(channel)
	if err != nil {
		return nil, nil, err
	}

	deliveriesChan, err := channel.Consume(
		common.QueueName, // queue
		"",               // consumer
		true,             // auto ack
		false,            // exclusive
		false,            // no local
		false,            // no wait
		nil,              //args
	)
	if err != nil {
		return nil, nil, err
	}

	msgChan := make(chan wire.Msg)
	errChan := make(chan error)
	go func() {
		defer func() {
			close(msgChan)
			close(errChan)
		}()

		c.logger.Info("starting")
		for {
			select {
			case d, ok := <-deliveriesChan:
				if !ok {
					c.logger.Info("channel closed")
					return
				}
				msg, err := parseMessage(d.Body)
				if err != nil {
					errChan <- err
					continue
				}
				msgChan <- msg
			case <-ctx.Done():
				c.logger.Info("closing")
				return
			}
		}
	}()

	return msgChan, errChan, nil
}

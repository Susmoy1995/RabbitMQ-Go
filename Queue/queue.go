package queue

import (
	"errors"

	"github.com/streadway/amqp"
)

func DeclareQueue(ch *amqp.Channel, queueName string) error {
	_, err := ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

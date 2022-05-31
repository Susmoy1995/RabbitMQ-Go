package consumer

import (
	channel "Rabbit-GOPkg/Channels"
	conn "Rabbit-GOPkg/Connection"
	"errors"
)

var body []byte

func Consumer(queueName string) ([]byte, error) {
	connection, err := conn.Connection()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer connection.Close()

	ch, err := channel.CreateChannel(connection)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, errors.New(err.Error())
	}

	for msg := range msgs {
		body = msg.Body
		msg.Ack(false)
	}

	return body, nil
}

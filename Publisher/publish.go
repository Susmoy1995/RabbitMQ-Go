package publisher

import (
	channel "Rabbit-GOPkg/Channels"
	conn "Rabbit-GOPkg/Connection"
	queue "Rabbit-GOPkg/Queue"
	"errors"

	"github.com/streadway/amqp"
)

func Publish(queueName string, msg []byte) (string, error) {

	// create connection to rabbitmq
	connection, err := conn.Connection()
	if err != nil {
		panic(err.Error())
	}
	defer connection.Close()

	// create channel
	ch, err := channel.CreateChannel(connection)
	if err != nil {
		// panic(err.Error())
		return "", errors.New(err.Error())
	}
	defer ch.Close()

	// declare a queue under channel
	err = queue.DeclareQueue(ch, queueName)
	if err != nil {
		// fmt.Printf("%v\n", err)
		return "", errors.New(err.Error())
	}

	// publish message to that queue
	err = ch.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        msg,
		},
	)

	if err != nil {
		return "", errors.New(err.Error())
	}

	return "Successfully Message Published", nil
}

package channels

import (
	"errors"

	"github.com/streadway/amqp"
)

func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		// fmt.Printf("%v\n", err)
		return nil, errors.New(err.Error())
	}

	return ch, nil
}

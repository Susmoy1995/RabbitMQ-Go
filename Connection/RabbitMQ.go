package connection

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

func Connection() (*amqp.Connection, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		fmt.Printf("%v\n", err)
		// panic(err.Error())
		return nil, errors.New(err.Error())
	}

	return conn, nil
}

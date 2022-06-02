package consumer

import (
	conn "Rabbit-GOPkg/Connection"
	"errors"
)

func Consume(queueName string) ([]byte, error) {
	// var cr conn.RabbitMQ
	var sub *conn.RabbitMQ
	var err error

	if sub == nil {
		sub, err = conn.New()
		if err != nil {
			return nil, errors.New(err.Error())
		}
	}

	// err = sub.OpenChannel()
	// if err != nil {
	// 	return nil, errors.New(err.Error())
	// }

	// err = sub.CreateQueue(queueName)
	// if err != nil {
	// 	return nil, errors.New(err.Error())
	// }

	result, err := sub.ConsumeMessage(queueName)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return result, nil
}

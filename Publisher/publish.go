package publisher

import (
	conn "Rabbit-GOPkg/Connection"
	"errors"
)

func Publish(queueName string, msg []byte) (string, error) {

	var pub *conn.RabbitMQ
	var err error

	if pub == nil {
		pub, err = conn.New()
		if err != nil {
			return "", errors.New(err.Error())
		}
	}

	// err = pub.OpenChannel()
	// if err != nil {
	// 	return "", errors.New(err.Error())
	// }

	err = pub.CreateQueue(queueName)
	if err != nil {
		return "", errors.New(err.Error())
	}

	err = pub.PublishMessage(queueName, msg)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return "Successfully Message Published", nil
}

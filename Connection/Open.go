package connection

import (
	"errors"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

var rq RabbitMQ

func New() (*RabbitMQ, error) {
	if rq.Connection != nil && rq.Channel != nil {
		return &rq, nil
	}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")

	if err != nil {
		rq = RabbitMQ{nil, nil}
		return &rq, errors.New(err.Error())
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		rq = RabbitMQ{nil, nil}
		return &rq, errors.New(err.Error())
	}

	rq = RabbitMQ{Connection: conn, Channel: ch}
	return &rq, nil
}

func (r *RabbitMQ) CloseConnection() {
	r.Channel.Close()
	r.Connection.Close()
}

func (r *RabbitMQ) CloseChannel() {
	r.Channel.Close()
}

func (r *RabbitMQ) CreateQueue(queueName string) error {
	_, err := r.Channel.QueueDeclare(
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

func (r *RabbitMQ) DeleteQueue() {}

func (r *RabbitMQ) PublishMessage(queueName string, msg []byte) error {
	err := r.Channel.Publish(
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
		return errors.New(err.Error())
	}

	return nil
}

func (r *RabbitMQ) ConsumeMessage(queueName string) ([]byte, error) {

	msg, err := r.Channel.Consume(
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
	var body []byte

	for d := range msg {
		body = d.Body
		d.Ack(true)
	}
	return body, nil
}

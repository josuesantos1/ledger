package component

import (
	"fmt"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (c *Component) QueueConnect(queueName string, addr string) {
	var conn *amqp.Connection
	var err error

	for attempts := 0; attempts < 3; attempts++ {
		conn, err = amqp.Dial(addr)
		if err == nil {
			break
		}

		fmt.Printf("Attempt %d failed to connect to RabbitMQ: %v\n", attempts+1, err)
		if attempts < 2 {
			time.Sleep(2 * time.Second)
		}
	}

	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ after 3 attempts: %v", err))
	}

	c.QueueConn = conn

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		panic(fmt.Sprintf("Failed to open channel: %v", err))
	}

	c.QueueChan = channel

	_, err = channel.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		channel.Close()
		conn.Close()
		panic(fmt.Sprintf("Failed to declare queue %s: %v", queueName, err))
	}
}

package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

const QUEUENAME string = "hotpotato"

func GetMessages(conn *amqp.Connection) (msgs <-chan amqp.Delivery) {
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")

	q := DeclareQueue(ch, QUEUENAME)
	msgs, err = ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	return
}

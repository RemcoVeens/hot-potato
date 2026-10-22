package queue

import (
	"encoding/json"
	"math/rand/v2"

	amqp "github.com/rabbitmq/amqp091-go"
)

type QueueMessage struct {
	Message   string `json:"message"`
	TrowsLeft int    `json:"trows"`
}

func MakeMessage(message string) QueueMessage {
	return QueueMessage{Message: message, TrowsLeft: rand.IntN(10)}
}
func (qm QueueMessage) ToBytes() (marshaled []byte) {
	marshaled, err := json.Marshal(qm)
	failOnError(err, "faild to decode json")
	return
}
func MessageFromDelivery(d amqp.Delivery) (msg QueueMessage) {
	err := json.Unmarshal(d.Body, &msg)
	failOnError(err, "could not decode")
	return

}

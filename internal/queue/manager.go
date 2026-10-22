package queue

import (
	"context"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

const QUEUENAME string = "hot potato"

type QueueManager struct {
	conn  *amqp.Connection
	queue amqp.Queue
	ch    *amqp.Channel
}

func GetQueueManager() QueueManager {
	return QueueManager{}
}

func (qm *QueueManager) Connect() {
	var err error
	qm.conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
}
func (qm *QueueManager) CloseConn() {
	qm.conn.Close()
}

func (qm *QueueManager) CloseChan() {
	qm.ch.Close()
}

func (qm *QueueManager) GetReciever(consumerName string) (msgs <-chan amqp.Delivery) {
	qm.openChannel()
	var err error
	q := DeclareQueue(qm.ch, QUEUENAME)
	msgs, err = qm.ch.Consume(
		q.Name,       // queue
		consumerName, // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	failOnError(err, "Failed to register a consumer")
	return
}
func (qm *QueueManager) openChannel() {
	var err error
	qm.ch, err = qm.conn.Channel()
	failOnError(err, "Failed to open a channel")
}
func (qm *QueueManager) SendMessage(msg QueueMessage) {
	var err error
	qm.openChannel()
	qm.queue = DeclareQueue(qm.ch, QUEUENAME)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = qm.ch.PublishWithContext(ctx,
		"",            // exchange
		qm.queue.Name, // routing key
		false,         // mandatory
		false,         // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msg.ToBytes(),
		})
	failOnError(err, "Failed to publish a message")
}

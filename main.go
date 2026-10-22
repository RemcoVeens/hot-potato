package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RemcoVeens/hot_potato/internal/queue"
	"github.com/icrowley/fake"
)

func main() {
	qm := queue.GetQueueManager()
	qm.Connect()
	defer qm.CloseConn()
	go func() {
		for i := range 5 {
			qm.SendMessage(queue.MakeMessage(fake.EmailSubject()))
			time.Sleep(time.Second * time.Duration(i))
		}
	}()
	reciever := qm.GetReciever(fake.FirstName())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done() // wait for Ctrl+C / SIGTERM
		stop()
		qm.CloseChan() // closes the AMQP channel → closes msgs → range ends
	}()
	for msg := range reciever {
		mess := queue.MessageFromDelivery(msg)

		log.Printf(" [*] Recieved: '%s' from: '%s'\n", mess.Message, msg.ConsumerTag)
	}
	fmt.Print("\b\b")
	log.Println("shutting down")

}

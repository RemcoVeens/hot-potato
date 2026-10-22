package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/RemcoVeens/hot_potato/internal/queue"
	"github.com/icrowley/fake"
)

func main() {
	qm := queue.QueueManager{}
	qm.Connect()
	defer qm.CloseConn()
	qm.SendMessage("this is a test")
	reciever := qm.GetMessages(fake.FirstName())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done() // wait for Ctrl+C / SIGTERM
		stop()
		qm.CloseChan() // closes the AMQP channel → closes msgs → range ends
	}()
	for msg := range reciever {
		log.Printf(" [*] Recieved: '%s' from: '%s'\n", msg.Body, msg.ConsumerTag)
	}
	fmt.Println("")
	log.Println("shutting down")

}

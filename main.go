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
	player := fake.FemaleFirstName()
	reciever := qm.GetReciever(player)
	log.Printf("%s is ready to play\n", player)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	go func() {
		<-ctx.Done() // wait for Ctrl+C / SIGTERM
		stop()
		qm.CloseChan() // closes the AMQP channel → closes msgs → range ends
	}()
	for msg := range reciever {
		mess := queue.MessageFromDelivery(msg)
		log.Printf("%s caught the %s\n", msg.ConsumerTag, mess.Message)
		if mess.TrowsLeft < 1 {
			log.Println("💥biem")
			os.Exit(0)
		}
		mess.TrowsLeft -= 1
		time.Sleep(time.Duration(1) * time.Second)
		qm.SendMessage(mess)

	}
	fmt.Print("\b\b")
	log.Println("shutting down")

}

package main

import (
	"github.com/RemcoVeens/hot_potato/internal/queue"
)

func main() {
	qm := queue.GetQueueManager()
	qm.Connect()
	defer qm.CloseConn()
	qm.SendMessage(queue.MakeMessage("potato"))
}

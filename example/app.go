package main

import (
	"github.com/ujunglangit-id/gotong-royong/pkg"
	"log"
	"sync"
	"time"
)

func showText() {
	time.Sleep(2 * time.Second)
	log.Printf(
		"remaining queue : %d, current channel length %d, execute current task ....\n",
		wk.GetRemainingQueueLength(), wk.GetChannelLength())
}

var (
	wk *pkg.WorkerContainer
)

func init() {
	wk = pkg.NewWorkers(5)
}

func main() {
	wg := sync.WaitGroup{}
	wk.RunInBackground()
	for i := 0; i < 100; i++ {
		go wk.AddNewEvent(showText)
	}
	wg.Add(1)
	wg.Wait()
}

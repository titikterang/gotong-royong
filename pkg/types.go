package pkg

import "sync"

type WorkerContainer struct {
	maxWorker    int
	workerQueue  chan bool
	wg           sync.WaitGroup
	mtx          sync.RWMutex
	eventHandler []func()
}

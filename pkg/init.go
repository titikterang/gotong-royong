package pkg

func NewWorkers(maxWorker int) *WorkerContainer {
	return &WorkerContainer{
		maxWorker:    maxWorker,
		workerQueue:  make(chan bool, maxWorker),
		eventHandler: []func(){},
	}
}

func (w *WorkerContainer) RunInBackground() {
	go w.Run()
}

func (w *WorkerContainer) GetChannelLength() (l int) {
	w.mtx.Lock()
	l = len(w.workerQueue)
	w.mtx.Unlock()
	return
}

func (w *WorkerContainer) GetRemainingQueueLength() (l int) {
	w.mtx.Lock()
	l = len(w.eventHandler)
	w.mtx.Unlock()
	return
}

func (w *WorkerContainer) Run() {
	//wait until total workers < max worker, and refresh worker queue
	for {
		if len(w.workerQueue) < w.maxWorker && len(w.eventHandler) > 0 {
			w.workerQueue <- true
			w.mtx.Lock()
			exec := w.eventHandler[0]
			w.eventHandler = w.eventHandler[1:]
			w.mtx.Unlock()
			go func(q chan bool) {
				exec()
				<-q
			}(w.workerQueue)
		}
	}
}

func (w *WorkerContainer) AddNewEvent(eventHandler func()) {
	w.mtx.Lock()
	w.eventHandler = append(w.eventHandler, eventHandler)
	w.mtx.Unlock()
}

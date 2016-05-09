package WebWorker

import (
	"fmt"

	"github.com/bolovsky/requester/WebRequest"
)

// Worker represents the worker that executes the job
type WebWorker struct {
	WorkerChannel	chan chan Job
	JobChannel		chan Job
	Quit			chan bool
	WorkerId		int
}

// NewWorker returns a new Worker
func NewWorker(workerPool chan chan Job, workerId int) WebWorker {
	return WebWorker{
		WorkerChannel:	workerPool,
		JobChannel:	make(chan Job),
		Quit:		make(chan bool),
		WorkerId:	workerId}
}

func (w WebWorker) Start() {
	fmt.Println("starting worker")
	go func() {
		req := WebRequester.New()

		for {
			// register the current worker into the worker queue.
			w.WorkerChannel <- w.JobChannel

			select {
				case job := <-w.JobChannel:
					resp, err := req.PostJson(job.Url, job.Payload)
					if nil != err {
						fmt.Println(err)
					} else {
						fmt.Println(resp)
					}
				case <-w.Quit:
					fmt.Println(fmt.Sprintf("quitting %d", w.WorkerId))
					return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests.
func (w WebWorker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

package WebWorker

import "fmt"

type Pool struct {
	// A pool of workers channels that are registered with the dispatcher
	WorkerChannel chan chan Job
	WorkerCollection []WebWorker
	MaxWorkers int
}

func NewPool(maxWorkers int) *Pool {
	pool := make(chan chan Job, maxWorkers)

	fmt.Println("pool starting")
	var wrkCl = make([]WebWorker, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		wrkCl[i] = NewWorker(pool, i)
		wrkCl[i].Start()
	}

	return &Pool{
		WorkerChannel: pool,
		WorkerCollection: wrkCl,
		MaxWorkers: maxWorkers,
	}
}

func (pool *Pool) QueueJob(job Job) {
	go func(job Job) {
		// try to obtain a worker job channel that is available.
		// this will block until a worker is idle
		jobChannel := <-pool.WorkerChannel

		// dispatch the job to the worker job channel
		jobChannel <- job
	}(job)
}

func (pool *Pool) ShutDown() (stat bool) {
	for _, wrk := range pool.WorkerCollection {
		wrk.Stop()
	}
	fmt.Println("pool stopped")

	return true
}
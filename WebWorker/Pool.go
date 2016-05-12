package WebWorker

// Pool of workers channels that are registered with the dispatcher
type Pool struct {
	WorkerChannel    chan chan Job
	WorkerCollection []WebWorker
	MaxWorkers       int
}

// NewPool returns a new instance of Pool
func NewPool(maxWorkers int, ResponseChannel chan JobResponse) *Pool {
	pool := make(chan chan Job, maxWorkers)

	var wrkCl = make([]WebWorker, maxWorkers)

	for i := 0; i < maxWorkers; i++ {
		wrkCl[i] = NewWorker(pool, ResponseChannel, i)
		wrkCl[i].Start()
	}

	return &Pool{
		WorkerChannel:    pool,
		WorkerCollection: wrkCl,
		MaxWorkers:       maxWorkers,
	}
}

// QueueJob queues a job to be attempted
func (pool *Pool) QueueJob(job Job) {
	go func(job Job) {
		// try to obtain a worker job channel that is available.
		// this will block until a worker is idle
		jobChannel := <-pool.WorkerChannel

		// dispatch the job to the worker job channel
		jobChannel <- job
	}(job)
}

// ShutDown terminates Pool
func (pool *Pool) ShutDown() (stat bool) {
	for _, wrk := range pool.WorkerCollection {
		wrk.Stop()
	}

	return true
}

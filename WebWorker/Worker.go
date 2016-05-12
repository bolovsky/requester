package WebWorker

// WebWorker represents the worker that executes the job
type WebWorker struct {
	WorkerChannel   chan chan Job
	JobChannel      chan Job
	ResponseChannel chan JobResponse
	Quit            chan bool
	WorkerID        int
}

// NewWorker returns a new Worker
func NewWorker(
	workerPool chan chan Job,
	ResponseChannel chan JobResponse,
	workerID int,
) WebWorker {
	return WebWorker{
		WorkerChannel: workerPool,
		JobChannel:    make(chan Job),
		Quit:          make(chan bool),
		WorkerID:      workerID,
	}
}

// Start initiates WebWorker
func (w WebWorker) Start() {
	go func() {
		req := NewRequester()

		for {
			// register the current worker into the worker queue.
			w.WorkerChannel <- w.JobChannel

			select {
			case job := <-w.JobChannel:
				resp, err := req.PostJSON(job.URL, job.Payload)
				if nil != err {
					w.ResponseChannel <- JobResponse{
						WorkerID: w.WorkerID,
						Status:   false,
						Response: "",
					}
				} else {
					w.ResponseChannel <- JobResponse{
						WorkerID: w.WorkerID,
						Status:   true,
						Response: resp,
					}
				}
			case <-w.Quit:
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

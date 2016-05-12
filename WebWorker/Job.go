package WebWorker

// Job represents the job to be run
type Job struct {
	Payload string
	URL     string
}

// JobResponse represents the response for a processed job
type JobResponse struct {
	WorkerID int
	Status   bool
	Response string
}

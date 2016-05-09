package main

import (
	"fmt"
	"time"

	"github.com/bolovsky/requester/WebWorker"
)

var pool *WebWorker.Pool

func main() {
	fmt.Printf("Starting Server...\n")

	pool = WebWorker.NewPool(5)

	// do requests ad eternum or until interrupt
	doRequests()

	pool.ShutDown()

	fmt.Printf("Stopping Server...\n")
}

func doRequests() {
	// just demostrative purpose, will implement possibly socket receive, possible external queue read
	for {
		for i := 0; i < 20; i += 1 {
			pool.QueueJob(
				WebWorker.Job{
					Payload: fmt.Sprintf("{\"SomePayloadKey\": %d}", i),
					Url:     "http://localhost:8080",
				})
		}

		time.Sleep(5 * time.Second)
	}
}

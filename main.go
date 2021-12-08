package main

import (
	"time"

	"github.com/secmohammed/golang-basic-concurrent-queue/worker"
)

func main() {
	// if calling defer worker.Wait here, it will block the main thread until all of the workers are done but all of them will fire at the same time. it won't wait for 2 sec for each (due to the print payload is sleeping for 2 secs), so It's firing asynchronously.
	// defer worker.Wait()
	for i := 0; i < 1000; i++ {
		job := worker.Job{
			Action: PrintPayload,
			Payload: map[string]string{
				"message": "Hello World",
				"time":    time.Now().String(),
			},
		}
		job.Fire()

	}
}

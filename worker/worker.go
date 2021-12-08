package worker

import "sync"

type Job struct {
	id      string
	Action  func(map[string]string)
	Payload map[string]string
}

var wg sync.WaitGroup

func Wait() {
	wg.Wait()
}

func (j Job) Fire() {
	wg.Add(1)
	go func() {
		defer wg.Done()
		j.Action(j.Payload)
	}()
	// if waiting for a job to finish, then wait for it to finish, if not, we will just dispatch the job and won't wait it to finish.
	// wg.Wait()
}

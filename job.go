package pool

import "sync"

// Job is a task that should be done in a pool
type Job struct {
	err  error
	work func() error
}

// NewJob creates new Job based on a work function
func NewJob(work func() error) *Job {
	return &Job{work: work}
}

// Run runs a Job and sets it as done in a given sync.WaitGroup
func (j *Job) Run(wg *sync.WaitGroup) {
	j.err = j.work()
	wg.Done()
}

// Err returns an error resulting from the work function
func (j *Job) Err() error {
	return j.err
}

// Package pool implements simple pool of workers
package pool

import (
	"sync"
)

// Pool is a pool of workers that runs jobs concurrently limited by
// given worker count
type Pool struct {
	wg          sync.WaitGroup
	jobs        []*Job
	workerCount int
	jobsChan    chan *Job
}

// NewPool creates new Pool with the given Jobs and the given number of workers
func NewPool(jobs []*Job, workerCount int) *Pool {
	return &Pool{
		jobs:        jobs,
		workerCount: workerCount,
		jobsChan:    make(chan *Job),
	}
}

// Run concurrently runs all jobs in workers and blocks until work is finished
func (p *Pool) Run() {
	for i := 0; i < p.workerCount; i++ {
		go p.work()
	}

	p.wg.Add(len(p.jobs))
	for _, job := range p.jobs {
		p.jobsChan <- job
	}

	// all workers return
	close(p.jobsChan)

	p.wg.Wait()
}

func (p *Pool) work() {
	for job := range p.jobsChan {
		job.Run(&p.wg)
	}
}

// Jobs returns Jobs of this worker pool
func (p *Pool) Jobs() []*Job {
	return p.jobs
}

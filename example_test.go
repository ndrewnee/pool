package pool_test

import (
	"errors"
	"log"

	"github.com/ndrewnee/pool"
)

func ExamplePool() {
	jobs := []*pool.Job{
		pool.NewJob(func() error { return nil }),
		pool.NewJob(func() error { return errors.New("some error") }),
	}

	p := pool.NewPool(jobs, 10)
	p.Run()

	for _, job := range p.Jobs() {
		if err := job.Err(); err != nil {
			log.Println(err)
		}
	}
}

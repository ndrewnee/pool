# pool

## Description

Simple pool of workers

## Installation

```bash
go get github.com/ndrewnee/pool
```

## Usage

```go
package main

import (
    "errors"
    "log"

    "github.com/ndrewnee/pool"
)

func main() {
    jobs := []*pool.Job{
        pool.NewJob(func() error { return nil }),
        pool.NewJob(func() error { return errors.New("some error") }),
    }

    p := pool.NewPool(jobs, 10)
    p.Run()

    for i, job := range p.Jobs() {
        if err := job.Err(); err != nil {
            log.Printf("Job with index %v failed: %v\n", i, err)
            continue
        }

        log.Printf("Job with index %v finished successfully\n", i)
    }
}
```
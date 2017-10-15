# pool

[![GoDoc](https://godoc.org/github.com/ndrewnee/pool?status.svg)](https://godoc.org/github.com/ndrewnee/pool)
[![Go Report Card](https://goreportcard.com/badge/github.com/ndrewnee/pool)](https://goreportcard.com/report/github.com/ndrewnee/pool)
[![Build Status](https://travis-ci.org/ndrewnee/pool.svg?branch=master)](https://travis-ci.org/ndrewnee/pool)
[![Coverage Status](https://coveralls.io/repos/github/ndrewnee/pool/badge.svg)](https://coveralls.io/github/ndrewnee/pool)
[![Contributions Welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/ndrewnee/pool/issues)

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

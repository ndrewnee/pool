package pool_test

import (
	"errors"
	"sync"
	"testing"

	"github.com/ndrewnee/pool"
	"github.com/stretchr/testify/assert"
)

func TestJob_Run(t *testing.T) {
	testErr := errors.New("test error")

	var wg sync.WaitGroup
	wg.Add(1)

	j := pool.NewJob(func() error {
		return testErr
	})

	j.Run(&wg)

	wg.Wait()

	assert.Equal(t, testErr, j.Err())
}

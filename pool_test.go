package pool_test

import (
	"errors"
	"testing"

	"github.com/ndrewnee/pool"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPool_Run(t *testing.T) {
	testErr1 := errors.New("test error 1")
	testErr2 := errors.New("test error 2")

	p := pool.NewPool([]*pool.Job{
		pool.NewJob(func() error { return testErr1 }),
		pool.NewJob(func() error { return testErr2 }),
	}, 1)

	p.Run()

	jobs := p.Jobs()
	require.Len(t, jobs, 2)
	assert.Equal(t, testErr1, jobs[0].Err())
	assert.Equal(t, testErr2, jobs[1].Err())
}

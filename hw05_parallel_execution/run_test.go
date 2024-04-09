package hw05parallelexecution

import (
	"errors"
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"
)

func TestRun(t *testing.T) {
	defer goleak.VerifyNone(t)

	t.Run("if were errors in first M tasks, than finished not more N+M tasks", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32

		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1)))
				atomic.AddInt32(&runTasksCount, 1)
				return err
			})
		}

		workersCount := 10
		maxErrorsCount := 23
		err := Run(tasks, workersCount, maxErrorsCount)

		require.Truef(t, errors.Is(err, ErrErrorsLimitExceeded), "actual err - %v", err)
		require.LessOrEqual(t, runTasksCount, int32(workersCount+maxErrorsCount), "extra tasks were started")
	})

	t.Run("tasks without errors", func(t *testing.T) {
		tasksCount := 50
		tasks := make([]Task, 0, tasksCount)

		var runTasksCount int32
		var sumTime time.Duration

		for i := 0; i < tasksCount; i++ {
			taskSleep := time.Millisecond * time.Duration(rand.Intn(100))
			sumTime += taskSleep

			tasks = append(tasks, func() error {
				time.Sleep(taskSleep)
				atomic.AddInt32(&runTasksCount, 1)
				return nil
			})
		}

		workersCount := 5
		maxErrorsCount := 1

		start := time.Now()
		err := Run(tasks, workersCount, maxErrorsCount)
		elapsedTime := time.Since(start)
		require.NoError(t, err)

		require.Equal(t, runTasksCount, int32(tasksCount), "not all tasks were completed")
		require.LessOrEqual(t, int64(elapsedTime), int64(sumTime/2), "tasks were run sequentially?")
	})
}

func TestTracker(t *testing.T) {
	t.Run("finish_by_max", func(t *testing.T) {
		n := 5
		stop := make(chan bool)
		tracked_channel := make(chan bool, n)
		max := 4
		go tracker(&stop, &tracked_channel, max)
		for i := 0; i < max; i++ {
			tracked_channel <- true
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		}
		<-stop
	})

	t.Run("finish_by_ok", func(t *testing.T) {
		stop := make(chan bool)
		tracked_channel := make(chan bool, 10)
		go tracker(&stop, &tracked_channel, 10)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
		close(tracked_channel)
	})
}

func TestTaskIerator(t *testing.T) {
	t.Run("finish_by_max", func(t *testing.T) {
		n := 10
		task_i := NewTaskIterator(n)
		for i := 0; i < n; i++ {
			_, ok := task_i.Get()
			require.True(t, ok)
		}
		ok := task_i.Close()
		require.True(t, ok)
	})

	t.Run("finish_by_close", func(t *testing.T) {
		n := 10
		task_i := NewTaskIterator(n)
		ok := task_i.Close()
		require.False(t, ok)
		i, ok := task_i.Get()
		require.False(t, ok)
		require.Equal(t, n, i)
	})
}

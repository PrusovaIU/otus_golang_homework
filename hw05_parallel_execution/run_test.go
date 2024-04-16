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
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
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

	t.Run("negative m", func(t *testing.T) {
		tasksCount := 10
		tasks := make([]Task, 0, tasksCount)
		for i := 0; i < tasksCount; i++ {
			err := fmt.Errorf("error from task %d", i)
			tasks = append(tasks, func() error {
				return err
			})
		}
		workersCount := 2
		maxErrorsCount := -1
		err := Run(tasks, workersCount, maxErrorsCount)
		require.NoError(t, err)
	})
}

// func TestHandler(t *testing.T) {
// 	tasksCount := 5
// 	tests := []struct {
// 		name            string
// 		taskReturn      error
// 		errsTrackerRes  int
// 		tasksTrackerRes int
// 		result          bool
// 	}{
// 		{name: "without errs", taskReturn: nil, errsTrackerRes: 0, tasksTrackerRes: tasksCount, result: true},
// 		// {name: "with errs", taskReturn: errors.New("Test"), errsTrackerRes: tasksCount, tasksTrackerRes: 0, result: false},
// 	}

// 	// for _, tc := range tests {
// 	// 	tc := tc
// 	// 	t.Run(tc.name, func(t *testing.T) {
// 	// 		stop := sync.WaitGroup{}
// 	// 		stop.Add(1)
// 	// 		iTask := NewIterator(tasksCount)
// 	// 		iErr := NewIterator(tasksCount)
// 	// 		tasks := make([]Task, 0, tasksCount)
// 	// 		for i := 0; i < tasksCount; i++ {
// 	// 			tasks = append(tasks, func() error {
// 	// 				return tc.taskReturn
// 	// 			})
// 	// 		}
// 	// 		handler(tasks, &stop, &iTask, &iErr)
// 	// 		stop.Wait()
// 	// 	})
// 	// }
// 	for _, tc := range tests {
// 		tc := tc
// 		t.Run(tc.name, func(t *testing.T) {
// 			stop := sync.WaitGroup{}
// 			stop.Add(1)
// 			iErr := NewIterator(tasksCount)
// 			tasks := make(chan Task, tasksCount)
// 			for i := 0; i < tasksCount; i++ {
// 				tasks <- func() error {
// 					return tc.taskReturn
// 				}
// 			}
// 			go handler(tasks, &stop, &iErr)
// 			stop.Wait()
// 			close(tasks)
// 		})
// 	}
// }

func TestHandler(t *testing.T) {
	tasks := make(chan Task, 2)
	tasks <- func() error {
		return nil
	}
	tasks <- func() error {
		return errors.New("Test")
	}
	iTasks := make(chan bool)
	iErrs := make(chan bool)
	go handler(tasks, iTasks, iErrs)
	<-iTasks
	<-iErrs
	close(tasks)
	close(iTasks)
	close(iErrs)
}

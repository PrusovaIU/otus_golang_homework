package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func handler(tasks <-chan Task, results chan<- error, done *sync.WaitGroup) {
	for {
		task, ok := <-tasks
		if !ok {
			break
		}
		results <- task()
	}
	done.Done()
}

func observe(tasksCount int, results <-chan error, n, errsMax int, tasksChan chan Task, tasks []Task) int {
	iErr := 0
	for i := n; i < tasksCount; i++ {
		result := <-results
		if result != nil {
			iErr++
		}
		if iErr > errsMax {
			break
		}
		tasksChan <- tasks[i]
	}
	close(tasksChan)
	return iErr
}

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	tasksCount := len(tasks)
	if m < 0 {
		m = tasksCount
	}
	tasksChan := make(chan Task, n)
	results := make(chan error, n)
	done := sync.WaitGroup{}
	for i := 0; i < n; i++ {
		done.Add(1)
		tasksChan <- tasks[i]
		go handler(tasksChan, results, &done)
	}
	iErr := observe(tasksCount, results, n, m, tasksChan, tasks)
	done.Wait()
	if iErr > m {
		return ErrErrorsLimitExceeded
	}
	return nil
}

package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func handler(tasks []Task, stop *sync.WaitGroup, iTask *Iterator, iErr *Iterator) {
	for iErr.Check() {
		j, ok := iTask.Get()
		if !ok {
			break
		} else {
			task := tasks[j]
			err := task()
			if err != nil {
				iErr.Get()
			}
		}
	}
	stop.Done()
}

func run(tasks []Task, n, m, tasksCount int) bool {
	stop := sync.WaitGroup{}
	iTask := NewIterator(tasksCount)
	iErr := NewIterator(m)
	for i := 0; i < n; i++ {
		stop.Add(1)
		go handler(tasks, &stop, &iTask, &iErr)
	}
	stop.Wait()
	return iErr.Check()
}

func Run(tasks []Task, n, m int) error {
	tasksCount := len(tasks)
	if tasksCount < n {
		n = tasksCount
	}
	if m < 0 {
		m = tasksCount
	}
	if ok := run(tasks, n, m, tasksCount); ok {
		return nil
	} else {
		return ErrErrorsLimitExceeded
	}
}

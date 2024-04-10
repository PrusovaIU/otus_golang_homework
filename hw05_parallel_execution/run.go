package hw05parallelexecution

import (
	"errors"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func tracker(stop *chan bool, tracked_channel *chan bool, max int) {
	i := 0
	for ok := true; i < max && ok; i++ {
		_, ok = <-*tracked_channel
	}
	if i == max {
		*stop <- true
	}
}

func handler(tasks []Task, i *TaskIterator, errsTracker *chan bool, taskTracker *chan bool) {
	for {
		j, ok := i.Get()
		if !ok {
			break
		} else {
			task := tasks[j]
			err := task()
			if err != nil {
				*errsTracker <- true
			} else {
				*taskTracker <- true
			}
		}
	}
}

func run(tasks []Task, n, m, tasksCount int) bool {
	taskTracker := make(chan bool, len(tasks))
	errsTracker := make(chan bool, m)
	stop := make(chan bool)
	go tracker(&stop, &taskTracker, tasksCount)
	go tracker(&stop, &errsTracker, m)
	iTask := NewTaskIterator(tasksCount)
	for i := 0; i < n; i++ {
		go handler(tasks, &iTask, &errsTracker, &taskTracker)
	}
	<-stop
	ok := iTask.Close()
	close(errsTracker)
	close(taskTracker)
	return ok
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

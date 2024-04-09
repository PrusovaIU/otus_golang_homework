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

func handler(tasks []Task, i *TaskIterator, errs_tracker *chan bool) {
	for {
		j, ok := i.Get()
		if !ok {
			break
		} else {
			task := tasks[j]
			err := task()
			if err != nil {
				*errs_tracker <- true
			}
		}
	}
}

func Run(tasks []Task, n, m int) error {
	tasks_count := len(tasks)
	if tasks_count < n {
		n = tasks_count
	}
	if m < 0 {
		m = tasks_count
	}
	task_tracker := make(chan bool, len(tasks))
	errs_tracker := make(chan bool, m)
	stop := make(chan bool)
	go tracker(&stop, &task_tracker, tasks_count)
	go tracker(&stop, &errs_tracker, m)
	i_task := NewTaskIterator(tasks_count)
	for i := 0; i < n; i++ {
		go handler(tasks, &i_task, &errs_tracker)
	}
	<-stop
	ok := i_task.Close()
	close(errs_tracker)
	close(task_tracker)
	if ok {
		return nil
	} else {
		return ErrErrorsLimitExceeded
	}
}

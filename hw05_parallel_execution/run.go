package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

func handler(tasks chan Task, iTasks chan bool, iErrs chan bool) {
	for {
		task, ok := <-tasks
		if !ok {
			break
		}
		err := task()
		if err != nil {
			iErrs <- true
		} else {
			iTasks <- true
		}
	}
}

func observer(i chan bool, max int, stop *sync.WaitGroup, j *Counter) {
	for {
		if _, ok := <-i; !ok {
			break
		} else {
			j.Inc()
		}
	}
	stop.Done()
}

func run(tasks []Task, n, m, tasksCount int) bool {
	var tasksChan = make(chan Task, tasksCount)
	for _, task := range tasks {
		tasksChan <- task
	}
	stop := sync.WaitGroup{}
	stop.Add(1)
	iTasks := make(chan bool, tasksCount)
	jTask := NewCounter()
	iErrs := make(chan bool, m)
	jErr := NewCounter()
	go observer(iTasks, tasksCount, &stop, &jTask)
	go observer(iErrs, m, &stop, &jErr)
	for i := 0; i < n; i++ {
		go handler(tasksChan, iTasks, iErrs)
	}
	stop.Wait()
	close(tasksChan)
	return jErr.Get() < m
}

func Run(tasks []Task, n, m int) error {
	tasksCount := len(tasks)
	if m < 0 {
		m = tasksCount + 1
	}
	if ok := run(tasks, n, m, tasksCount); ok {
		return nil
	}
	return ErrErrorsLimitExceeded
}

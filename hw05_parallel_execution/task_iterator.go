package hw05parallelexecution

import (
	"sync"
)

type TaskIterator struct {
	mu  sync.Mutex
	i   int
	max int
}

func (i *TaskIterator) check(j int) bool {
	var ok bool
	if j < i.max {
		ok = true
	} else {
		ok = false
	}
	return ok
}

func (i *TaskIterator) Get() (int, bool) {
	i.mu.Lock()
	j := i.i
	ok := i.check(j)
	if ok {
		i.i++
	}
	i.mu.Unlock()
	return j, ok
}

func (i *TaskIterator) Close() bool {
	i.mu.Lock()
	j := i.i
	i.i = i.max
	i.mu.Unlock()
	return !i.check(j)
}

func NewTaskIterator(max int) TaskIterator {
	return TaskIterator{sync.Mutex{}, 0, max}
}

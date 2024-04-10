package hw05parallelexecution

import (
	"sync"
)

type Iterator struct {
	mu  sync.Mutex
	i   int
	max int
}

func (i *Iterator) check(j int) bool {
	var ok bool
	if j < i.max {
		ok = true
	} else {
		ok = false
	}
	return ok
}

func (i *Iterator) Check() bool {
	i.mu.Lock()
	j := i.i
	i.mu.Unlock()
	return i.check(j)
}

func (i *Iterator) Get() (int, bool) {
	i.mu.Lock()
	j := i.i
	ok := i.check(j)
	if ok {
		i.i++
	}
	i.mu.Unlock()
	return j, ok
}

func NewIterator(max int) Iterator {
	return Iterator{sync.Mutex{}, 0, max}
}

package hw05parallelexecution

import (
	"sync"
)

type Counter struct {
	mu sync.Mutex
	i  int
}

func (i *Counter) Inc() {
	i.mu.Lock()
	i.i++
	i.mu.Unlock()
}

func (i *Counter) Get() int {
	i.mu.Lock()
	j := i.i
	i.mu.Unlock()
	return j
}

func NewCounter() Counter {
	return Counter{sync.Mutex{}, 0}
}

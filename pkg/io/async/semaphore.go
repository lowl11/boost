package async

import "github.com/lowl11/boost/data/interfaces"

type Semaphore struct {
	c chan struct{}
}

func (semaphore *Semaphore) Acquire() {
	semaphore.c <- struct{}{}
}

func (semaphore *Semaphore) Release() {
	<-semaphore.c
}

func NewSemaphore(size int) interfaces.Semaphore {
	return &Semaphore{
		c: make(chan struct{}, size),
	}
}

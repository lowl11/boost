package async

import "github.com/lowl11/boost/data/interfaces"

type Semaphore struct {
	C chan struct{}
}

func (semaphore *Semaphore) Acquire() {
	semaphore.C <- struct{}{}
}

func (semaphore *Semaphore) Release() {
	<-semaphore.C
}

func NewSemaphore(size int) interfaces.Semaphore {
	return &Semaphore{
		C: make(chan struct{}, size),
	}
}

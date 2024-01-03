package async

import "github.com/lowl11/boost/data/interfaces"

type semaphore struct {
	c chan struct{}
}

func (s *semaphore) Acquire() {
	s.c <- struct{}{}
}

func (s *semaphore) Release() {
	<-s.c
}

func (s *semaphore) Close() {
	close(s.c)
}

func NewSemaphore(size int) interfaces.Semaphore {
	return &semaphore{
		c: make(chan struct{}, size),
	}
}

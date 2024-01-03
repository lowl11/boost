package interfaces

type Semaphore interface {
	Acquire()
	Release()
}

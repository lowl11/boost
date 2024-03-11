package cancel

import "sync"

type Cancel struct {
	wg *sync.WaitGroup
}

var instance *Cancel

func Get() *Cancel {
	if instance != nil {
		return instance
	}

	instance = &Cancel{
		wg: &sync.WaitGroup{},
	}
	return instance
}

func (c *Cancel) Add() {
	c.wg.Add(1)
}

func (c *Cancel) Done() {
	c.wg.Done()
}

func (c *Cancel) Wait() {
	c.wg.Wait()
}

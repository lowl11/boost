package di_container

import "reflect"

type tQueue struct {
	queues map[reflect.Type]*queue
}

func newTQueue() *tQueue {
	return &tQueue{
		queues: make(map[reflect.Type]*queue),
	}
}

func (tq *tQueue) Enqueue(value any) *tQueue {
	valueType := reflect.TypeOf(value)
	_, ok := tq.queues[valueType]
	if !ok {
		tq.queues[valueType] = newQueue(valueType)
	}

	tq.queues[valueType].Enqueue(value)
	return tq
}

func (tq *tQueue) Dequeue(t reflect.Type) any {
	typeQueue, exist := tq.queues[t]
	if !exist {
		return nil
	}

	return typeQueue.Dequeue()
}

func (tq *tQueue) Len() int {
	var length int
	for _, q := range tq.queues {
		length += q.Len()
	}
	return length
}

func (tq *tQueue) Copy() *tQueue {
	cp := newTQueue()
	for _, q := range tq.queues {
		cp.queues[q.DT()] = q.Copy()
	}
	return cp
}

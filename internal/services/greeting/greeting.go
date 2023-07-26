package greeting

import "github.com/lowl11/boost/internal/services/counter"

type Greeting struct {
	message string
	counter *counter.Counter

	mainColor string
}

func New(counter *counter.Counter) *Greeting {
	return &Greeting{
		counter: counter,
	}
}

package greeting

import (
	"github.com/lowl11/boost/internal/services/counter"
	"github.com/lowl11/boost/pkg/enums/colors"
)

type Greeting struct {
	message string
	counter *counter.Counter

	mainColor     string
	specificColor string
}

func New(counter *counter.Counter) *Greeting {
	return &Greeting{
		counter: counter,

		mainColor:     colors.Gray,
		specificColor: colors.Cyan,
	}
}

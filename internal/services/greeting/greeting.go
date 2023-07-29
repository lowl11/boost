package greeting

import (
	"github.com/lowl11/boost/internal/services/counter"
	"github.com/lowl11/boost/pkg/enums/colors"
)

type Context struct {
	Mode string
	Port string
}

type Greeting struct {
	message string
	counter *counter.Counter
	ctx     Context

	mainColor     string
	specificColor string

	printed bool
}

func New(counter *counter.Counter, ctx Context) *Greeting {
	return &Greeting{
		counter: counter,
		ctx:     ctx,

		mainColor:     colors.Gray,
		specificColor: colors.Cyan,
	}
}

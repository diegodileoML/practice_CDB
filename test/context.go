
package test

import (
	"time"
)

type FakeContext struct {
	ValueData interface{}
}

func (f FakeContext) Deadline() (deadline time.Time, ok bool) {
	return
}

func (f FakeContext) Done() <-chan struct{} {
	return nil
}

func (f FakeContext) Err() error {
	return nil
}

func (f FakeContext) Value(key interface{}) (i interface{}) {
	switch key {
	case "metrics":
		// this is just for test coverage of bits of code about tracing
		return FakeTracing{}
	default:
		return f.ValueData
	}
}
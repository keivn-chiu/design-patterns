package state

import "testing"

func TestDemo(t *testing.T) {
	ctx := NewCtx()
	f := func() {
		ctx.Current()
		ctx.Next()
	}
	for i := 0; i < 4; i++ {
		f()
	}
}

// idle -> choose -> pay -> idle -> choose

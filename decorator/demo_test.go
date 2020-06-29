package decorator

import "testing"

func TestDemo(t *testing.T) {
	// get base type obj
	c := NewComponent()
	// warpAdd
	c = NewWarpAddComponent(c, 10)
	// warpMul
	c = NewWarpMulComponent(c, 8)
	// get calc result
	result := c.Calc()
	if result != 80 {
		t.Fail()
	}
}

package interpreter

import "testing"

func TestDemo(t *testing.T) {
	parser := NewParser("1 + 1 - 1 + 11")
	ret := parser.Result().Interpret()
	if ret != 12 {
		t.Error("result != 12", "ret = ", ret)
	}
}

package iterator

import "testing"

func TestDemo(t *testing.T) {
	var aggregate Aggregate
	aggregate = NewNumber(2, 10)
	iterator := aggregate.Iterator()
	Print(iterator)
}

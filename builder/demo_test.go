package builder

import "testing"

func TestDemo(t *testing.T) {
	b := &Builder{}
	director := NewDirector(b)
	director.Construct()
	if b.result != "abc" {
		t.Errorf("Expect: %s, Actual: %s", "abc", b.result)
	}
}

package observer

import "testing"

func TestDemo(t *testing.T) {
	subject := NewSubject()

	kevin := NewReader("kevin")
	mary := NewReader("mary")
	peter := NewReader("peter")

	subject.Attach(kevin)
	subject.Attach(mary)
	subject.Attach(peter)

	subject.notify("Hello guys!")
}

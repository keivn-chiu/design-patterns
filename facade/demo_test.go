package facade

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	api := NewApi()
	fmt.Println(api.Do())
	fmt.Println(api.Done())
}

/*
=== RUN   TestDemo
A do-B do
A done-B done
--- PASS: TestDemo (0.00s)
PASS

Process finished with exit code 0
*/
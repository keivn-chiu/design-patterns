package proxy

import "testing"

func TestDemo(t *testing.T) {
	// no proxy
	do := &ImplDo{}
	do.Do()
	println("-----------")
	// to use proxy
	proxy := &IDoProxy{
		impl: *do,
	}
	proxy.Do()
}

/*
=== RUN   TestDemo
Process task
-----------
Before task
Process task
After task
--- PASS: TestDemo (0.00s)
PASS
*/

package adapter

import "testing"

func TestDemo(t *testing.T) {
	// mount usbB directly
	b := NewUsbB()
	MountUsb(b)
	// mount usbA via an adapter
	a := NewUsbA()
	adapter := NewUsbAdapter(a)
	MountUsb(adapter)
}

// MountUsb func for test adapter
func MountUsb(b IUsbB) {
	println(b.MountB())
}

/*
=== RUN   TestDemo
B
A
--- PASS: TestDemo (0.00s)
PASS
*/

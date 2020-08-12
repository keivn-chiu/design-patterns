package visitor

import "testing"

func TestDemo(t *testing.T) {
	collector := NewCustomerCollector()
	collector.Add(&CustomerOne{})
	collector.Add(&CustomerTwo{})

	collector.Accept(&VTwo{})
	/*
		vTwo - customer two visit
	*/
	collector.Accept(&VOne{})
	/*
		vOne - customer one visit
		vOne - customer two visit
	*/

}

package flyweight

import "testing"

func TestDemo(t *testing.T) {
	v1 := NewDataViewer("kevin.data")
	v2 := NewDataViewer("kevin.data")
	v1.ReadData()
	v2.ReadData()
	if v1.IData != v2.IData {
		t.Fail()
	}

	factory := GetFactory()
	data := factory.GetData("kevin.data")
	if v1.IData != data {
		t.Fail()
	}
}

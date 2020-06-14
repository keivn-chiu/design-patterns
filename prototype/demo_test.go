package prototype

import (
	"testing"
)

var manager *Manager

func TestDemo(t *testing.T) {
	// check if student in manager
	sp := manager.GetPrototype("student")
	obj, ok := sp.(*CloneObj)
	if !ok {
		t.Fatal("Type incorrect")
		return
	}
	// clone
	clone := obj.Clone()
	if obj == clone {
		t.Fatal("Clone failed")
		return
	}
	// check clone obj parameters
	newObj, ok := clone.(*CloneObj)
	if !ok {
		t.Fatal("Clone Type incorrect")
		return
	}
	if newObj.Age != obj.Age || newObj.Name != obj.Name {
		t.Fatal("Clone failed")
	}
}

func init() {
	// Get an instance of Prototype
	manager = NewManager()
	manager.AddPrototype("student", &CloneObj{
		Name: "Peter",
		Age:  19,
	})
}

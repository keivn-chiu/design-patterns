package composite

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	root := NewComponent(NODE, "root")
	node1 := NewComponent(NODE, "node-1")
	root.AddChild(node1)
	node2 := NewComponent(NODE, "node-2")
	root.AddChild(node2)
	leaf1 := NewComponent(LEAF, "leaf-1")
	root.AddChild(leaf1)
	leaf2 := NewComponent(LEAF, "leaf-2")
	node2.AddChild(leaf2)

	if n := root.GetChild("node-1"); n != nil && n.Name() == "node-1" {
		fmt.Println("node-1 in root")
	}
	if l := root.GetChild("leaf-1"); l != nil && l.Name() == "leaf-1" {
		fmt.Println("leaf-1 in root")
	}
	if n := root.GetChild("node-2"); n != nil && n.Name() == "node-2" {
		fmt.Println("node-2 in root")
		if l := n.GetChild("leaf-2"); l != nil && l.Name() == "leaf-2" {
			fmt.Println("leaf-2 in node-2")
		}
	}

	if root.GetChild("leaf-2") == nil {
		fmt.Println("leaf-2 is not under root")
	}

}

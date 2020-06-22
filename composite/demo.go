package composite

import "fmt"

// composite has a root node to save all notes or leaves

// IComponent the base api for base type component, type node or type leaf
type IComponent interface {
	Name() string
	SetName(name string)
	AddChild(ic IComponent)
	RemoveChild(name string)
	GetChild(name string) IComponent
	Parent() IComponent
	SetParent(c IComponent)
	Print()
}

// TYPE node or leaf
type TYPE int

const (
	NODE TYPE = iota
	LEAF
)

// NewComponent to get a component according to TYPE
func NewComponent(t TYPE, name string) IComponent {
	var comp IComponent
	switch t {
	case LEAF:
		comp = &leaf{}
	case NODE:
		comp = &node{
			children: make(map[string]IComponent),
		}
	}
	if comp != nil {
		comp.SetName(name)
	}
	return comp
}

// component the base struct for type node or type leaf
type component struct {
	name   string
	parent IComponent
}

func (c *component) Name() string {
	return c.name
}
func (c *component) SetName(name string) {
	c.name = name
}
func (c *component) AddChild(ic IComponent) {

}
func (c *component) RemoveChild(name string) {

}
func (c *component) GetChild(name string) IComponent {
	return nil
}
func (c *component) Parent() IComponent {
	return c.parent
}
func (c *component) SetParent(ic IComponent) {
	c.parent = ic
}
func (c *component) Print() {
	fmt.Printf("base print{name:%s}\n", c.name)
}

// node the base struct as a node, nodes can include some nodes or leaves
type node struct {
	component
	children map[string]IComponent
}

func (n *node) AddChild(c IComponent) {
	c.SetParent(n)
	n.children[c.Name()] = c
}

func (n *node) RemoveChild(name string) {
	delete(n.children, name)
}

func (n *node) GetChild(name string) IComponent {
	return n.children[name]
}

func (n *node) Print() {
	fmt.Println(n.name)
	fmt.Printf(" %s children:\n", n.name)
	for _, v := range n.children {
		v.Print()
	}
	fmt.Println()
}

// leaf the base struct ad a leaf, it just can be a child for nodes
type leaf struct {
	component
}

func (l *leaf) Print() {
	fmt.Println("leaf:", l.name)
}

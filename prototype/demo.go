package prototype

// cloneable interface, when need to create lots of instances of designated
type Cloneable interface {
	Clone() Cloneable
}

// manager to save all prototypes
type Manager struct {
	prototypes map[string]Cloneable
}

// NewManager to get a instance of prototypes manager
func NewManager() *Manager {
	return &Manager{
		prototypes: make(map[string]Cloneable),
	}
}

// AddPrototype to add a prototype into manager
func (manager *Manager) AddPrototype(name string, prototype Cloneable) {
	manager.prototypes[name] = prototype
}

// RemovePrototype to remove a prototype from manager
func (manager *Manager) RemovePrototype(name string) {
	delete(manager.prototypes, name)
}

// GetPrototype to get a prototype from manager
func (manager *Manager) GetPrototype(name string) Cloneable {
	return manager.prototypes[name]
}

// a type implement Cloneable
type CloneObj struct {
	Name string
	Age  int
}

func (c *CloneObj) Clone() Cloneable {
	newObj := *c
	return &newObj
}

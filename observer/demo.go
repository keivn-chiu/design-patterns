package observer

import "fmt"

// IObserver the interface for observers
type IObserver interface {
	update(data interface{})
	Name() string
}

// IObservable the interface for observable type
type IObservable interface {
	Attach(IObserver)
	notify(data interface{})
	Remove(string)
}

// Subject the implement of IObservable
type Subject struct {
	observers map[string]IObserver
	context   string
}

func NewSubject() *Subject {
	return &Subject{
		observers: make(map[string]IObserver),
		context:   "default",
	}
}
func (s *Subject) Attach(o IObserver) {
	s.observers[o.Name()] = o
}
func (s *Subject) notify(data interface{}) {
	for _, observer := range s.observers {
		observer.update(data)
	}
}
func (s *Subject) Remove(name string) {
	delete(s.observers, name)
}

// Reader the one of implement of IObserver
type Reader struct {
	name string
}

func NewReader(name string) *Reader {
	return &Reader{name: name}
}
func (r *Reader) update(data interface{}) {
	fmt.Printf("%s got: %v\n", r.name, data)
}
func (r *Reader) Name() string {
	return r.name
}

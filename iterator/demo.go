package iterator

import "fmt"

// iterator it can help for traversing such as list, map

// Aggregate interface for the type has iterator
type Aggregate interface {
	Iterator() Iterator
}

// Iterator the interface of iterator
type Iterator interface {
	Next() interface{}
	HasNext() bool
	First()
}

// Number the implement
type Number struct {
	start int
	end   int
}

func (n *Number) Iterator() Iterator {
	return &NumberIterator{next: n.start, Number: n}
}
func NewNumber(start, end int) Aggregate {
	return &Number{start: start, end: end}
}

// NumberIterator the iterator of Number type
type NumberIterator struct {
	*Number
	next int
}

func (ni *NumberIterator) Next() interface{} {
	if ni.HasNext() {
		next := ni.next
		ni.next++
		return next
	}
	return nil
}
func (ni *NumberIterator) HasNext() bool {
	return ni.next<=ni.end
}
func (ni *NumberIterator) First() {
	ni.next = ni.start
}

// Print to print out to stdin
func Print(i Iterator) {
	for i.First(); i.HasNext(); {
		fmt.Printf("%#v\n", i.Next())
	}
}

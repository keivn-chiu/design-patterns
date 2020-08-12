package visitor

import "fmt"

// visitor mode helps to separate data structure and data operation

type IVisitor interface {
	Visit(Customer)
}

type Customer interface {
	Accept(IVisitor)
}

type CustomerCollector struct {
	customers []Customer
}

func NewCustomerCollector() *CustomerCollector {
	return &CustomerCollector{make([]Customer, 0)}
}

func (cc *CustomerCollector) Accept(visitor IVisitor) {
	for _, customer := range cc.customers {
		customer.Accept(visitor)
	}
}
func (cc *CustomerCollector) Add(customer Customer) {
	cc.customers = append(cc.customers, customer)
}

type CustomerOne struct {
}

func (co *CustomerOne) Accept(visitor IVisitor) {
	visitor.Visit(co)
}

type CustomerTwo struct {
}

func (ct *CustomerTwo) Accept(visitor IVisitor) {
	visitor.Visit(ct)
}

type VOne struct {
}

func (v *VOne) Visit(customer Customer) {
	switch customer.(type) {
	case *CustomerOne:
		fmt.Println("vOne - customer one visit")
	case *CustomerTwo:
		fmt.Println("vOne - customer two visit")
	}
}

type VTwo struct {
}

func (v *VTwo) Visit(customer Customer) {
	switch customer.(type) {
	case *CustomerTwo:
		fmt.Println("vTwo - customer two visit")
	}
}

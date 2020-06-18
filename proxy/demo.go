package proxy

import "fmt"

// proxy to do something before or after actual jobs

// IDoSomething the interface for proxy and the obj types need proxy
type IDoSomething interface {
	Do()
}

// ImplDo the obj type need proxy
type ImplDo struct {
}

func (*ImplDo) Do() {
	fmt.Println("Process task")
}

// IDoProxy the proxy for ImplDo
type IDoProxy struct {
	impl ImplDo
}

func (p *IDoProxy) Do() {
	fmt.Println("Before task")
	p.impl.Do()
	fmt.Println("After task")
}

package chain_of_responsibility

import "fmt"

// this mode helps to separate different responsibilities and combine related responsibilities dynamically.
// 1. chain need to include current responsibility obj and next chain
// 2. responsibility obj need to provide interface for checking if it can handle the related request
// 3. responsibility obj need to provide handling function for handling related requests

type Manager interface {
	CanHandle(level int) bool
	HandleRequest(name string, level int) bool
}

// RequestChain the base request chain
type RequestChain struct {
	Manager
	successor *RequestChain
}

func (r *RequestChain) SetSuccessor(s *RequestChain) {
	r.successor = s
}

func (r *RequestChain) HandleRequest(name string, level int) bool {
	if r.Manager.CanHandle(level) {
		return r.Manager.HandleRequest(name, level)
	}
	if r.successor != nil {
		return r.successor.HandleRequest(name, level)
	}
	return false
}

func (r *RequestChain) CanHandle(level int) bool {
	return true
}

// RequestChain children
type RequestLevel1 struct {
}

func (r *RequestLevel1) CanHandle(level int) bool {
	return level <= 1
}
func (r *RequestLevel1) HandleRequest(name string, level int) bool {
	if name == "peter" {
		fmt.Printf("%s is level %d, job done\n", name, level)
		return true
	}
	fmt.Printf("%s can't handle the request, level is %d\n", name, level)
	return false
}
func NewRequestLevel1() *RequestChain {
	return &RequestChain{Manager: &RequestLevel1{}}
}

type RequestLevel2 struct {
}

func (r *RequestLevel2) CanHandle(level int) bool {
	return level <= 2
}
func (r *RequestLevel2) HandleRequest(name string, level int) bool {
	if name == "pop" {
		fmt.Printf("%s is level %d, job done\n", name, level)
		return true
	}
	fmt.Printf("%s can't handle the request, level is %d\n", name, level)
	return false
}
func NewRequestLevel2() *RequestChain {
	return &RequestChain{Manager: &RequestLevel2{}}
}

type RequestLevel3 struct {
}

func (r *RequestLevel3) CanHandle(level int) bool {
	return level <= 3
}
func (r *RequestLevel3) HandleRequest(name string, level int) bool {
	if name == "may" {
		fmt.Printf("%s is level %d, job done\n", name, level)
		return true
	}
	fmt.Printf("%s can't handle the request, level is %d\n", name, level)
	return false
}
func NewRequestLevel3() *RequestChain {
	return &RequestChain{Manager: &RequestLevel3{}}
}

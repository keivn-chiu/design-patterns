package state

import "fmt"

// to separate states and behaviors

// interface of state
type IState interface {
	Current()
	Next(ctx *Ctx)
}

// the top state struct
type Ctx struct {
	IState
}

func (c *Ctx) Current() {
	c.IState.Current()
}
func (c *Ctx) Next() {
	c.IState.Next(c)
}
func NewCtx() *Ctx {
	return &Ctx{&Idle{}}
}

// states, idle -> choose -> pay -> idle
type Choose struct {
}

func (*Choose) Current() {
	fmt.Println("choose state -> pay state")
}
func (*Choose) Next(ctx *Ctx) {
	ctx.IState = &Pay{}
}

type Pay struct {
}

func (*Pay) Current() {
	fmt.Println("pay state -> idle state")
}
func (*Pay) Next(ctx *Ctx) {
	ctx.IState = &Idle{}
}

type Idle struct {
}

func (*Idle) Current() {
	fmt.Println("idle state -> choose state")
}
func (*Idle) Next(ctx *Ctx) {
	ctx.IState = &Choose{}
}

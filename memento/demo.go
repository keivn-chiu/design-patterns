package memento

import "fmt"

// to save and output the internal state of program

// the interface of memento
type Memento interface {
	State() string
}

type Game struct {
	hp, mp int
}

func (g *Game) Save() Memento {
	return &gameMemento{hp: g.hp, mp: g.mp}
}
func (g *Game) Load(m Memento) {
	l, ok := m.(*gameMemento)
	if ok {
		g.mp = l.mp
		g.hp = l.hp
	}
}
func (g *Game) State() {
	fmt.Printf("game state: hp=%d,mp=%d\n", g.hp, g.mp)
}
func (g *Game) Change(hp, mp int) {
	g.hp = hp
	g.mp = mp
}

// the memento of Game
type gameMemento struct {
	hp, mp int
}

func (gm *gameMemento) State() string {
	return fmt.Sprintf("hp=%d,mp=%d", gm.hp, gm.mp)
}

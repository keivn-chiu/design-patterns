package memento

import "testing"

func TestDemo(t *testing.T) {
	game := &Game{100, 100}
	temp := game.Save()
	game.State()
	game.Change(200, 300)
	game.State()
	println("load ...")
	game.Load(temp)
	game.State()
}

/*
game state: hp=100,mp=100
game state: hp=200,mp=300
load ...
game state: hp=100,mp=100
*/

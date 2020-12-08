package memento

import "testing"

func TestGameMemento(t *testing.T)  {
	game := new(Game)
	game.hp = 10
	game.mp = 10

	game.Status()
	progress := game.Save()

	game.Play(-1, -2)
	game.Status()

	game.Back(progress)
	game.Status()
}
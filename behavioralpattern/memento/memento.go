package memento

import "fmt"

type Memento interface {}

// 备忘录类
type gameMemento struct {
	hp int
	mp int
}

// 发起人类
type Game struct {
	hp int
	mp int
}

func (g *Game) Play(mpDelta, hpDelta int)  {
	g.hp += hpDelta
	g.mp += mpDelta
}

func (g *Game) Save() Memento  {
	return &gameMemento{
		hp: g.hp,
		mp: g.mp,
	}
}

func (g *Game) Back(m Memento)  {
	gm := m.(*gameMemento)
	g.hp = gm.hp
	g.mp = gm.mp
}

func (g *Game) Status() {
	fmt.Printf("current HP:%d, MP:%d\n", g.hp, g.mp)
}



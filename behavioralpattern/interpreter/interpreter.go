package interpreter

import (
	"strconv"
	"strings"
)

type Node interface {
	Interpret() int
}

type ValNode struct {
	val int
}

func (n *ValNode) Interpret()int {
	return n.val
}

type AddNode struct {
	left Node
	right Node
}

func (a *AddNode) Interpret()int {
	return a.left.Interpret() + a.right.Interpret()
}

type MinNode struct {
	left Node
	right Node
}

func (m *MinNode) Interpret()int {
	return m.left.Interpret() - m.right.Interpret()
}

type Parser struct {
	exp []string
	index int
	prev Node
}

func (p *Parser) newValNode() Node{
	v, _ := strconv.Atoi(p.exp[p.index])
	p.index ++
	return &ValNode{val:v}
}

func (p *Parser) newAddNode() Node {
	p.index ++
	return &AddNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) newMinNode() Node {
	p.index ++
	return &MinNode{
		left:  p.prev,
		right: p.newValNode(),
	}
}

func (p *Parser) Parse(exp string) {
	p.exp = strings.Split(exp, " ")

	for {
		if p.index >= len(p.exp) {
			return
		}

		switch p.exp[p.index] {
		case "+":
			p.prev = p.newAddNode()
		case "-":
			p.prev = p.newMinNode()
		default:
			p.prev = p.newValNode()
		}
	}
}

func (p *Parser) Result()Node {
	return p.prev
}


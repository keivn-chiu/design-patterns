package interpreter

import (
	"strconv"
	"strings"
)

// for parsing some sentences via some interpreters, such as sql sentence parsing
type INode interface {
	Interpret() int
}

type ValNode struct {
	value int
}

func (v *ValNode) Interpret() int {
	return v.value
}

type AddNode struct {
	pre, next INode
}

func (a *AddNode) Interpret() int {
	return a.pre.Interpret() + a.next.Interpret()
}

type MinNode struct {
	pre, next INode
}

func (m *MinNode) Interpret() int {
	return m.pre.Interpret() - m.next.Interpret()
}

// Parser to parse a sentence
type Parser struct {
	index   int
	preNode INode
	exp     []string
}

func (p *Parser) Interpret() int {
	return p.preNode.Interpret()
}
func NewParser(sentence string) *Parser {
	return &Parser{exp: strings.Split(sentence, " ")}
}
func (p *Parser) Result() INode {
	for {
		if p.index >= len(p.exp) {
			return p.preNode
		}
		switch p.exp[p.index] {
		case "+":
			p.preNode = p.newAddNode()
		case "-":
			p.preNode = p.newMinNode()
		default:
			p.preNode = p.newValNode()
		}
	}
}
func (p *Parser) newValNode() INode {
	num, _ := strconv.Atoi(p.exp[p.index])
	p.index++
	return &ValNode{num}
}
func (p *Parser) newAddNode() INode {
	p.index++
	return &AddNode{pre: p.preNode, next: p.newValNode()}
}
func (p *Parser) newMinNode() INode {
	p.index++
	return &MinNode{pre: p.preNode, next: p.newValNode()}
}

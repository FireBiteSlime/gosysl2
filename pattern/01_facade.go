package pattern

import (
	"strings"
)

func NewMan() *Man {
	return &Man{
		house: &House{},
		tree:  &Tree{},
		child: &Child{},
	}
}

type Man struct {
	house *House
	tree  *Tree
	child *Child
}

func (m *Man) Todo() string {
	result := []string{
		m.house.Build(),
		m.tree.Grow(),
		m.child.Born(),
	}
	return strings.Join(result, "\n")
}

type House struct {
}

func (h *House) Build() string {
	return "Build house"
}

type Tree struct {
}

func (t *Tree) Grow() string {
	return "Tree grow"
}

type Child struct {
}

func (c *Child) Born() string {
	return "Child born"
}

package main

//
// The Go way SOLUTION
// ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~ ~~~~
// This solution is leveraging Go modern principles and implements
// Prototype based on that.
//

type TreeComponent interface {
	Clone() TreeComponent
}

type Node struct {
	parent   *Node
	children []*Node
}

func (n Node) Clone() TreeComponent {
	return &Node{
		parent:   n.parent,
		children: n.children,
	}
}

type Edge struct {
	node1 Node
	node2 Node
}

func (e Edge) Clone() TreeComponent {
	return &Edge{
		node1: e.node1,
		node2: e.node2,
	}
}

package main

type Node struct {
	left  *Node
	right *Node
	value int
}

func CreateTree() *Node {

	return nil
}

func (this *Node) GetLeftChild() *Node {
	return this.left
}

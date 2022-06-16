package main

type Node struct {
	Val      int
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []int) []int {
	array = append(array, n.Val)
	for _, child := range n.Children {
		array = child.DepthFirstSearch(array)
	}
}

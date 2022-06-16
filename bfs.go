package main

import "fmt"

type Node struct {
	Value    int
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []int) []int {
	queue := []*Node{n}
	for len(queue) > 0 {
		current := queue[0]
		queue := queue[1:]
		array = append(array, current.Value)
		for _, child := range n.Children {
			queue = append(queue, child)
		}
	}
	return array
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := &Node{Value: 1}
	n.Children = []*Node{&Node{Value: 2}, &Node{Value: 3}}
	n.Children[0].Children = []*Node{&Node{Value: 4}, &Node{Value: 5}}
	n.Children[1].Children = []*Node{&Node{Value: 6}, &Node{Value: 7}}
	fmt.Println(n.BreadthFirstSearch(array))
}

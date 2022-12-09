package main

import "fmt"

type Value int
type Node struct {
	value Value
	next  *Node
}

func (n *Node) createNode(value Value) *Node {
	newNode := &Node{value: value}
	n.next = newNode
	return newNode
}

func printList(node *Node) {
	fmt.Println(node)
	if node.next != nil {
		printList(node.next)
	}
}
func main() {
	node := Node{value: 1}
	node.createNode(2).createNode(3).createNode(4)
	printList(&node)
}

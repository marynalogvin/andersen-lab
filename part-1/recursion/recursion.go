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

func printListReverse(node *Node) {
	if node.next != nil {
		printListReverse(node.next)
	}
	fmt.Println(node)
}

func main() {
	node := Node{value: 1}
	node.createNode(2).createNode(3).createNode(4)
	printList(&node)
	printListReverse(&node)
}

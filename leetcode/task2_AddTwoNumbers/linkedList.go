package addTwo

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (n *Node) createNode(value int) *Node {
	newNode := &Node{value: value}
	n.next = newNode
	return newNode
}

func createList(list []int) *Node {
	node := &Node{value: list[0]}
	tmp := node
	for i := 1; i < (len(list)); i++ {
		tmp = tmp.createNode(list[i])
	}
	return node
}

func printList(node *Node) {
	fmt.Println(node)
	if node.next != nil {
		printList(node.next)
	}
}

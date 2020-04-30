package main

import "fmt"

type listNode struct {
	next *listNode
	value int
}

func (root *listNode) nodePrint() {
	node := root
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (node *listNode) reverse() *listNode {
	var preNode *listNode
	var nextNode *listNode
	//pre = node
	nextNode = node
	for nextNode != nil {

		tmp := nextNode.next
		nextNode.next = preNode
		preNode = nextNode
		nextNode = tmp

	}
	return preNode
}

func main()  {
	a := listNode{nil, 1}
	b := listNode{&a, 2}
	c := listNode{&b, 3}
	d := listNode{&c, 4}

	d.nodePrint()

	fmt.Println("reverse!")
	res := d.reverse()

	res.nodePrint()
}

package main

type Node[T any] struct { // Linked List Node
	next *Node[T]
	prev *Node[T]
	val  T
}

func main() {
	root := &Node[int]{val: 1}
	root.next = &Node[int]{val: 2, prev: root}
	leafNode := &Node[int]{val: 3, prev: root.next}
	root.next.next = leafNode

	for n := root; n != nil; n = n.next {
		println(n.val)
	}

	for n := leafNode; n != nil; n = n.prev {
		println(n.val)
	}
}

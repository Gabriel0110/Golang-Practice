package main

type Node struct {
	Value int
	Next  *Node
}

func hasCycle(head *Node) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			return true // cycle found
		}
	}
	return false
}

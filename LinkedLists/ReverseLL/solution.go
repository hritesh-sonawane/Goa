package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(1) time | O(1) space
func ReverseLinkedList(head *LinkedList) *LinkedList {
	var prevNode, currNode *LinkedList = nil, head
	for currNode != nil {
		nextNode := currNode.Next
		currNode.Next = prevNode
		prevNode = currNode
		currNode = nextNode
	}
	return prevNode
}
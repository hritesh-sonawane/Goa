package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// O(n) time | O(1) space - where n is the number of nodes in the LinkedList
func NoDuplicates(linkedList *LinkedList) *LinkedList {
	currentNode := linkedList
	for currentNode != nil {
		nextUniqueNode := currentNode.Next
		for nextUniqueNode != nil && nextUniqueNode.Value == currentNode.Value {
			nextUniqueNode = nextUniqueNode.Next
		}
		currentNode.Next = nextUniqueNode
		currentNode = nextUniqueNode
	}
	return linkedList
}
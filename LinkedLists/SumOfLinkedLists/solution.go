package main

type LinkedList struct {
	Value int
	Next  *LinkedList
}

// 0(max(m, n)) time | 0(max(m, n)) space
func SumOfLinkedLists(linkedListOne *LinkedList, linkedListTwo *LinkedList) *LinkedList {
	// This variable will store a dummy node whose .next
	// attribute will point to the head of our new LL.
	newLinkedListHeadPointer := &LinkedList{Value: 0}
	currentNode := newLinkedListHeadPointer
	carry := 0

	nodeOne := linkedListOne
	nodeTwo := linkedListTwo
	for nodeOne != nil || nodeTwo != nil || carry != 0 {
		var valueOne, valueTwo int
		if nodeOne != nil {
			valueOne = nodeOne.Value
		}
		if nodeTwo != nil {
			valueTwo = nodeTwo.Value
		}
		sumOfValues := valueOne + valueTwo + carry

		newValue := sumOfValues % 10
		newNode := &LinkedList{Value: newValue}
		currentNode.Next = newNode
		currentNode = newNode

		carry = sumOfValues / 10
		if nodeOne != nil {
			nodeOne = nodeOne.Next
		}
		if nodeTwo != nil {
			nodeTwo = nodeTwo.Next
		}
	}
	return newLinkedListHeadPointer.Next
}
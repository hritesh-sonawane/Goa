package main

type BST struct {
	Value int
	Left *BST
	Right *BST
}

// O(n) time | O(n) space, n -> number of nodes in Binary Tree
func BranchSums(root *BST) []int {
	sum := []int{}
	calcBranchSums(root, 0, &sum)
	return sum
}

func calcBranchSums(node *BST, runningSum int, sum *[]int) {
	if node == nil {
		return
	}

	runningSum += node.Value
	if node.Left == nil && node.Right == nil {
		*sum = append(*sum, runningSum)
		return
	}

	calcBranchSums(node.Left, runningSum, sum)
	calcBranchSums(node.Right, runningSum, sum)
}
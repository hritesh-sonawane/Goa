package main

type BinaryTree struct {
	Value int
	Left, Right *BinaryTree
}

// Average case: when the tree is balanced
// O(n) time | O(h) space | n => no. of nodes | h => height
func NodeDepths(root *BinaryTree) int {
	return nodeDepthsHelper(root, 0)
}

func nodeDepthsHelper(root *BinaryTree, depth int) int {
	if root == nil {
		return 0
	}
	return depth + nodeDepthsHelper(root.Left, depth+1) + nodeDepthsHelper(root.Right, depth+1)
}
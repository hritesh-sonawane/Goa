package main

// O(n) time | O(1) space 0 where n: len of input array
func WaterArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	leftIdx := 0
	rightIdx := len(heights) - 1
	leftMax := heights[leftIdx]
	rightMax := heights[rightIdx]
	surfaceArea := 0

	for leftIdx < rightIdx {
		if heights[leftIdx] < heights[rightIdx] {
			leftIdx++
			leftMax = max(leftMax, heights[leftIdx])
			surfaceArea += leftMax - heights[leftIdx]
		} else {
			rightIdx--
			rightMax = max(rightMax, heights[rightIdx])
			surfaceArea += rightMax - heights[rightIdx]
		}
	}
	return surfaceArea
}

func max(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}

package main

import "fmt"

// O(n) time | O(1) space
func HasSingleCycle(array []int) bool {
	numElementsVisited := 0
	currentIdx := 0
	for numElementsVisited < len(array) {
		if numElementsVisited > 0 && currentIdx == 0 {
			return false
		}
		numElementsVisited++
		currentIdx = getNextIdx(currentIdx, array)
	}
	return currentIdx == 0
}

func getNextIdx(currentIdx int, array []int) int {
	jump := array[currentIdx]
	nextIdx := (currentIdx + jump) % len(array)
	if nextIdx >= 0 {
		return nextIdx
	}
	return nextIdx + len(array)
}

func main() {
	arr := []int{2, 3, 1, -4, -4, 2}
	fmt.Println(HasSingleCycle(arr))
}
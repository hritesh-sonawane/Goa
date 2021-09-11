package main

import "fmt"

// O(nlog(n)) time | O(1) space
func HeapSort(array []int) []int {
	buildMaxHeap(array)

	for endIndex := len(array) - 1; endIndex >= 1; endIndex-- {
		swap(0, endIndex, array)
		shiftDown(0, endIndex-1, array)
	}
	return array
}

func buildMaxHeap(array []int) {
	first := (len(array) - 2) / 2

	for currentIndex := first + 1; currentIndex >= 0; currentIndex-- {
		shiftDown(currentIndex, len(array)-1, array)
	}
}

func shiftDown(currentIndex, endIndex int, heap []int) {
	childOneIndex := currentIndex*2 + 1

	for childOneIndex <= endIndex {
		childTwoIndex := -1

		if currentIndex*2+2 <= endIndex {
			childTwoIndex = currentIndex*2 + 2
		}

		indexToSwap := childOneIndex

		if childTwoIndex > -1 && heap[childTwoIndex] > heap[childOneIndex] {
			indexToSwap = childTwoIndex
		}

		if heap[indexToSwap] > heap[currentIndex] {
			swap(currentIndex, indexToSwap, heap)
			currentIndex = indexToSwap
			childOneIndex = currentIndex*2 + 1
		} else {
			return
		}
	}
}

func swap(i, j int, array []int) {
	array[i], array[j] = array[j], array[i]
}

func main() {
	array := []int{8, 5, 2, 9, 5, 6, 3}
	fmt.Println(HeapSort(array))
}

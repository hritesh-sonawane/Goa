package main

import "fmt"

func ShiftedBinarySearch(array []int, target int) int {
	return helper(array, target, 0, len(array)-1)
}

func helper(array []int, target, left, right int) int {
	if left > right {
		return -1
	}

	middle := (left + right) / 2
	potentialMatch := array[middle]
	leftNum, rightNum := array[left], array[right]
	newLeft, newRight := middle+1, right

	if target == potentialMatch {
		return middle
	} else if leftNum <= potentialMatch {
		if target < potentialMatch && target >= leftNum {
			newLeft, newRight = left, middle-1
		}
	} else {
		if !(target > potentialMatch && target <= rightNum) {
			newLeft, newRight = left, middle-1
		}
	}
	return helper(array, target, newLeft, newRight)
}

func main() {
	array := []int{45, 61, 71, 72, 73, 0, 1, 21, 33, 37}
	target := 33
	fmt.Println(ShiftedBinarySearch(array, target))
}
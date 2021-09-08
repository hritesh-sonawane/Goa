package main

import "fmt"

func SearchForRange(array []int, target int) []int {
	finalRange := []int{-1, -1}
	alteredBinarySearch(array, target, 0, len(array)-1, finalRange, true)
	alteredBinarySearch(array, target, 0, len(array)-1, finalRange, false)
	return finalRange
}

func alteredBinarySearch(array []int, target, left, right int, finalRange []int, goLeft bool) {
	if left > right {
		return
	}
	mid := (left+ right)/2
	if array[mid] < target {
		alteredBinarySearch(array, target, mid+1, right, finalRange, goLeft)
	} else if array[mid] > target {
		alteredBinarySearch(array, target, left, mid-1, finalRange, goLeft)
	} else {
		if goLeft {
			if mid == 0 || array[mid-1] != target {
				finalRange[0] = mid
			} else {
				alteredBinarySearch(array, target, left, mid-1, finalRange, goLeft)
			}
		} else {
			if mid == len(array)-1 || array[mid+1] != target {
				finalRange[1] = mid
			} else {
				alteredBinarySearch(array, target, mid+1, right, finalRange, goLeft)
			}
		}
	}
}

func main() {
	arr := []int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}
	target := 45
	fmt.Println(SearchForRange(arr, target))
}
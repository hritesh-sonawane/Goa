package main

import "fmt"

// O(n) time | O(1) space - n => length of the array
func IsMonotonic(array []int) bool {
	isNonDecr := true
	isNonIncr := true
	for i := 1; i < len(array); i++ {
		if array[i] < array[i-1] {
			isNonDecr = false
		}
		if array[i] > array[i-1] {
			isNonIncr = false
		}
	}
	return isNonDecr || isNonIncr
}

func main() {
	arr := []int{-1, -5, -10, -10, -1102, -9001}
	fmt.Println(IsMonotonic(arr))
}
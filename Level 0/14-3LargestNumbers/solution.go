package main

import (
	"fmt"
	"math"
)

// O(n) time | O(1) space
func ThreeLargestNumbers(array []int) []int {
	three := []int{math.MinInt32, math.MinInt32, math.MinInt32}
	for _, num := range array {
		updateLargest(three, num)
	}
	return three
}

func updateLargest(three []int, num int) {
	if num > three[2] {
		shiftAndUpdate(three, num, 2)
	} else if num> three[1] {
		shiftAndUpdate(three, num, 1)
	} else if num > three[0] {
		shiftAndUpdate(three, num, 0)
	}
}

func shiftAndUpdate(array []int, num int, idx int) {
	for i := 0; i < idx+1; i++ {
		if i == idx {
			array[i] = num
		} else {
			array[i] = array[i+1]
		}
	}
}

func main() {
	arr := []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	fmt.Println(ThreeLargestNumbers(arr))
}
package main

import "fmt"

// O(n) time | O(n) space
func TwoNumberSum(arr []int, target int) []int {
	nums := map[int]bool{}
	for _, num := range arr {
		potentialMatch := target - num
		if _, found := nums[potentialMatch]; found {
			return []int{potentialMatch, num}
		}
		nums[num] = true
	}
	return []int{}
}

func main() {
	testArr := []int{3, 5, -4, 8, 11, 1, -1, 6} 
	fmt.Println(TwoNumberSum(testArr, 10))
}
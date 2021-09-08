package main

import "fmt"

// O(n) time | O(n) space
func NextGreaterElement(array []int) []int {
	result := make([]int, 0)
	for range array {
		result = append(result, -1)
	}
	stack := make([]int, 0)

	for idx := 2*len(array) - 1; idx >= 0; idx-- {
		circularIdx := idx % len(array)

		for len(stack) > 0 {
			if stack[len(stack)-1] <= array[circularIdx] {
				stack = stack[:len(stack)-1]
			} else {
				result[circularIdx] = stack[len(stack)-1]
				break
			}
		}
		stack = append(stack, array[circularIdx])
	}
	return result
}

func main() {
	arr := []int{2, 5, -3, -4, 6, 7, 2}
	fmt.Println(NextGreaterElement(arr))
}

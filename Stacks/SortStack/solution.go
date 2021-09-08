package main

import "fmt"

func SortStack(stack []int) []int {
	if len(stack) == 0 {
		return stack
	}

	top := stack[len(stack)-1]
	stack = stack[:len(stack)-1] //pop?
	SortStack(stack)

	insertInSortedOrder(&stack, top)
	return stack
}

func insertInSortedOrder(stack *[]int, value int) {
	if len(*stack) == 0 || (*stack)[len(*stack)-1] <= value {
		*stack = append(*stack, value) //append?
		return
	}

	top := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1] // pop?
	insertInSortedOrder(stack, value)
	*stack = append(*stack, top) //append?
}

func main() {
	stack := []int{-5, 2, -2, 4, 3, 1}
	fmt.Println(SortStack(stack))
}
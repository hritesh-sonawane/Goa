package main

import "fmt"

type MinMaxStack struct {
	stack []int
	minMaxStack []entry
}

type entry struct {
	min int
	max int
}

func NewMinMaxStack() *MinMaxStack {
	return &MinMaxStack{}
}

// O(1) time | O(1) space
func (stack *MinMaxStack) Peek() int {
	return stack.stack[len(stack.stack)-1]
}

// O(1) time | O(1) space
func (stack *MinMaxStack) Pop() int {
	stack.minMaxStack = stack.minMaxStack[:len(stack.minMaxStack)-1]
	out := stack.stack[len(stack.stack)-1]
	stack.stack = stack.stack[:len(stack.stack)-1]
	return out
}

// O(1) time | O(1) space
func (stack *MinMaxStack) Push(number int) {
	newMinMax := entry{min: number, max: number}
	if len(stack.minMaxStack) > 0 {
		lastMinMax := stack.minMaxStack[len(stack.minMaxStack)-1]
		newMinMax.min = min(lastMinMax.min, number)
		newMinMax.max = max(lastMinMax.max, number)
	}
	stack.minMaxStack = append(stack.minMaxStack, newMinMax)
	stack.stack = append(stack.stack, number)
}

// O(1) time | O(1) space
func (stack *MinMaxStack) GetMin() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].min
}

// O(1) time | O(1) space
func (stack *MinMaxStack) GetMax() int {
	return stack.minMaxStack[len(stack.minMaxStack)-1].max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	stx := MinMaxStack{}
	stx.Push(1)
	fmt.Println("Push: 1")
	fmt.Println(stx.stack)
	stx.Push(2)
	fmt.Println("Push: 2")
	fmt.Println(stx.stack)
	stx.Push(3)
	fmt.Println("Push: 3")
	fmt.Println(stx.stack)
	fmt.Print("Max in stack: ")
	fmt.Println(stx.GetMax())
	fmt.Print("Min in stack: ")
	fmt.Println(stx.GetMin())
	stx.Peek()
	fmt.Println("Peek stack")
	fmt.Println(stx.stack)
	stx.Pop()
	fmt.Println("Pop element")
	fmt.Println(stx.stack)
}
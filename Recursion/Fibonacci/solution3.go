package main

import "fmt"

// O(n) time | O(1) space
func Fibo(n int) int {
	lastTwo := []int{0, 1}
	counter := 3
	for counter <= n {
		nextFib := lastTwo[0] + lastTwo[1]
		lastTwo = []int{lastTwo[1], nextFib}
		counter += 1
	}
	if n > 1 {
		return lastTwo[1]
	}
	return lastTwo[0]
}

func main() {
	fmt.Println(Fibo(6))
}
package main

import "fmt"

// O(n) time | O(n) space
func Fibo(n int) int {
	return helper(n, map[int]int{
		1: 0,
		2: 1,
	})
}

func helper(n int, memorize map[int]int) int {
	if val, found := memorize[n]; found {
		return val
	}
	memorize[n] = helper(n-1, memorize) + helper(n-2, memorize)
	return memorize[n]
}

func main() {
	fmt.Println(Fibo(6))
}
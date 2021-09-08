package main

import "fmt"

// O(2^n) time | O(n) space
func Fibo(n int) int {
	if n == 1 {
		return 0
	} else if n == 2 {
		return 1
	}
	return Fibo(n-1) + Fibo(n-2)
}

func main() {
	fmt.Println(Fibo(6))
}
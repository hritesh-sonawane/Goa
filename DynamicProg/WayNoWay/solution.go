package main

import "fmt"

// O(nd) time | O(n) space
// n => target | d => no. of coin denominations
func WaysToMakeChange(n int, denoms []int) int {
	ways := make([]int, n+1)
	ways[0] = 1
	for _, denom := range denoms {
		for amount := 1; amount < n+1; amount++ {
			if denom <= amount {
				ways[amount] += ways[amount-denom]
			}
		}
	}
	return ways[n]
}

func main() {
	arr := []int{1, 5, 10, 25}
	fmt.Println(WaysToMakeChange(10, arr))
}
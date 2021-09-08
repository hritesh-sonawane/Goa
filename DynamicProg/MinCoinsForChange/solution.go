package main

import (
	"fmt"
	"math"
)

// O(nd) time | O(n) space
// n => target | d => no. of coin denominations
func MinCoinsForChange(n int, denoms []int) int {
	numOfCoins := make([]int, n+1)
	for i := range numOfCoins {
		numOfCoins[i] = math.MaxInt32
	}
	numOfCoins[0] = 0
	for _, denom := range denoms {
		for amount := range numOfCoins {
			if denom <= amount {
				numOfCoins[amount] = min(numOfCoins[amount], numOfCoins[amount-denom]+1)
			}
		}
	}
	if numOfCoins[n] != math.MaxInt32 {
		return numOfCoins[n]
	}
	return -1
}

func min(arg1 int, rest ...int) int {
	curr := arg1
	for _, num := range rest {
		if num < curr {
			curr = num
		}
	}
	return curr
}

func main() {
	denoms := []int{1, 5, 10}
	fmt.Println(MinCoinsForChange(7, denoms))
}
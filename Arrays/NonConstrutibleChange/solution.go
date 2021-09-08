package main

import (
	"fmt"
	"sort"
)

// O(nlogn) time | O(1) space - where n is the number of coins
func Thanos(coins []int) int {
	sort.Ints(coins)

	var currentChangeCreated = 0
	for _, coin := range coins {
		if coin > currentChangeCreated + 1 {
			return currentChangeCreated + 1
		}
		currentChangeCreated += coin
	}
	return currentChangeCreated + 1
}

func main() {
	dodgecoin := []int{1, 2, 5} 
	fmt.Println(Thanos(dodgecoin))
}
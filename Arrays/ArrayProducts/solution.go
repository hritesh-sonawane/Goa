package main

import "fmt"

// O(n^2) time | O(n) space
func ArrayOfProducts(array []int) []int {
	products := make([]int, len(array))

	for i := range array {
		res := 1
		for j := range array {
			if i != j {
				res *= array[j]
			}
		}
		products[i] = res
	}
	return products
}

func main() {
	arr := []int{5, 1, 4, 2}
	fmt.Println(ArrayOfProducts(arr))
}
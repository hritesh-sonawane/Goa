package main

import "fmt"

type SpecialArray []interface{}

// O(n) time | O(d) space - n is the total elements in array
// including sub-elements & d is the max depth of special array
func ProductSum(array SpecialArray) int {
	return helper(array, 1)
}

func helper(array SpecialArray, multiplier int) int {
	sum := 0
	for _, el:= range array {
		if cast, ok := el.(SpecialArray); ok {
			sum += helper(cast, multiplier+1)
		} else if cast, ok := el.(int); ok {
			sum += cast
		}
	}
	return sum * multiplier
}

func main() {
	arr := SpecialArray{5, 2, 4}
	fmt.Println(ProductSum(arr))
}
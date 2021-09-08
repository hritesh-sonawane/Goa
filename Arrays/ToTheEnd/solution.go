package main

import "fmt"

func ToTheEnd(array []int, toMove int) []int {
	i, j := 0, len(array)-1
	for i < j {
		for i < j && array[j] == toMove {
			j--
		}
		if array[i] == toMove {
			array[i], array[j] = array[j], array[i]
		}
		i++
	}
	return array
}

func main() {
	arr := []int{2, 1, 2, 2, 2, 2, 3, 4, 2}
	fmt.Println(ToTheEnd(arr, 2))
}
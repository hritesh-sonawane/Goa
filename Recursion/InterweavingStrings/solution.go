package main

import "fmt"

// O(mn) time | O(mn) space
func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(two)+len(one) {
		return false
	}
	cache := make([][]*bool, len(one)+1)
	for i := 0; i < len(one)+1; i++ {
		cache[i] = make([]*bool, len(two)+1)
	}
	return areInterwoven(one, two, three, 0, 0, cache)
}

func areInterwoven(one, two, three string, i, j int, cache [][]*bool) bool {
	if cache[i][j] != nil {
		return *cache[i][j]
	}

	k := i + j
	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		result := areInterwoven(one, two, three, i+1, j, cache)
		cache[i][j] = &result
		if result {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		result := areInterwoven(one, two, three, i, j+1, cache)
		cache[i][j] = &result
		return result
	}

	result := false
	cache[i][j] = &result
	return result
}

func main() {
	one := "abc"
	two := "def"
	three := "abdecf"
	fmt.Println(InterweavingStrings(one, two, three))
}
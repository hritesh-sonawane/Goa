// Easier solution, but not effecient
package main

import "fmt"

// O(2^(m+n)) time | O(m+n) space
func InterweavingStrings(one, two, three string) bool {
	if len(three) != len(two)+len(one) {
		return false
	}
	return areInterwoven(one, two, three, 0, 0)
}

func areInterwoven(one, two, three string, i, j int) bool {
	k := i + j
	if k == len(three) {
		return true
	}

	if i < len(one) && one[i] == three[k] {
		if areInterwoven(one, two, three, i+1, j) {
			return true
		}
	}

	if j < len(two) && two[j] == three[k] {
		return areInterwoven(one, two, three, i, j+1)
	}
	return false
}

func main() {
	one := "abc"
	two := "def"
	three := "abdecf"
	fmt.Println(InterweavingStrings(one, two, three))
}
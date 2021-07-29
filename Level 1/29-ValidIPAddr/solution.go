package main

import (
	"fmt"
	"strconv"
	"strings"
)

// O(1) time | O(1) space
func ValidIPAddresses(str string) []string {
	ipAddressesFound := make([]string, 0)

	for i := 1; i < min(len(str), 4); i++ {
		curentIPAddressParts := []string{"", "", "", ""}

		curentIPAddressParts[0] = str[:i]
		if !isValidPart(curentIPAddressParts[0]) {
			continue
		}

		for j := i+1; j < i+min(len(str)-i, 4); j++ {
			curentIPAddressParts[1] = str[i:j]
			if !isValidPart(curentIPAddressParts[1]) {
				continue
			}

			for k := j+1; k < j+min(len(str)-j, 4); k++ {
				curentIPAddressParts[2] = str[j:k]
				curentIPAddressParts[3] = str[k:]

				if isValidPart(curentIPAddressParts[2]) && isValidPart(curentIPAddressParts[3]) {
					ipAddressesFound = append(ipAddressesFound, strings.Join(curentIPAddressParts, "."))
				}
			}
		}
	}
	return ipAddressesFound
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isValidPart(str string) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	if i > 255 {
		return false
	}
	return len(str) == len(strconv.Itoa(i))
}

func main() {
	string := "1921680"
	fmt.Println(ValidIPAddresses(string))
}
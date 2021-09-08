package main

import (
	"fmt"
	"strings"
)

type intervals []*interval
type interval struct {
	left  int
	right int
}

// Average case: O(m+n) | O(m) space
// m: length of str | n: length of substr
func UnderscorifySubstring(str, substr string) string {
	locations := getLocations(str, substr)
	locations = locations.collapse()
	return underscorify(str, locations)
}

func getLocations(str, substr string) intervals {
	result := intervals{}
	for start := 0; start < len(str); {
		nextIdx := strings.Index(str[start:], substr)
		if nextIdx == -1 {
			break
		}
		nextIdx += start
		result = append(result, &interval{nextIdx, nextIdx+len(substr)})
		start = nextIdx + 1
	}
	return result
}

func (array intervals) collapse() intervals {
	// If the array is empty, do nothing
	if len(array) == 0 {
		return array
	}

	result := intervals{array[0]}
	previous := array[0]
	for i := 1; i < len(array); i++ {
		current := array[i]
		if current.left <= previous.right {
			// Collapse the two intervals
			previous.right = current.right
		} else {
			result = append(result, current)
			previous = current
		}
	}
	return result
}

func underscorify(str string, locations intervals) string {
	if len(locations) == 0 {
		return str
	}

	// We know that the resulting string will have an additional 
	// 2*len(intervals) characters
	result := make([]rune, len(str)+2*len(locations))
	resultIndex := 0
	locationIndex := 0
	for i, r := range str {
		location := locations[locationIndex]
		if i == location.left {
			result[resultIndex] = '_'
			resultIndex += 1
		} else if i == location.right {
			result[resultIndex] = '_'
			resultIndex += 1
			if locationIndex+1 < len(locations) {
				locationIndex += 1
			}
 		}
		 result[resultIndex] = r
		 resultIndex += 1
	}
	
	if locations[locationIndex].right == len(str) {
		result[len(result)-1] = '_'
	}
	return string(result)
}

func main() {
	str := "testthis is a testtest to see if testestest it works"
	substr := "test"
	fmt.Println(UnderscorifySubstring(str, substr))
}
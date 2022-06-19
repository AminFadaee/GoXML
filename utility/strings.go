package utility

import "strings"

func firstIndex(s string, substrs ...string) int {
	minIndex := -1
	for _, substr := range substrs {
		index := strings.Index(s, substr)
		if index != -1 && (minIndex == -1 || index < minIndex) {
			minIndex = index
		}
	}
	return minIndex
}

func IndeciesInString(s string, substrs ...string) []int {
	result := make([]int, 0, 10)
	hit := firstIndex(s, substrs...)
	index := 0
	for hit != -1 {
		result = append(result, index+hit)
		index = result[len(result)-1] + 1
		hit = firstIndex(s[index:], substrs...)
	}
	return result
}

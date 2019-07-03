package main

import "strings"

// Slice Window
func lengthOfLongestSubstring2(s string) int {
	sLen := len(s)
	temp := ""

	result := 0
	i := 0
	j := 0

	for i < sLen && j < sLen {
		str := string(s[j])

		if strings.Contains(temp, str) {
			temp = temp[1:]
			i += 1
		} else {
			temp += str
			j += 1
		}

		if result < len(temp) {
			result = len(temp)
		}
	}

	return result
}

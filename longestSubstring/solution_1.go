package main

import (
	"strings"
)

// Slice Window
func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	temp := ""

	result := 0
	i := 0
	j := 0

	for i < sLen && j < sLen {
		str := string(s[j])

		if !strings.Contains(temp, str) {
			temp += str
			j++
		} else {
			temp = temp[1:]
			i++
		}

		if result < len(temp) {
			result = len(temp)
		}
	}

	return result
}

// Brute Force
func lengthOfLongestSubstringByBruteForce(s string) int {
	longest := 0

	for i := 0; i < len(s); i++ {
		temp := map[string]string{}
		for j := i; j < len(s); j++ {
			str := string(s[j])
			_, exists := temp[str]

			if exists {
				break
			}

			temp[str] = str
		}

		if len(temp) > longest {
			longest = len(temp)
		}
	}

	return longest
}

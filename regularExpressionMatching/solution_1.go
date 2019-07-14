package main

import "fmt"

func isMatch(s string, p string) bool {
	i := 0
	curr := 0

	for curr < len(p) {
		// In '*' situation
		if curr+1 < len(p) && string(p[curr+1]) == "*" {
			if i >= len(s) || (s[i] != p[curr] && string(p[curr]) != ".") {
				curr += 2
			} else {
				i += 1
			}

			continue
		}

		// Check the index if it out of range
		if i == len(s) {
			fmt.Println("=====")
			fmt.Println("i is out of range")
			fmt.Println("=====")
			return false
		}

		// Validate string
		if !inLowercaseAToZ(s[i]) {
			fmt.Println("=====")
			fmt.Println("s[i] is invalid")
			fmt.Println("=====")
			return false
		}

		// In `.` situation
		if string(p[curr]) == "." {
			i += 1
			curr += 1
			continue
		}

		if p[curr] == s[i] {
			i += 1
			curr += 1
		} else {
			fmt.Println("=====")
			fmt.Println("p[curr] != s[i]")
			fmt.Println("=====")
			return false
		}
	}

	if i != len(s) {
		return false
	}

	return true
}

func inLowercaseAToZ(s byte) bool {
	return s >= 97 && s <= 122
}

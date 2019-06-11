package main

func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i += 1 {
		len1 := expandAroundCenter(s, i, i)
		len2 := expandAroundCenter(s, i, i+1)
		longest := len1
		if len2 > len1 {
			longest = len2
		}
		if longest > end-start {
			start = i - (longest-1)/2
			end = i + longest/2
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) int {
	l, r := left, right

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l -= 1
		r += 1
	}

	return r - l - 1
}

package main

func longestPalindrome2(s string) string {
	if len(s) <= 1 {
		return s
	}

	start, end := 0, 0

	for i := 0; i < len(s); i += 1 {
		l1, r1 := expandAroundCenter(s, i, i)
		l2, r2 := expandAroundCenter(s, i, i+1)
		len1 := r1 - l1 + 1
		len2 := r2 - l2 + 1

		longest := len1
		l := l1
		r := r1
		if len2 > len1 {
			longest = len2
			l = l2
			r = r2
		}

		if longest > end-start {
			start = l
			end = r
		}
	}

	return s[start : end+1]
}

func expandAroundCenter(s string, left int, right int) (int, int) {
	l, r := left, right

	for l >= 0 && r < len(s) && s[l] == s[r] {
		l -= 1
		r += 1
	}

	return l + 1, r - 1
}

package main

func longestPalindrome(s string) string {
	l := len(s)

	for {
		if l == 0 {
			break
		}

		i, j := 0, l-1
		for {
			if j >= len(s) {
				break
			}

			if isPalindrome(s[i:j]) {
				return s[i:j]
			}

			i += 1
			j += 1
		}

		l -= 1
	}

	return ""
}

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if i == j {
			break
		}

		if s[i] != s[j] {
			return false
		}
	}

	return true
}

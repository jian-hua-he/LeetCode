package main

func longestPalindrome(s string) string {
	for l := len(s); l > 0; l -= 1 {
		for i, j := 0, l-1; j < len(s); i, j = i+1, j+1 {
			if isPalindrome(s[i:j]) {
				return s[i:j]
			}
		}
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

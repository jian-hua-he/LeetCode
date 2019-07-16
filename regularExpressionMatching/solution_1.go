package main

func isMatch(s string, p string) bool {
	if p == "" {
		return s == ""
	}

	firstMatch := s != "" && (s[0] == p[0] || string(p[0]) == ".")

	if len(p) >= 2 && string(p[1]) == "*" {
		return isMatch(s, p[2:]) || firstMatch && isMatch(s[1:], p)
	}

	return firstMatch && isMatch(s[1:], p[1:])
}

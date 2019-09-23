package main

/**
 * 1. Define base case
 * 2. Extend base case or update base case
 * 3. Reach the goal
 */

func isMatch2(s string, p string) bool {
	// dp[i][j] = string s with len i matches string p with len j
	// Default elements in dp are false
	dp := make([][]bool, len(s)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, len(p)+1)
	}

	// Build Base case
	// s = "", p = "" => is matched
	dp[0][0] = true

	// If s is not empty, p = "" => dp[i][0] => not matched
	// Since every default elements are false. We don’t have to do anything here

	// If s = "", p is not empty => dp[0][j]
	// It depend on p[j]. If p[j] is "*" than matched
	for j := 2; j < len(p)+1; j += 1 {
		if string(p[j-1]) == "*" && dp[0][j-2] {
			dp[0][j] = true
		}
	}

	// Extend base case
	// i and j is for var s and p. not for array dp
	// So when match s[i] and p[j]. current dp would be dp[i+1][j+1]
	for i := 0; i < len(s); i += 1 {
		for j := 0; j < len(p); j += 1 {
			// If p[j] is "." or s[i] == p[j]
			// currently dp should equal the previous matched
			if string(p[j]) == "." || (s[i] == p[j]) {
				dp[i+1][j+1] = dp[i][j]
				continue
			}

			// If p[j] == "*", match p[j-2] and s[i]
			// If not matched then compare p[j-1] and s[i]
			// If p[j-1] == s[i], get result from dp[i][j+1]
			if string(p[j]) == "*" {
				if dp[i+1][j-1] == true {
					dp[i+1][j-1] = true
					continue
				}

				if p[j-1] == s[i] || string(p[j-1]) == "." {
					dp[i+1][j+1] = dp[i][j+1]
				}
			}
		}
	}

	return dp[len(s)][len(p)]
}

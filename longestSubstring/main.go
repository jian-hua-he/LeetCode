package main

func lengthOfLongestSubstring(s string) int {
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

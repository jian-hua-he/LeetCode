package main

func lengthOfLongestSubstring(s string) int {
	longest := 0

	temp := map[string]string{}
	for _, v := range s {
		str := string(v)

		if _, ok := temp[str]; ok {
			if len(temp) > longest {
				longest = len(temp)
			}

			temp = map[string]string{}
		}

		temp[str] = str
	}

	return longest
}

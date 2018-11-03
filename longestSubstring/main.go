package main

func lengthOfLongestSubstring(s string) int {
	longest := 0

	if len(s) == 1 {
		return 1
	}

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

	if len(temp) > longest {
		longest = len(temp)
	}

	return longest
}

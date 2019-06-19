package main

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	m := make(map[int]string, 0)
	for i := 0; i < numRows; i += 1 {
		m[i] = ""
	}

	m = recoursiveConvert(s, numRows, m)
	result := ""
	for i := 0; i < numRows; i += 1 {
		result += m[i]
	}

	return result
}

func recoursiveConvert(s string, nums int, m map[int]string) map[int]string {
	if s == "" {
		return m
	}

	min := nums
	if len(s) < min {
		min = len(s)
	}

	for k, v := range s[:min] {
		m[k] += string(v)
	}

	max := nums*2 - 2
	if len(s) < max {
		max = len(s)
	}

	for k, v := range s[min:max] {
		m[nums-k-2] += string(v)
	}

	return recoursiveConvert(s[max:], nums, m)
}

package main

import (
	"math"
	"strconv"
)

func reverse1(x int) int {
	n := x
	base := 1
	if x < 0 {
		base = -1
		n *= -1
	}

	s := strconv.Itoa(n)
	r := len(s) / 2
	l := r
	if len(s)%2 == 0 {
		r = len(s) / 2
		l = r - 1
	}

	rStr := ""
	for {
		if r > len(s) || l < 0 {
			break
		}

		if r == l {
			rStr += string(s[r])
		} else {
			rStr = string(s[r]) + rStr + string(s[l])
		}

		r += 1
		l -= 1
	}

	result, _ := strconv.Atoi(rStr)
	result *= base

	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}

	return result
}

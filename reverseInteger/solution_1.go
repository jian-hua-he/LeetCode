package main

import "math"

func reverse1(x int) int {
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}

	return x
}

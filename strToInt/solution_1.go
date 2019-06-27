package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

func myAtoi1(str string) int {
	re := regexp.MustCompile(`^(\s*\+?\-?\d*)`)
	s := re.FindString(str)
	s = strings.Trim(s, " ")
	r, _ := strconv.Atoi(s)

	if r > math.MaxInt32 {
		return math.MaxInt32
	}

	if r < math.MinInt32 {
		return math.MinInt32
	}

	return r
}

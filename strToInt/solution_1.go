package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func myAtoi1(str string) int {
	var result int32 = 0

	re := regexp.MustCompile(`^\s*\-?\d*`)
	s := re.FindString(str)
	s = strings.Trim(s, " ")
	r, _ := strconv.Atoi(s)
	result = int32(r)

	return int(result)
}

package main

import (
	"fmt"
)

func convert1(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	m := make(map[int]string, 0)
	for i := 0; i < numRows; i += 1 {
		m[i] = ""
	}

	length := numRows*2 - 2

	for k, v := range s[:numRows] {
		m[k] += string(v)
	}

	for k, v := range s[numRows:length] {
		m[numRows-k-2] += string(v)
	}

	for i := 0; i < numRows; i += 1 {
		fmt.Println(m[i])
	}

	return s
}

package main

import "testing"

func TestCaseOne(t *testing.T) {
	input := "42"
	want := 42
	got := myAtoi1(input)

	if got != want {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, got, want)
	}
}

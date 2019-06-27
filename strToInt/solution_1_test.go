package main

import "testing"

func TestCase1ForSolution1(t *testing.T) {
	input := "42"
	want := 42
	got := myAtoi1(input)

	if got != want {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, got, want)
	}
}

func TestCase2ForSolution1(t *testing.T) {
	input := "-42"
	want := -42
	got := myAtoi1(input)

	if got != want {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, got, want)
	}
}

func TestCase3ForSolution1(t *testing.T) {
	input := "   -42"
	want := -42
	got := myAtoi1(input)

	if got != want {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, got, want)
	}
}

func TestCase4ForSolution1(t *testing.T) {
	input := "words and 987"
	want := 0
	got := myAtoi1(input)

	if got != want {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, got, want)
	}
}

func TestCase5ForSolution1(t *testing.T) {
	input := "-91283472332"
	want := -2147483648
	got := myAtoi1(input)

	if got != want {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, got, want)
	}
}

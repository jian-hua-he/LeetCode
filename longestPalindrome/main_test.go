package main

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	input := "babad"
	expected := "bab"
	result := longestPalindrome(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

func TestCaseTwo(t *testing.T) {
	input := "cbbd"
	expected := "bb"
	result := longestPalindrome(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

func TestCaseThree(t *testing.T) {
	input := "12ece41abcba41"
	expected := "abcba"
	result := longestPalindrome(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

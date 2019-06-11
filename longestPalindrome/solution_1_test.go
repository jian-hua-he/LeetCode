package main

import (
	"testing"
)

func TestCaseOne(t *testing.T) {
	input := "babad"
	expected1 := "bab"
	expected2 := "aba"
	result := longestPalindrome2(input)

	if result != expected1 && result != expected2 {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\" or \"%v\"", input, result, expected1, expected2)
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

func TestCaseFour(t *testing.T) {
	input := "a"
	expected := "a"
	result := longestPalindrome(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

func TestCaseFive(t *testing.T) {
	input := ""
	expected := ""
	result := longestPalindrome(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

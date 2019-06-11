package main

import (
	"testing"
)

func TestCaseOneForSolutionTwo(t *testing.T) {
	input := "babad"
	expected1 := "bab"
	expected2 := "aba"
	result := longestPalindrome2(input)

	if result != expected1 && result != expected2 {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\" or \"%v\"", input, result, expected1, expected2)
	}
}

func TestCaseTwoForSolutionTwo(t *testing.T) {
	input := "cbbd"
	expected := "bb"
	result := longestPalindrome2(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

func TestCaseThreeForSolutionTwo(t *testing.T) {
	input := "12ece41abcba41"
	expected := "abcba"
	result := longestPalindrome2(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

func TestCaseFourForSolutionTwo(t *testing.T) {
	input := "a"
	expected := "a"
	result := longestPalindrome2(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

func TestCaseFiveForSolutionTwo(t *testing.T) {
	input := ""
	expected := ""
	result := longestPalindrome2(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
	}
}

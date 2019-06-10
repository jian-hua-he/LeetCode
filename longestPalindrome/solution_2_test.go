package main

import (
	"testing"
)

func TestCaseOneForSolutionTwo(t *testing.T) {
	input := "babad"
	expected := "bab"
	result := longestPalindrome2(input)

	if result != expected {
		t.Errorf("Data was incorrect, input \"%v\", got \"%v\", want \"%v\"", input, result, expected)
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

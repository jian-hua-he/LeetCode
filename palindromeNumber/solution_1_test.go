package main

import "testing"

func TestCase1IsPalindrome(t *testing.T) {
	input := 121
	got := isPalindrome(input)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase2IsPalindrome(t *testing.T) {
	input := -121
	got := isPalindrome(input)
	want := false

	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase3IsPalindrome(t *testing.T) {
	input := 10
	got := isPalindrome(input)
	want := false

	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

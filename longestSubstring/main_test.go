package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input1 := "abcabcbb"
	expected1 := 3
	len1 := lengthOfLongestSubstring(input1)
	if len1 != expected1 {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input1, len1, expected1)
	}

	input2 := "bbbbb"
	expected2 := 1
	len2 := lengthOfLongestSubstring(input2)
	if len2 != expected2 {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input2, len2, expected2)
	}

	input3 := "pwwkew"
	expected3 := 3
	len3 := lengthOfLongestSubstring(input3)
	if len3 != expected3 {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input3, len3, expected3)
	}
}

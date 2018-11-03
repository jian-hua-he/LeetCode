package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	input1 := "abcabcbb"
	expected1 := 3
	len1 := lengthOfLongestSubstring(input1)
	if len1 != expected1 {
		t.Errorf("Data was incorrect, input \"%v\", got %v, want %v", input1, len1, expected1)
	}

	input2 := "bbbbb"
	expected2 := 1
	len2 := lengthOfLongestSubstring(input2)
	if len2 != expected2 {
		t.Errorf("Data was incorrect, input \"%v\", got %v, want %v", input2, len2, expected2)
	}

	input3 := "pwwkew"
	expected3 := 3
	len3 := lengthOfLongestSubstring(input3)
	if len3 != expected3 {
		t.Errorf("Data was incorrect, input \"%v\", got %v, want %v", input3, len3, expected3)
	}

	input4 := " "
	expected4 := 1
	len4 := lengthOfLongestSubstring(input4)
	if len4 != expected4 {
		t.Errorf("Data was incorrect, input \"%v\", got %v, want %v", input4, len4, expected4)
	}

	input5 := "au"
	expected5 := 2
	len5 := lengthOfLongestSubstring(input5)
	if len5 != expected5 {
		t.Errorf("Data was incorrect, input \"%v\", got %v, want %v", input5, len5, expected5)
	}

	input6 := "dvdf"
	expected6 := 3
	len6 := lengthOfLongestSubstring(input6)
	if len6 != expected6 {
		t.Errorf("Data was incorrect, input \"%v\", got %v, want %v", input6, len6, expected6)
	}
}

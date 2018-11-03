package main

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	len1 := lengthOfLongestSubstring("abcabcbb")
	if len1 != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", len1, 3)
	}

	len2 := lengthOfLongestSubstring("bbbbb")
	if len2 != 1 {
		t.Errorf("Data was incorrect, got %v, want %v", len2, 1)
	}

	len3 := lengthOfLongestSubstring("pwwkew")
	if len3 != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", len3, 3)
	}
}

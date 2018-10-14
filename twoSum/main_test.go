package main

import "testing"

func TestTwoNum(t *testing.T) {
	t1 := []int{2, 7, 11, 15}
	r1 := twoSum(t1, 9)

	if r1[0] != 0 {
		t.Errorf("Data was incorrect, got %v, want %v", r1[0], 0)
	}
	if r1[1] != 1 {
		t.Errorf("Data was incorrect, got %v, want %v", r1[1], 1)
	}

	t2 := []int{2, 5, 10, 12}
	r2 := twoSum(t2, 14)

	if r2[0] != 0 {
		t.Errorf("Data was incorrect, got %v, want %v", r2[0], 0)
	}
	if r2[1] != 3 {
		t.Errorf("Data was incorrect, got %v, want %v", r2[1], 3)
	}
}
package main

import "testing"

func TestCase1TwoSum(t *testing.T) {
	input := []int{2, 7, 11, 15}
	got := twoSum(input, 9)

	if (got[0] != 0 || got[1] != 1) && (got[0] != 1 || got[1] != 0) {
		t.Errorf(
			"Data was incorrect, got %v, want %v or %v",
			got,
			[]int{0, 1},
			[]int{1, 0},
		)
	}
}

func TestCase2TwoSum(t *testing.T) {
	input := []int{2, 5, 10, 12}
	got := twoSum(input, 14)

	if (got[0] != 0 || got[1] != 3) && (got[0] != 3 || got[1] != 0) {
		t.Errorf(
			"Data was incorrect, got %v, want %v or %v",
			got,
			[]int{0, 3},
			[]int{3, 0},
		)
	}
}

func TestCase3TwoSum(t *testing.T) {
	input := []int{3, 2, 4}
	got := twoSum(input, 6)

	if (got[0] != 1 || got[1] != 2) && (got[0] != 2 || got[1] != 1) {
		t.Errorf(
			"Data was incorrect, got %v, want %v or %v",
			got,
			[]int{1, 2},
			[]int{2, 1},
		)
	}
}

func TestCase4TwoSum(t *testing.T) {
	input := []int{4, 6, 12, 5, 2}
	got := twoSum(input, 11)

	if (got[0] != 1 || got[1] != 3) && (got[0] != 3 || got[1] != 1) {
		t.Errorf(
			"Data was incorrect, got %v, want %v or %v",
			got,
			[]int{1, 3},
			[]int{3, 1},
		)
	}
}

func TestCase5TwoSum(t *testing.T) {
	input := []int{-1, -2, -3, -4, -5}
	got := twoSum(input, -8)

	if (got[0] != 2 || got[1] != 4) && (got[0] != 4 || got[1] != 2) {
		t.Errorf(
			"Data was incorrect, got %v, want %v or %v",
			got,
			[]int{2, 4},
			[]int{4, 2},
		)
	}
}

package main

import (
	"math"
	"testing"
)

func TestCase1ForReverse2(t *testing.T) {
	input := 123
	want := 321
	got := reverse2(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase2ForReverse2(t *testing.T) {
	input := -123
	want := -321
	got := reverse2(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase3ForReverse2(t *testing.T) {
	input := 120
	want := 21
	got := reverse2(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase4ForReverse2(t *testing.T) {
	input := math.MaxInt32 + 1
	want := 0
	got := reverse2(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase5ForReverse2(t *testing.T) {
	input := math.MinInt32 - 1
	want := 0
	got := reverse2(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase6ForReverse2(t *testing.T) {
	input := 1534236469
	want := 0
	got := reverse2(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

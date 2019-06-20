package main

import (
	"math"
	"testing"
)

func TestCase1ForReverse1(t *testing.T) {
	input := 123
	want := 321
	got := reverse1(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase2ForReverse1(t *testing.T) {
	input := -123
	want := -321
	got := reverse1(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase3ForReverse1(t *testing.T) {
	input := 120
	want := 21
	got := reverse1(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase4ForReverse1(t *testing.T) {
	input := math.MaxInt32 + 1
	want := 0
	got := reverse1(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

func TestCase5ForReverse1(t *testing.T) {
	input := math.MinInt32 - 1
	want := 0
	got := reverse1(input)
	if got != want {
		t.Errorf("Data was incorrect, input %v, got %v, want %v", input, got, want)
	}
}

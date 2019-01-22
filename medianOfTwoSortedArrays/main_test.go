package main

import "testing"

func TestTwoEmptyArray(t *testing.T) {
	input1 := []int{}
	input2 := []int{}

	result := findMedianSortedArrays(input1, input2)
	expected := 0.0
	if result != expected {
		t.Errorf("Data was incorrect, input1 \"%v\", input2 \"%v\", got %v, want %v", input1, input2, result, expected)
	}
}

func TestOneEmptyArray(t *testing.T) {
	input1 := []int{}
	input2 := []int{1, 4, 6}
	result := findMedianSortedArrays(input1, input2)
	expected := 4.0
	if result != expected {
		t.Errorf("Data was incorrect, input1 \"%v\", input2 \"%v\", got %v, want %v", input1, input2, result, expected)
	}
}

func TestCaseOne(t *testing.T) {
	input1 := []int{1, 3}
	input2 := []int{2}
	result := findMedianSortedArrays(input1, input2)
	expected := 2.0
	if result != expected {
		t.Errorf("Data was incorrect, input1 \"%v\", input2 \"%v\", got %v, want %v", input1, input2, result, expected)
	}
}

func TestCaseTwo(t *testing.T) {
	input1 := []int{1, 2}
	input2 := []int{3, 4}
	result := findMedianSortedArrays(input1, input2)
	expected := 2.5
	if result != expected {
		t.Errorf("Data was incorrect, input1 \"%v\", input2 \"%v\", got %v, want %v", input1, input2, result, expected)
	}
}

func TestCaseThree(t *testing.T) {
	input1 := []int{1, 1, 2, 5}
	input2 := []int{3, 4, 10}
	result := findMedianSortedArrays(input1, input2)
	expected := 5.0
	if result != expected {
		t.Errorf("Data was incorrect, input1 \"%v\", input2 \"%v\", got %v, want %v", input1, input2, result, expected)
	}
}

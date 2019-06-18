package main

import "testing"

func TestCase1ForConvert1(t *testing.T) {
	input := "PAYPALISHIRING"
	rows := 3
	want := "PAHNAPLSIIGYIR"
	got := convert1(input, rows)
	if got != want {
		t.Errorf("Data was incorrect, input %v, rows %v got %v, want %v", input, rows, got, want)
	}
}

func TestCase2ForConvert1(t *testing.T) {
	input := "PAYPALISHIRING"
	rows := 4
	want := "PINALSIGYAHRPI"
	got := convert1(input, rows)
	if got != want {
		t.Errorf("Data was incorrect, input %v, rows %v got %v, want %v", input, rows, got, want)
	}
}

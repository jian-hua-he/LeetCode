package main

import "testing"

func TestCase1IsMatch1(t *testing.T) {
	s := "aa"
	p := "a"
	got := isMatch(s, p)
	want := false

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase2IsMatch1(t *testing.T) {
	s := "aa"
	p := "a*"
	got := isMatch(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase3IsMatch1(t *testing.T) {
	s := "ab"
	p := ".*"
	got := isMatch(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase4IsMatch1(t *testing.T) {
	s := "aab"
	p := "c*a*b"
	got := isMatch(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase5IsMatch1(t *testing.T) {
	s := "mississippi"
	p := "mis*is*ip*."
	got := isMatch(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase6IsMatch1(t *testing.T) {
	s := "foo"
	p := "foo"
	got := isMatch(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase7IsMatch1(t *testing.T) {
	s := "aaa"
	p := "a*a"
	got := isMatch(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

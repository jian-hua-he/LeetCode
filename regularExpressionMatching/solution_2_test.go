package main

import "testing"

func TestCase1IsMatch2(t *testing.T) {
	s := "aa"
	p := "a"
	got := isMatch2(s, p)
	want := false

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase2IsMatch2(t *testing.T) {
	s := "aa"
	p := "a*"
	got := isMatch2(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase3IsMatch2(t *testing.T) {
	s := "ab"
	p := ".*"
	got := isMatch2(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase4IsMatch2(t *testing.T) {
	s := "aab"
	p := "c*a*b"
	got := isMatch2(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase5IsMatch2(t *testing.T) {
	s := "mississippi"
	p := "mis*is*p*."
	got := isMatch2(s, p)
	want := false

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase6IsMatch2(t *testing.T) {
	s := "foo"
	p := "foo"
	got := isMatch2(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

func TestCase7IsMatch2(t *testing.T) {
	s := "aaa"
	p := "a*a"
	got := isMatch2(s, p)
	want := true

	if got != want {
		t.Errorf("Data was incorrect, string \"%v\", pattern \"%v\", got %v, want %v", s, p, got, want)
	}
}

package main

import "testing"

func TestSum5Plus5(t *testing.T) {
	got := sum(5, 5)
	expected := 10

	if got != expected {
		t.Errorf("Error TestSum: got: %d\nwant: %d\n", got, expected)
	}
}

func TestSum56Plus1(t *testing.T) {
	got := sum(6, 1)
	expected := 7

	if got != expected {
		t.Errorf("Error TestSum: got: %d\nwant: %d\n", got, expected)
	}
}

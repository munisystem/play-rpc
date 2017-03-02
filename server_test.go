package main

import (
	"testing"
)

func TestFibTail(t *testing.T) {
	actual := FibTail(10)
	expected := 55
	if actual != expected {
		t.Errorf("got %v, want %v", actual, expected)
	}
}

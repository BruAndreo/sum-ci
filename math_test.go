package main

import "testing"

func TestSum(t *testing.T) {
	expected := 30
	total := sum(15, 15)

	if total != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", total, expected)
	}
}

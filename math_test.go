package main

import "testing"

func TestSum(t *testing.T) {
	expected := 30
	total := sum(15, 15)

	if total != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", total, expected)
	}
}

func TestSub(t *testing.T) {
	expected := 0
	result := sub(15, 15)

	if result != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", result, expected)
	}
}

func TestTimes(t *testing.T) {
	expected := 15
	result := times(3, 5)

	if result != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", result, expected)
	}
}

func TestSumX(t *testing.T) {
	expected := 45
	result := sumX(15, 15)

	if result != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", result, expected)
	}
}

func TestSumX2(t *testing.T) {
	expected := 60
	result := sumX2(15, 15)

	if result != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", result, expected)
	}
}

func TestSumX3(t *testing.T) {
	expected := 75
	result := sumX3(15, 15)

	if result != expected {
		t.Errorf("Result is invalid. Result: %d. Expected: %d", result, expected)
	}
}

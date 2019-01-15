package main

import "testing"

func TestGetDivisors(t *testing.T) {
	expectedResult := []int{1, 2, 3, 4, 6}
	result := getDivisors(12)

	if compareSlices(expectedResult, result) {
		t.Errorf("Expected: %v, Got: %v", expectedResult, result)
	}
}

func TestSumSlice(t *testing.T) {
	expectedResult := 16
	numbers := []int{1, 2, 3, 4, 6}
	result := sumSlice(numbers)

	if expectedResult != result {
		t.Errorf("Expected: %v, Got: %v", expectedResult, result)
	}
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

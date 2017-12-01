package main

import "testing"

func TestStringToDigits(t *testing.T) {
	tests := []struct {
		input    string
		expected []int
	}{
		{input: "", expected: []int{}},
		{input: "1122", expected: []int{1, 1, 2, 2}},
		{input: "1111", expected: []int{1, 1, 1, 1}},
		{input: "1234", expected: []int{1, 2, 3, 4}},
		{input: "91212129", expected: []int{9, 1, 2, 1, 2, 1, 2, 9}},
	}

	for _, test := range tests {
		digits := stringToDigits(test.input)

		if digits == nil {
			t.Fatalf("digits was nil")
		}

		if len(digits) != len(test.expected) {
			t.Fatalf("unequal lengths for input '%s'. expected=%d. got=%d",
				test.input, len(test.expected), len(digits))
		}

		for i, d := range digits {
			if d != test.expected[i] {
				t.Errorf("input: '%s'. digit did not match. expected=%d. got=%d.",
					test.input, test.expected[i], d)
			}
		}
	}
}

func TestCalculateFirstCaptcha(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: 0},
		{input: []int{1, 1, 2, 2}, expected: 3},
		{input: []int{1, 1, 1, 1}, expected: 4},
		{input: []int{1, 2, 3, 4}, expected: 0},
		{input: []int{9, 1, 2, 1, 2, 1, 2, 9}, expected: 9},
	}

	for _, test := range tests {
		sum := calculateFirstCaptcha(test.input)
		if sum != test.expected {
			t.Errorf("for %v, captcha sum didn't match. expected=%d. got=%d",
				test.input, test.expected, sum)
		}
	}
}

func TestCalculateSecondCaptcha(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{}, expected: 0},
		{input: []int{1, 2, 1, 2}, expected: 6},
		{input: []int{1, 2, 2, 1}, expected: 0},
		{input: []int{1, 2, 3, 4, 2, 5}, expected: 4},
		{input: []int{1, 2, 3, 1, 2, 3}, expected: 12},
		{input: []int{1, 2, 1, 3, 1, 4, 1, 5}, expected: 4},
	}

	for _, test := range tests {
		sum := calculateSecondCaptcha(test.input)
		if sum != test.expected {
			t.Errorf("for %v, captcha sum didn't match. expected=%d. got=%d",
				test.input, test.expected, sum)
		}
	}
}

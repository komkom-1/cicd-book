package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{name: "positive even", input: 10, expected: "even"},
		{name: "positive odd", input: 9, expected: "odd"},
		{name: "zero", input: 0, expected: "even"},
		{name: "negative even", input: -10, expected: "even"},
		{name: "negative odd", input: -9, expected: "odd"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := EvenOrOdd(test.input)
			if actual != test.expected {
				t.Errorf("EvenOrOdd(%d) = %q; expected %q", test.input, actual, test.expected)
			}
		})
	}
}

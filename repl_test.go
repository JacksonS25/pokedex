package main

import (
	"fmt"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Alabaster Canada Dry misfotune fortune ABSOLUTE power itty BiTy Lving space",
			expected: []string{"alabaster", "canada", "dry", "misfotune", "fortune", "absolute", "power", "itty", "bity", "lving", "space"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			fmt.Printf("Length mismatch: expected %d, got %d\n", len(c.expected), len(actual))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				fmt.Printf("Word mismatch at index %d: expected %q, got %q\n", i, expectedWord, word)
			}
		}
	}

}

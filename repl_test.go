package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello moto",
			expected: []string{"hello", "moto"},
		},
		{
			input:    "  WAGWAN mi chaMpiOn  ",
			expected: []string{"wagwan", "mi", "champion"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("got string slice len: %d\nexpected len: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("%s does not equal expected word: %s", word, expectedWord)
			}
		}
	}
}

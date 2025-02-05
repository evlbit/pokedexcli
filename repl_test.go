package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander  BulbASaur PIKACHU  ",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	errMsg := "Input: %v\nExpected: %v\nActual: %v"

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("lengths mismatch\n"+errMsg, c.input, c.expected, actual)
			continue
		}

		match := true
		for i := range actual {
			if actual[i] != c.expected[i] {
				match = false
				break
			}
		}

		if !match {
			t.Errorf(errMsg, c.input, c.expected, actual)
		}
	}
}

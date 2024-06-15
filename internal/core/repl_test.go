package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HELLO world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		actualLength := len(actual)
		expectedLength := len(cs.expected)
		if actualLength != expectedLength {
			t.Errorf("Unequal lengths: %v, %v", actualLength, expectedLength)
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			if actualWord != expectedWord {
				t.Errorf("Unequal words: %v, %v", actualWord, expectedWord)
			}
		}
	}
}

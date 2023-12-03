package main

import "testing"

type Test struct {
	input    string
	expected string
}

type Tests = []Test

func TestUnpackString(t *testing.T) {
	tests := &Tests{
		{"a4bc2d5e", "aaaabccddddde"},
		{"abcd", "abcd"},
		{"45", ""},
		{"", ""},
		{"qwe\\4\\5", "qwe45"},
		{"qwe\\45", "qwe44444"},
		{"qwe\\\\5", "qwe\\\\\\\\\\"},
	}

	application := &Application{}

	for _, test := range *tests {
		result := application.UnpackString(test.input)

		if result != test.expected {
			t.Errorf("Input: %s, Expected: %s, Result: %s", test.input, test.expected, result)
		}
	}
}

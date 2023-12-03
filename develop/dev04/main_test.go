package main

import (
	"reflect"
	"testing"
)

type Test struct {
	words    []string
	expected map[string][]string
}

type Tests = []Test

func TestFindAnagrams(t *testing.T) {
	application := &Application{}

	tests := Tests{
		{
			words:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
			expected: map[string][]string{"акптя": {"пятак", "пятка", "тяпка"}, "иклост": {"листок", "слиток", "столик"}},
		},
		{
			words:    []string{"мама", "амам", "тот", "тот"},
			expected: map[string][]string{"аамм": {"амам", "мама"}, "отт": {"тот", "тот"}},
		},
	}

	for _, test := range tests {
		result := application.FindAnagrams(test.words)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("for words %v expected %v but got %v", test.words, test.expected, result)
		}
	}
}

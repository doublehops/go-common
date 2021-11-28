package strings

import (
	"testing"
)

func TestSliceContains(t *testing.T) {

	Map := []string{
		"apple",
		"banana",
		"mandarin",
	}

	var tests = []struct {
		Name     string
		Value    string
		Expected bool
	}{
		{"key not found", "grape", false},
		{"not found", "apple", true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			received := sliceContains(test.Value, Map)
			if received != test.Expected {
				t.Errorf("expected: %t; received: %t", test.Expected, received)
			}
		})
	}
}

func TestKeyStringExists(t *testing.T) {

	Map := map[string]interface{}{
		"apple":    "foo",
		"banana":   123,
		"mandarin": 123,
	}

	var tests = []struct {
		Name     string
		Value    string
		Expected bool
	}{
		{"key not found", "grape", false},
		{"not found", "apple", true},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			received := KeyStringExists(test.Value, Map)
			if received != test.Expected {
				t.Errorf("expected: %t; received: %t", test.Expected, received)
			}
		})
	}
}

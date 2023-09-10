package str

import (
	"testing"
)

func TestSliceContains(t *testing.T) {
	fruits := []string{"apple", "banana", "mandarin"}
	if !SliceContains("banana", fruits) {
		t.Errorf("banana not found in fruits. expected: %t; received: %t", true, false)
	}

	cars := []string{"Mazda", "Ford", "Holden"}
	if SliceContains("Nissan", cars) {
		t.Errorf("Nissan found in cars. expected: %t; received: %t", false, true)
	}

	type testString string

	drinks := []testString{"coffee", "tea", "water"}
	if !SliceContains("water", drinks) {
		t.Errorf("water not found in drinks. expected: %t; received: %t", true, false)
	}

	OSes := []testString{"Linux", "Windows", "Macos"}
	if SliceContains("FreeBSD", OSes) {
		t.Errorf("FreeBSD found in OSes. expected: %t; received: %t", false, true)
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

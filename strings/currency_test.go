package strings

import "testing"

func TestAddSeparator(t *testing.T) {
	var tests = []struct{
		Name string
		Value float64
		Expected string
	}{
		{"hundreds", 134.81122, "134.81"},
		{"thousands", 1234.8012, "1,234.80"},
		{"millions", 12345678, "12,345,678.00"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			received := ConvertToCurrency(test.Value)
			if received != test.Expected {
				t.Errorf("expected: %s; received: %s", test.Expected, received)
			}
		})
	}
}

func TestAddEuropeanSeparator(t *testing.T) {
	var tests = []struct{
		Name string
		Value float64
		Expected string
	}{
		{"hundreds", 134.81, "134,81"},
		{"thousands", 1234.8012, "1.234,80"},
		{"millions", 12345678.80, "12.345.678,80"},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			received := ConvertToCurrency(test.Value, ".", ",")
			if received != test.Expected {
				t.Errorf("expected: %s; received: %s", test.Expected, received)
			}
		})
	}
}

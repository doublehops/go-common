package strings

import "testing"

func TestAddSeparator(t *testing.T) {
	var value float64
	value = 1234567.89

	//type subTestStruct struct{
	//	Name string
	//	Value float64
	//	Expected string
	//}
	//
	//subTests := []subTestStruct{
	//
	//
	//}

	expected := "1,234,567.89"
	foundVal := AddSeparator(value)

	if foundVal != expected {
		t.Fatalf("Value incorrect. Expected: %s; found: %s", expected, foundVal)
	}
}

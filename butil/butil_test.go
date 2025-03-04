package butil

import "testing"

// TestPrintReadableTypeValue tests the PrintReadableTypeValue function.
func TestPrintReadableTypeValue(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
	}{
		{42, "(int)(42)"},
		{"hello", "(string)(hello)"},
		{3.14, "(float64)(3.14)"},
		{new(int), "(*int)(0)"},
		{new(string), "(*string)()"},
		{new(float64), "(*float64)(0)"},
		{(*int)(nil), "(*int)(nil)"},
		{(*string)(nil), "(*string)(nil)"},
		{(*float64)(nil), "(*float64)(nil)"},
		{&struct{ A int }{A: 1}, "(*struct { A int })({A:1})"},
		//{&struct{ A *int }{A: new(int)}, "(*struct { A *int })(struct { A *int }{A:(int)(0)})"}, // TODO: fix this test case
	}

	for _, test := range tests {
		actual := PrintReadableTypeValue(test.input)
		if actual != test.expected {
			t.Errorf("PrintReadableTypeValue(%v) = %s; expected %s", test.input, actual, test.expected)
		}
	}
}

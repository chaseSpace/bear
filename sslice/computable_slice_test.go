package sslice

import "testing"

func TestComputableSlice_Sum_EmptySlice_ReturnsZero(t *testing.T) {
	slice := NewComputableSlice[int]()
	sum := slice.Sum()
	if sum != 0 {
		t.Errorf("Expected sum to be 0, got %d", sum)
	}
}

func TestComputableSlice_Sum_SingleElement_ReturnsElement(t *testing.T) {
	slice := NewComputableSlice(5)
	sum := slice.Sum()
	if sum != 5 {
		t.Errorf("Expected sum to be 5, got %d", sum)
	}
}

func TestComputableSlice_Sum_MultipleElements_ReturnsCorrectSum(t *testing.T) {
	slice := NewComputableSlice(1, 2, 3, 4)
	sum := slice.Sum()
	if sum != 10 {
		t.Errorf("Expected sum to be 10, got %d", sum)
	}
}

func TestComputableSlice_Sum_Floats_ReturnsCorrectSum(t *testing.T) {
	slice := NewComputableSlice(1.5, 2.5, 3.0)
	sum := slice.Sum()
	if sum != 7.0 {
		t.Errorf("Expected sum to be 7.0, got %f", sum)
	}
}

func TestComputableSlice_Max_MultipleElements_ReturnsMax(t *testing.T) {
	os := NewComputableSlice(1, 3, 2, 5, 4)
	if os.Max() != 5 {
		t.Errorf("Max() = %v, want %v", os.Max(), 5)
	}
	// does make influence on the original slice
	if !os.Equal(NewComputableSlice(1, 3, 2, 5, 4)) {
		t.Errorf("the origin slice has been changed")
	}
}

func TestComputableSlice_Min_MultipleElements_ReturnsMax(t *testing.T) {
	os := NewComputableSlice(1, 3, 2, 5, 4)
	if os.Min() != 1 {
		t.Errorf("Max() = %v, want %v", os.Max(), 1)
	}
	// does make influence on the original slice
	if !os.Equal(NewComputableSlice(1, 3, 2, 5, 4)) {
		t.Errorf("the origin slice has been changed")
	}
}

// TestAvg tests the Avg method of the ComputableSlice.
func TestAvg(t *testing.T) {
	tests := []struct {
		name     string
		input    *ComputableSlice[int]
		expected int
	}{
		{"NonEmptySlice", NewComputableSlice(1, 2, 3, 4, 5), 3},
		{"EmptySlice", NewComputableSlice[int](), 0},
		{"SingleElementSlice", NewComputableSlice(10), 10},
		{"NegativeNumbers", NewComputableSlice(-1, -2, -3), -2},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.Avg()
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}

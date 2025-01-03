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

func TestComputableSlice_Sum_ComplexNumbers_ReturnsCorrectSum(t *testing.T) {
	slice := NewComputableSlice(complex(1, 2), complex(3, 4))
	sum := slice.Sum()
	expected := complex(4, 6)
	if sum != expected {
		t.Errorf("Expected sum to be %v, got %v", expected, sum)
	}
}

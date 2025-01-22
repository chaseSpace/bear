package sslice

import (
	"reflect"
	"testing"
)

// TestSort_OrderedSlice_Int tests the Sort method of OrderedSlice.
func TestSort_OrderedSlice_Int(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		desc     []bool
		expected []int
	}{
		{"Sort Ascending", []int{5, 2, 9, 1, 5, 6}, []bool{false}, []int{1, 2, 5, 5, 6, 9}},
		{"Sort Descending", []int{5, 2, 9, 1, 5, 6}, []bool{true}, []int{9, 6, 5, 5, 2, 1}},
		{"Sort Default Ascending", []int{5, 2, 9, 1, 5, 6}, nil, []int{1, 2, 5, 5, 6, 9}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := NewOrderedSlice(test.input...)
			is.Sort(test.desc...)
			if !reflect.DeepEqual(is.slice.data, test.expected) {
				t.Errorf("Sort(%v) = %v, want %v", test.input, is.slice.data, test.expected)
			}
		})
	}
}

// TestSort_OrderedSlice_String tests the Sort method of OrderedSlice.
func TestSort_OrderedSlice_String(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		desc     []bool
		expected []string
	}{
		{"Sort Ascending", []string{"c", "a", "b"}, []bool{false}, []string{"a", "b", "c"}},
		{"Sort Descending", []string{"c", "a", "b"}, []bool{true}, []string{"c", "b", "a"}},
		{"Sort Default Ascending", []string{"c", "a", "b"}, nil, []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := NewOrderedSlice(test.input...)
			is.Sort(test.desc...)
			if !reflect.DeepEqual(is.slice.data, test.expected) {
				t.Errorf("Sort(%v) = %v, want %v", test.input, is.slice.data, test.expected)
			}
		})
	}
}

// TestSort_OrderedSlice_Byte tests the Sort method of OrderedSlice.
func TestSort_OrderedSlice_Byte(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		desc     []bool
		expected []byte
	}{
		{"Sort Ascending", []byte{5, 2, 9, 1, 5, 6}, []bool{false}, []byte{1, 2, 5, 5, 6, 9}},
		{"Sort Descending", []byte{5, 2, 9, 1, 5, 6}, []bool{true}, []byte{9, 6, 5, 5, 2, 1}},
		{"Sort Default Ascending", []byte{5, 2, 9, 1, 5, 6}, nil, []byte{1, 2, 5, 5, 6, 9}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			is := NewOrderedSlice(test.input...)
			is.Sort(test.desc...)
			if !reflect.DeepEqual(is.slice.data, test.expected) {
				t.Errorf("Sort(%v) = %v, want %v", test.input, is.slice.data, test.expected)
			}
		})
	}
}

func TestFilter_OrderedSlice_ReturnsFilteredSlice(t *testing.T) {
	os := NewOrderedSlice([]int{1, 2, 3, 4, 5}...)
	filtered := os.Filter(func(x int) bool { return x > 2 })
	if expected := []int{1, 2}; !reflect.DeepEqual(filtered.slice.data, expected) {
		t.Errorf("Filter(func(x int) bool { return x > 2 }) = %v, want %v", filtered.slice.data, expected)
	}
}

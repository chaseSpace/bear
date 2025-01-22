package sslice

import (
	"reflect"
	"testing"
)

// TestClone_EmptySlice_ShouldReturnEmptySlice tests cloning an empty slice.
func TestClone_EmptySlice_ShouldReturnEmptySlice(t *testing.T) {
	original := New[int]()
	cloned := original.Clone()

	if !reflect.DeepEqual(cloned.data, original.data) {
		t.Errorf("Expected cloned slice to be empty, got %v", cloned.data)
	}
}

// TestClone_NonEmptySlice_ShouldReturnClonedSlice tests cloning a non-empty slice.
func TestClone_NonEmptySlice_ShouldReturnClonedSlice(t *testing.T) {
	original := New(1, 2, 3)
	cloned := original.Clone()

	if !reflect.DeepEqual(cloned.data, original.data) {
		t.Errorf("Expected cloned slice to be %v, got %v", original.data, cloned.data)
	}
}

// TestClone_SliceWithDifferentTypes_ShouldReturnClonedSlice tests cloning a slice with different data types.
func TestClone_SliceWithDifferentTypes_ShouldReturnClonedSlice(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	original := New([]Person{{"Alice", 30}, {"Bob", 25}}...)
	cloned := original.Clone()

	if !reflect.DeepEqual(cloned.data, original.data) {
		t.Errorf("Expected cloned slice to be %v, got %v", original.data, cloned.data)
	}
}

// TestFilter_EmptySlice_ShouldReturnEmptySlice tests the Filter method with an empty slice.
func TestFilter_EmptySlice_ShouldReturnEmptySlice(t *testing.T) {
	s := New[int]()
	s.Filter(func(x int) bool { return x > 0 })
	if len(s.data) != 0 {
		t.Errorf("Expected empty slice, got %v", s.data)
	}
}

// TestFilter_AlwaysTruePredicate_ShouldReturnSameSlice tests the Filter method with a predicate that always returns true.
func TestFilter_AlwaysTruePredicate_ShouldReturnSameSlice(t *testing.T) {
	s := New(1, 2, 3)
	s.Filter(func(x int) bool { return true })
	expected := []int{}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestFilter_AlwaysFalsePredicate_ShouldReturnEmptySlice tests the Filter method with a predicate that always returns false.
func TestFilter_AlwaysFalsePredicate_ShouldReturnEmptySlice(t *testing.T) {
	s := New(1, 2, 3)
	s.Filter(func(x int) bool { return false })
	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestFilter_MixedPredicate_ShouldReturnFilteredSlice tests the Filter method with a predicate that returns mixed results.
func TestFilter_MixedPredicate_ShouldReturnFilteredSlice(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	s.Filter(func(x int) bool { return x%2 == 0 })
	expected := []int{1, 3, 5}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestFilter_ValueDependentPredicate_ShouldReturnFilteredSlice tests the Filter method with a predicate that depends on element values.
func TestFilter_ValueDependentPredicate_ShouldReturnFilteredSlice(t *testing.T) {
	s := New("apple", "banana", "cherry")
	s.Filter(func(x string) bool { return len(x) > 5 })
	expected := []string{"apple"}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestUnique_NoDuplicates_ShouldRemainUnchanged tests the Unique method on a slice with no duplicates.
func TestUnique_NoDuplicates_ShouldRemainUnchanged(t *testing.T) {
	s := New(1, 2, 3)
	expected := []int{1, 2, 3}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestUnique_AllDuplicates_ShouldRemoveAllButOne tests the Unique method on a slice with all elements the same.
func TestUnique_AllDuplicates_ShouldRemoveAllButOne(t *testing.T) {
	s := New(1, 1, 1)
	expected := []int{1}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestUnique_MixedElements_ShouldRemoveDuplicates tests the Unique method on a slice with mixed elements.
func TestUnique_MixedElements_ShouldRemoveDuplicates(t *testing.T) {
	s := New(1, 2, 2, 3, 4, 4, 5)
	expected := []int{1, 2, 3, 4, 5}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestUnique_StringSlice_ShouldHandleStrings tests the Unique method on a slice of strings.
func TestUnique_StringSlice_ShouldHandleStrings(t *testing.T) {
	s := New("a", "b", "a", "c")
	expected := []string{"a", "b", "c"}
	s = s.Unique()
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got: %v", expected, s.data)
	}
}

// TestReduce_EmptySlice_ReturnsZeroValue tests the Reduce method with an empty slice.
func TestReduce_EmptySlice_ReturnsZeroValue(t *testing.T) {
	s := New[int]()
	result := s.Reduce(func(x, y int) int { return x + y })
	if result != 0 {
		t.Errorf("Expected zero value, got %d", result)
	}
}

// TestReduce_SingleElement_ReturnsElement tests the Reduce method with a single element.
func TestReduce_SingleElement_ReturnsElement(t *testing.T) {
	s := New(5)
	result := s.Reduce(func(x, y int) int { return x + y })
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// TestReduce_MultipleElements_SumsCorrectly tests the Reduce method with multiple elements.
func TestReduce_MultipleElements_SumsCorrectly(t *testing.T) {
	s := New(1, 2, 3, 4)
	result := s.Reduce(func(x, y int) int { return x + y })
	if result != 10 {
		t.Errorf("Expected 10, got %d", result)
	}
}

// TestReduce_MultipleElements_ProductCorrectly tests the Reduce method with a product function.
func TestReduce_MultipleElements_ProductCorrectly(t *testing.T) {
	s := New(1, 2, 3, 4)
	result := s.Reduce(func(x, y int) int { return x * y })
	if result != 24 {
		t.Errorf("Expected 24, got %d", result)
	}
}

// TestReduce_MultipleElements_MaxCorrectly tests the Reduce method with a max function.
func TestReduce_MultipleElements_MaxCorrectly(t *testing.T) {
	s := New(1, 5, 3, 4)
	result := s.Reduce(func(x, y int) int {
		if x > y {
			return x
		}
		return y
	})
	if result != 5 {
		t.Errorf("Expected 5, got %d", result)
	}
}

// TestReverse_EmptySlice_ShouldRemainEmpty tests reversing an empty slice.
func TestReverse_EmptySlice_ShouldRemainEmpty(t *testing.T) {
	s := New[int]()
	s.Reverse()
	if len(s.data) != 0 {
		t.Errorf("Expected empty slice, got: %v", s.data)
	}
}

// TestReverse_SingleElement_ShouldRemainUnchanged tests reversing a slice with a single element.
func TestReverse_SingleElement_ShouldRemainUnchanged(t *testing.T) {
	s := New(1)
	s.Reverse()
	expected := []int{1}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestReverse_EvenNumberOfElements_ShouldReverseCorrectly tests reversing a slice with an even number of elements.
func TestReverse_EvenNumberOfElements_ShouldReverseCorrectly(t *testing.T) {
	s := New(1, 2, 3, 4)
	s.Reverse()
	expected := []int{4, 3, 2, 1}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestReverse_OddNumberOfElements_ShouldReverseCorrectly tests reversing a slice with an odd number of elements.
func TestReverse_OddNumberOfElements_ShouldReverseCorrectly(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	s.Reverse()
	expected := []int{5, 4, 3, 2, 1}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

// TestReverse_DuplicateElements_ShouldReverseCorrectly tests reversing a slice with duplicate elements.
func TestReverse_DuplicateElements_ShouldReverseCorrectly(t *testing.T) {
	s := New(1, 2, 2, 1)
	s.Reverse()
	expected := []int{1, 2, 2, 1}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s.data)
	}
}

func TestJoin_MultipleElements_ReturnsJoinedString(t *testing.T) {
	os := New(1, 2, 3)
	result := os.Join(",")
	if result != "1,2,3" {
		t.Errorf("Expected '1,2,3', got %s", result)
	}
}

func TestJoin_DifferentSeparator_ReturnsJoinedStringWithSeparator(t *testing.T) {
	os := New(1, 2, 3)
	result := os.Join("-")
	if result != "1-2-3" {
		t.Errorf("Expected '1-2-3', got %s", result)
	}
}

func TestJoin_StringSlice_ReturnsJoinedString(t *testing.T) {
	os := New("a", "b", "c")
	result := os.Join(",")
	if result != "a,b,c" {
		t.Errorf("Expected 'a,b,c', got %s", result)
	}
}

// TestPopLeft_EmptySlice_NoChange tests PopLeft on an empty slice.
func TestPopLeft_EmptySlice_NoChange(t *testing.T) {
	s := New[int]()
	result := s.PopLeft()
	if !reflect.DeepEqual(result.data, []int{}) {
		t.Errorf("Expected empty slice, got %v", result.data)
	}
}

// TestPopLeft_NonEmptySlice_RemovesFirstElement tests PopLeft on a non-empty slice.
func TestPopLeft_NonEmptySlice_RemovesFirstElement(t *testing.T) {
	s := New(1, 2, 3)
	result := s.PopLeft()
	expected := []int{2, 3}
	if !reflect.DeepEqual(result.data, expected) {
		t.Errorf("Expected %v, got %v", expected, result.data)
	}
}

// TestPopRight_EmptySlice_NoChange tests PopRight on an empty slice.
func TestPopRight_EmptySlice_NoChange(t *testing.T) {
	s := &Slice[int]{data: []int{}}
	result := s.PopRight()
	if !reflect.DeepEqual(result.data, []int{}) {
		t.Errorf("Expected empty slice, got %v", result.data)
	}
}

// TestPopRight_NonEmptySlice_RemovesLastElement tests PopRight on a non-empty slice.
func TestPopRight_NonEmptySlice_RemovesLastElement(t *testing.T) {
	s := New(1, 2, 3)
	result := s.PopRight()
	expected := []int{1, 2}
	if !reflect.DeepEqual(result.data, expected) {
		t.Errorf("Expected %v, got %v", expected, result.data)
	}
}

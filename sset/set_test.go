package sset

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"strings"
	"testing"
)

// TestAdd_EmptySet_AddsElements tests adding elements to an empty set.
func TestAdd_EmptySet_AddsElements(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	expected := map[int]struct{}{1: {}, 2: {}, 3: {}}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s)
	}
}

// TestAdd_NonEmptySet_AddsNewElements tests adding new elements to a non-empty set.
func TestAdd_NonEmptySet_AddsNewElements(t *testing.T) {
	s := New(1, 2)
	s.Add(3, 4)
	expected := map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s)
	}
}

// TestAdd_ExistingElements_NoDuplicates tests adding existing elements to a set.
func TestAdd_ExistingElements_NoDuplicates(t *testing.T) {
	s := New(1, 2)
	s.Add(1, 2)
	expected := map[int]struct{}{1: {}, 2: {}}
	if !reflect.DeepEqual(s.data, expected) {
		t.Errorf("Expected %v, got %v", expected, s)
	}
}

// TestClone_EmptySet_ReturnsEmptySet tests cloning an empty set.
func TestClone_EmptySet_ReturnsEmptySet(t *testing.T) {
	set := New[int]()
	clonedSet := set.Clone()

	if !clonedSet.IsEmpty() {
		t.Errorf("Expected cloned set to be empty, but it is not")
	}
}

// TestClone_NonEmptySet_ReturnsClonedSet tests cloning a non-empty set.
func TestClone_NonEmptySet_ReturnsClonedSet(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)

	clonedSet := set.Clone()

	if clonedSet.Size() != set.Size() {
		t.Errorf("Expected cloned set size to be %d, but got %d", set.Size(), clonedSet.Size())
	}

	for _, item := range set.Slice() {
		if _, ok := clonedSet.data[item]; !ok {
			t.Errorf("Expected cloned set to contain %d, but it does not", item)
		}
	}
}

// TestClone_ModifyClonedSet_DoesNotAffectOriginalSet tests modifying a cloned set does not affect the original set.
func TestClone_ModifyClonedSet_DoesNotAffectOriginalSet(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)

	clonedSet := set.Clone()
	clonedSet.Add(3)

	if _, ok := set.data[3]; ok {
		t.Errorf("Expected original set not to contain 3 after modifying cloned set")
	}
}

// TestClone_CloneOfClone_ReturnsSameElements tests cloning a cloned set returns the same elements.
func TestClone_CloneOfClone_ReturnsSameElements(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)

	clonedSet := set.Clone()
	clonedClonedSet := clonedSet.Clone()

	if !reflect.DeepEqual(clonedSet.data, clonedClonedSet.data) {
		t.Errorf("Expected cloned set and cloned cloned set to be equal")
	}
}

func TestSet_Filter_AllElementsSatisfyCondition_ReturnsSameSet(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	filteredSet := set.Filter(func(x int) bool { return x > 0 })
	if !reflect.DeepEqual(set.data, filteredSet.data) {
		t.Errorf("Expected filtered set shoulbe be same as original set")
	}
}

func TestSet_Filter_NoElementsSatisfyCondition_ReturnsEmptySet(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	filteredSet := set.Filter(func(x int) bool { return x > 3 })
	if !reflect.DeepEqual(filteredSet.data, map[int]struct{}{}) {
		t.Errorf("Expected filtered set to be empty")
	}
}

func TestSet_Filter_MixedConditions_RemovesUnsatisfyingElements(t *testing.T) {
	set := New[int]()
	set.Add(1)
	set.Add(2)
	set.Add(3)
	filteredSet := set.Filter(func(x int) bool { return x > 1 })
	expectedSet := New[int]()
	expectedSet.Add(2)
	expectedSet.Add(3)
	if !reflect.DeepEqual(filteredSet.data, expectedSet.data) {
		t.Errorf("Expected filtered set to be %v, got %v", expectedSet.data, filteredSet.data)
	}
}

func TestClear_NonEmptySet_ClearsData(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	assert.NotEmpty(t, s.data, "The set should not be empty initially")
	s.Clear()
	assert.Empty(t, s.data, "The set should be empty after Clear")
}

func TestClear_ReturnsSameInstance(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)
	clearedSet := s.Clear()
	assert.Same(t, s, clearedSet, "The returned set should be the same instance")
}

func TestDelete_EmptyInput_NoChange(t *testing.T) {
	s := New(1, 2, 3)
	s.Delete()
	assert.Equal(t, 3, s.Size())
}

func TestDelete_SingleElement_Remove(t *testing.T) {
	s := New(1, 2, 3)
	s.Delete(2)
	assert.Equal(t, 2, s.Size())
	assert.False(t, s.Has(2))
}

func TestDelete_MultipleElements_Remove(t *testing.T) {
	s := New(1, 2, 3, 4, 5)
	s.Delete(2, 4)
	assert.Equal(t, 3, s.Size())
	assert.False(t, s.Has(2))
	assert.False(t, s.Has(4))
}

func TestDelete_NonExistentElement_NoChange(t *testing.T) {
	s := New(1, 2, 3)
	s.Delete(4)
	assert.Equal(t, 3, s.Size())
}

func TestDelete_AllElements_EmptySet(t *testing.T) {
	s := New(1, 2, 3)
	s.Delete(1, 2, 3)
	assert.Equal(t, 0, s.Size())
}

func TestForEach_EmptySet_NoFunctionCall(t *testing.T) {
	set := New[int]()
	var called bool
	set.ForEach(func(int) {
		called = true
	})
	if called {
		t.Error("Expected no function call for an empty set")
	}
}

func TestForEach_SingleElementSet_FunctionCalledOnce(t *testing.T) {
	set := New[int]()
	set.Add(1)
	var count int
	set.ForEach(func(int) {
		count++
	})
	if count != 1 {
		t.Errorf("Expected function to be called once, got %d", count)
	}
}

func TestForEach_MultipleElementsSet_FunctionCalledForEachElement(t *testing.T) {
	set := New[int]()
	set.Add(1).Add(2).Add(3)
	var count int
	set.ForEach(func(int) {
		count++
	})
	if count != 3 {
		t.Errorf("Expected function to be called three times, got %d", count)
	}
}

func TestForEach_ModifyDuringIteration_NoModification(t *testing.T) {
	set := New[int]()
	set.Add(1).Add(2)
	var result []int
	set.ForEach(func(item int) {
		result = append(result, item)
	})
	if !reflect.DeepEqual(result, []int{1, 2}) {
		t.Errorf("Expected result to have 2 elements, got %d", len(result))
	}
}

func TestForEach_ChainableMethod_ReturnsSelf(t *testing.T) {
	set := New[int]()
	set.Add(1)
	result := set.ForEach(func(int) {}).Add(2)
	if result != set {
		t.Error("Expected ForEach to return the same set instance")
	}
}

func TestMap_EmptySet_ReturnsEmptySet(t *testing.T) {
	s := New[int]()
	s.Map(func(x int) int { return x + 1 })
	if !s.IsEmpty() {
		t.Errorf("Expected empty set, got: %v", s)
	}
}

func TestMap_SingleElementSet_TransformsCorrectly(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Map(func(x int) int { return x + 1 })
	expected := New[int]()
	expected.Add(2)
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected set: %v, got: %v", expected, s)
	}
}

func TestMap_MultipleElements_TransformsCorrectly(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	s.Map(func(x int) int { return x + 1 })
	expected := New[int]()
	expected.Add(2, 3, 4)
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected set: %v, got: %v", expected, s)
	}
}

func TestMap_ComplexTransformation_TransformsCorrectly(t *testing.T) {
	s := New[string]()
	s.Add("hello", "world")
	s.Map(func(x string) string { return strings.ToUpper(x) })
	expected := New[string]()
	expected.Add("HELLO", "WORLD")
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected set: %v, got: %v", expected, s)
	}
}

func TestMap_IdentityFunction_NoChange(t *testing.T) {
	s := New[int]()
	s.Add(1, 2, 3)
	s.Map(func(x int) int { return x })
	expected := New[int]()
	expected.Add(1, 2, 3)
	if !reflect.DeepEqual(s, expected) {
		t.Errorf("Expected set: %v, got: %v", expected, s)
	}
}

func TestIntersect_EmptySets_ReturnsEmptySet(t *testing.T) {
	set1 := New[int]()
	set2 := New[int]()
	result := set1.Intersect(set2)
	assert.Empty(t, result.data)
}

func TestIntersect_OneEmptySet_ReturnsEmptySet(t *testing.T) {
	set1 := New[int]()
	set2 := New[int]()
	set2.Add(1)
	result := set1.Intersect(set2)
	assert.Empty(t, result.data)
}

func TestIntersect_NoCommonElements_ReturnsEmptySet(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(3)
	set2.Add(4)
	result := set1.Intersect(set2)
	assert.Empty(t, result.data)
}

func TestIntersect_AllCommonElements_ReturnsSameSet(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(1)
	set2.Add(2)
	result := set1.Intersect(set2)
	assert.Equal(t, set1.data, result.data)
}

func TestIntersect_PartialCommonElements_ReturnsCommonElements(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(2)
	set2.Add(3)
	result := set1.Intersect(set2)
	expected := New[int]()
	expected.Add(2)
	assert.Equal(t, expected.data, result.data)
}

func TestUnion_EmptySets_ReturnsEmptySet(t *testing.T) {
	set1 := New[int]()
	set2 := New[int]()
	result := set1.Union(set2)
	assert.Empty(t, result.data)
}

func TestUnion_OneEmptySet_ReturnsNonEmptySet(t *testing.T) {
	set1 := New[int]()
	set2 := New[int]()
	set2.Add(1)
	result := set1.Union(set2)
	assert.Equal(t, map[int]struct{}{1: {}}, result.data)
}

func TestUnion_BothSetsHaveElements_ReturnsUnion(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(2)
	set2.Add(3)
	result := set1.Union(set2)
	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}}, result.data)
}

func TestUnion_OverlappingElements_ReturnsUnion(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(2)
	set2.Add(3)
	result := set1.Union(set2)
	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}}, result.data)
}

func TestUnion_NoOverlappingElements_ReturnsUnion(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(3)
	set2.Add(4)
	result := set1.Union(set2)
	assert.Equal(t, map[int]struct{}{1: {}, 2: {}, 3: {}, 4: {}}, result.data)
}

func TestIsSubsetOf_EmptySet_ShouldReturnTrue(t *testing.T) {
	s := New[int]()
	other := New[int]()
	assert.True(t, s.IsSubsetOf(other), "Empty set should be a subset of any set, including another empty set")
}

func TestIsSubsetOf_NonEmptySubset_ShouldReturnTrue(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)

	other := New[int]()
	other.Add(1)
	other.Add(2)
	other.Add(3)

	assert.True(t, s.IsSubsetOf(other), "Set s is a subset of other")
}

func TestIsSubsetOf_NonSubset_ShouldReturnFalse(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(4)

	other := New[int]()
	other.Add(1)
	other.Add(2)
	other.Add(3)

	assert.False(t, s.IsSubsetOf(other), "Set s is not a subset of other")
}

func TestIsSubsetOf_OtherEmpty_ShouldReturnFalse(t *testing.T) {
	s := New[int]()
	s.Add(1)

	other := New[int]()

	assert.False(t, s.IsSubsetOf(other), "Non-empty set should not be a subset of an empty set")
}

func TestIsSubsetOf_EqualSets_ShouldReturnTrue(t *testing.T) {
	s := New[int]()
	s.Add(1)
	s.Add(2)

	other := New[int]()
	other.Add(1)
	other.Add(2)

	assert.True(t, s.IsSubsetOf(other), "Equal sets should be subsets of each other")
}

// TestDiff_EmptySets_ShouldReturnEmptySet tests the Diff method when both sets are empty.
func TestDiff_EmptySets_ShouldReturnEmptySet(t *testing.T) {
	set1 := New[int]()
	set2 := New[int]()
	result := set1.Diff(set2)
	if !result.IsEmpty() {
		t.Errorf("Expected an empty set, got: %v", result)
	}
}

// TestDiff_OneEmptySet_ShouldReturnNonEmptySet tests the Diff method when one set is empty.
func TestDiff_OneEmptySet_ShouldReturnNonEmptySet(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set2 := New[int]()
	result := set1.Diff(set2)
	expected := New[int]()
	expected.Add(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestDiff_SameSets_ShouldReturnEmptySet tests the Diff method when both sets are the same.
func TestDiff_SameSets_ShouldReturnEmptySet(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(1)
	set2.Add(2)
	result := set1.Diff(set2)
	if !result.IsEmpty() {
		t.Errorf("Expected an empty set, got: %v", result)
	}
}

// TestDiff_DifferentSets_ShouldReturnAllElementsOfFirstSet tests the Diff method when sets are different.
func TestDiff_DifferentSets_ShouldReturnAllElementsOfFirstSet(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set2 := New[int]()
	set2.Add(3)
	set2.Add(4)
	result := set1.Diff(set2)
	expected := New[int]()
	expected.Add(1)
	expected.Add(2)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestDiff_PartialOverlap_ShouldReturnNonOverlappingElements tests the Diff method with partial overlap.
func TestDiff_PartialOverlap_ShouldReturnNonOverlappingElements(t *testing.T) {
	set1 := New[int]()
	set1.Add(1)
	set1.Add(2)
	set1.Add(3)
	set2 := New[int]()
	set2.Add(2)
	set2.Add(3)
	set2.Add(4)
	result := set1.Diff(set2)
	expected := New[int]()
	expected.Add(1)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

func TestSet_Equal(t *testing.T) {
	tests := []struct {
		name     string
		set1     *Set[int]
		set2     *Set[int]
		expected bool
	}{
		{"DifferentSizes", New(1, 2, 3), New(1, 2), false},
		{"SameSizeDifferentElements", New(1, 2, 3), New(4, 5, 6), false},
		{"SameSizeSameElements", New(1, 2, 3), New(3, 2, 1), true},
		{"BothEmpty", New[int](), New[int](), true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.set1.Equal(test.set2)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestJoin_EmptySet_ReturnsEmptyString(t *testing.T) {
	s := New[string]()
	result := s.Join(",")
	assert.Equal(t, "", result)
}

func TestJoin_SingleElement_ReturnsElementWithoutSeparator(t *testing.T) {
	s := New[string]()
	s.Add("hello")
	result := s.Join(",")
	assert.Equal(t, "hello", result)
}

func TestJoin_MultipleElements_ReturnsJoinedString(t *testing.T) {
	s := New[string]()
	s.Add("hello")
	s.Add("world")
	result := s.Join(",")
	assert.Equal(t, "hello,world", result)
}

func TestJoin_EmptySeparator_ReturnsConcatenatedString(t *testing.T) {
	s := New[string]()
	s.Add("hello")
	s.Add("world")
	result := s.Join("")
	assert.Equal(t, "helloworld", result)
}

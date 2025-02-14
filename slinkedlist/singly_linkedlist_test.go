package slinkedlist

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

//go test slinkedlist/* -v

// TestAppend 测试 Append 方法的各种情况
func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		list     *SinglyLinkedList[int]
		val      []int
		expected *SinglyLinkedList[int]
	}{
		{
			name:     "empty input",
			list:     &SinglyLinkedList[int]{},
			val:      []int{},
			expected: &SinglyLinkedList[int]{},
		},
		{
			name:     "single value",
			list:     &SinglyLinkedList[int]{},
			val:      []int{1},
			expected: &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: nil}},
		},
		{
			name:     "multiple values",
			list:     &SinglyLinkedList[int]{},
			val:      []int{1, 2, 3},
			expected: &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: &SinglyNode[int]{Val: 2, Next: &SinglyNode[int]{Val: 3, Next: nil}}}},
		},
		{
			name:     "non-empty list",
			list:     &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: nil}},
			val:      []int{2, 3},
			expected: &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: &SinglyNode[int]{Val: 2, Next: &SinglyNode[int]{Val: 3, Next: nil}}}},
		},
		{
			name:     "multi-nodes list",
			list:     &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: &SinglyNode[int]{Val: 2}}},
			val:      []int{3, 4},
			expected: &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: &SinglyNode[int]{Val: 2, Next: &SinglyNode[int]{Val: 3, Next: &SinglyNode[int]{Val: 4, Next: nil}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.list.Append(tt.val...)
			if !reflect.DeepEqual(tt.list, tt.expected) {
				t.Errorf("Append() = %v, want %v", tt.list, tt.expected)
			}
		})
	}
}

// TestInsertAt_NegativeIndex_ShouldReturnError tests the InsertAt method with a negative index.
func TestInsertAt_NegativeIndex_ShouldReturnError(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	err := list.InsertAt(-1, 10)
	assert.Error(t, err, "index must be zero or a positive number")
}

// TestInsertAt_ZeroIndex_ShouldInsertAtHead tests the InsertAt method with index zero.
func TestInsertAt_ZeroIndex_ShouldInsertAtHead(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.InsertAt(0, 10)
	list.InsertAt(0, 20)
	assert.Equal(t, 20, list.Head.Val)
}

// TestInsertAt_ValidIndex_ShouldInsertAtCorrectPosition tests the InsertAt method with a valid index.
func TestInsertAt_ValidIndex_ShouldInsertAtCorrectPosition(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.InsertAt(0, 10)
	list.InsertAt(1, 20)
	list.InsertAt(1, 30)
	assert.Equal(t, 10, list.Head.Val)
	assert.Equal(t, 30, list.Head.Next.Val)
	assert.Equal(t, 20, list.Head.Next.Next.Val)
}

// TestInsertAt_OutOfRangeIndex_ShouldReturnError tests the InsertAt method with an out-of-range index.
func TestInsertAt_OutOfRangeIndex_ShouldReturnError(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.InsertAt(0, 10)
	err := list.InsertAt(2, 20)
	assert.Error(t, err, "index out of range")
}

// TestRemove_EmptyList_NoError
func TestRemove_EmptyList_NoError(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Remove(0)
	if list.Head != nil {
		t.Errorf("Expected list to remain empty after Remove on an empty list")
	}
}

// TestRemove_NegativeIndex_NoRemoval
func TestRemove_NegativeIndex_NoRemoval(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Remove(-1)
	if list.Head.Val != 1 {
		t.Errorf("Expected no removal for negative index")
	}
}

// TestRemove_RemoveHeadNode_HeadUpdated
func TestRemove_RemoveHeadNode_HeadUpdated(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Remove(0)
	if list.Head.Val != 2 {
		t.Errorf("Expected head to be updated to the next node after removing head")
	}
}

// TestRemove_RemoveMiddleNode_NodeRemoved
func TestRemove_RemoveMiddleNode_NodeRemoved(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(1)
	if list.Head.Next.Val != 3 {
		t.Errorf("Expected middle node to be removed")
	}
}

// TestRemove_RemoveTailNode_TailUpdated
func TestRemove_RemoveTailNode_TailUpdated(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Remove(1)
	if list.Head.Next != nil {
		t.Errorf("Expected tail to be updated after removing tail node")
	}
}

// TestRemove_OutOfRangeIndex_NoRemoval
func TestRemove_OutOfRangeIndex_NoRemoval(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Remove(10)
	if list.Head.Val != 1 {
		t.Errorf("Expected no removal for out of range index")
	}
}

// TestIndexOf_EmptyList_ReturnsMinusOne tests the IndexOf method on an empty list.
func TestIndexOf_EmptyList_ReturnsMinusOne(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	assert.Equal(t, -1, list.IndexOf(1))
}

// TestIndexOf_ValueExists_ReturnsCorrectIndex tests the IndexOf method when the value exists in the list.
func TestIndexOf_ValueExists_ReturnsCorrectIndex(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, 1, list.IndexOf(2))
}

// TestIndexOf_ValueDoesNotExist_ReturnsMinusOne tests the IndexOf method when the value does not exist in the list.
func TestIndexOf_ValueDoesNotExist_ReturnsMinusOne(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, -1, list.IndexOf(4))
}

// TestIndexOf_ValueAtStart_ReturnsZero tests the IndexOf method when the value is at the start of the list.
func TestIndexOf_ValueAtStart_ReturnsZero(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, 0, list.IndexOf(1))
}

// TestIndexOf_ValueAtEnd_ReturnsLastIndex tests the IndexOf method when the value is at the end of the list.
func TestIndexOf_ValueAtEnd_ReturnsLastIndex(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, 2, list.IndexOf(3))
}

// TestIndexOf_DuplicateValues_ReturnsFirstIndex tests the IndexOf method with duplicate values.
func TestIndexOf_DuplicateValues_ReturnsFirstIndex(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, 1, list.IndexOf(2))
}

// TestFind_EmptyList_ReturnsNil tests the Find method on an empty list.
func TestFind_EmptyList_ReturnsNil(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	node := list.Find(0)
	assert.Nil(t, node, "Expected nil for empty list")
}

// TestFind_ValidIndex_ReturnsNode tests the Find method with a valid index.
func TestFind_ValidIndex_ReturnsNode(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	node := list.Find(1)
	assert.NotNil(t, node, "Expected non-nil node")
	assert.Equal(t, 20, node.Val, "Expected value at index 1 to be 20")
}

// TestFind_IndexOutOfRange_ReturnsNil tests the Find method with an out-of-range index.
func TestFind_IndexOutOfRange_ReturnsNil(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	node := list.Find(2)
	assert.Nil(t, node, "Expected nil for out-of-range index")
}

// TestFind_NegativeIndex_ReturnsNil tests the Find method with a negative index.
func TestFind_NegativeIndex_ReturnsNil(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	node := list.Find(-1)
	assert.Nil(t, node, "Expected nil for negative index")
}

// TestFind_ZeroIndex_ReturnsHead tests the Find method with index zero.
func TestFind_ZeroIndex_ReturnsHead(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	node := list.Find(0)
	assert.NotNil(t, node, "Expected non-nil node")
	assert.Equal(t, 10, node.Val, "Expected value at index 0 to be 10")
}

// TestUpdate_ValidIndex_NodeUpdated tests the Update method when the index is valid.
func TestUpdate_ValidIndex_NodeUpdated(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Update(1, 25) // Update the second node's value to 25

	expected := []int{10, 25, 30}
	actual := list.ToSlice()

	if !assert.Equal(t, expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// TestUpdate_InvalidIndex_NoUpdate tests the Update method when the index is invalid.
func TestUpdate_InvalidIndex_NoUpdate(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	list.Update(5, 100) // Attempt to update a non-existent node

	expected := []int{10, 20, 30}
	actual := list.ToSlice()

	if !assert.Equal(t, expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// TestWalk_EmptyList_NoFunctionCall tests the Walk method on an empty list.
func TestWalk_EmptyList_NoFunctionCall(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	var called bool
	list.Walk(func(int) { called = true })
	if called {
		t.Errorf("Expected no function call on an empty list")
	}
}

// TestWalk_SingleNodeList_FunctionCalledOnce tests the Walk method on a single node list.
func TestWalk_SingleNodeList_FunctionCalledOnce(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	var count int
	list.Walk(func(int) { count++ })
	if count != 1 {
		t.Errorf("Expected function to be called once, got %d calls", count)
	}
}

// TestWalk_MultipleNodesList_FunctionCalledForEachNode tests the Walk method on a multiple nodes list.
func TestWalk_MultipleNodesList_FunctionCalledForEachNode(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	var count int
	list.Walk(func(int) { count++ })
	if count != 3 {
		t.Errorf("Expected function to be called three times, got %d calls", count)
	}
}

// TestWalk_ModifyNodeValues_DuringWalk tests the Walk method with modification of node values.
func TestWalk_ModifyNodeValues_DuringWalk(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Walk(func(val int) {
		if val == 2 {
			list.Head.Next.Val = 10 // Modify the second node's value
		}
	})
	var result []int
	list.Walk(func(val int) {
		result = append(result, val)
	})
	expected := []int{1, 10, 3}
	if !assert.Equal(t, result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}

// TestReverse_EmptyList_ShouldRemainUnchanged tests that reversing an empty list does not change it.
func TestReverse_EmptyList_ShouldRemainUnchanged(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Reverse()
	if list.Head != nil {
		t.Errorf("Expected list to remain empty after reverse, got: %v", list)
	}
}

// TestReverse_SingleNode_ShouldRemainUnchanged tests that reversing a single-node list does not change it.
func TestReverse_SingleNode_ShouldRemainUnchanged(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Reverse()
	if list.Head.Val != 1 || list.Head.Next != nil {
		t.Errorf("Expected single-node list to remain unchanged after reverse, got: %v", list)
	}
}

// TestReverse_MultipleNodes_ShouldReverseCorrectly tests that reversing a multi-node list reverses it correctly.
func TestReverse_MultipleNodes_ShouldReverseCorrectly(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()

	expectedValues := []int{3, 2, 1}
	actualValues := list.ToSlice()
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Expected list to be reversed to %v, got: %v", expectedValues, actualValues)
	}
}

// TestReverse_ReversingTwice_ShouldRestoreOriginalOrder tests that reversing a list twice restores the original order.
func TestReverse_ReversingTwice_ShouldRestoreOriginalOrder(t *testing.T) {
	list := NewSinglyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Reverse()
	list.Reverse()

	expectedValues := []int{1, 2, 3}
	actualValues := list.ToSlice()
	if !reflect.DeepEqual(actualValues, expectedValues) {
		t.Errorf("Expected list to restore original order after two reversals, got: %v", actualValues)
	}
}

// TestMerge tests the Merge method of SinglyLinkedList.
func TestMerge(t *testing.T) {
	tests := []struct {
		list1    *SinglyLinkedList[int]
		list2    *SinglyLinkedList[int]
		expected []int
	}{
		{
			list1:    &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1, Next: &SinglyNode[int]{Val: 2}}},
			list2:    &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 3, Next: &SinglyNode[int]{Val: 4}}},
			expected: []int{1, 2, 3, 4},
		},
		{
			list1:    &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1}},
			list2:    &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 2}},
			expected: []int{1, 2},
		},
		{
			list1:    &SinglyLinkedList[int]{Head: nil},
			list2:    &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1}},
			expected: []int{1},
		},
		{
			list1:    &SinglyLinkedList[int]{Head: &SinglyNode[int]{Val: 1}},
			list2:    &SinglyLinkedList[int]{Head: nil},
			expected: []int{1},
		},
		{
			list1:    &SinglyLinkedList[int]{Head: nil},
			list2:    &SinglyLinkedList[int]{Head: nil},
			expected: []int{},
		},
	}

	for i, test := range tests {
		test.list1.Merge(test.list2)
		var actual = []int{}
		current := test.list1.Head
		for current != nil {
			actual = append(actual, current.Val)
			current = current.Next
		}
		assert.Equal(t, test.expected, actual, "Merge method should merge two lists correctly, No.%d", i)
	}
}

func TestSinglyLinkedList_Length(t *testing.T) {
	tests := []struct {
		name     string
		listFunc func() *SinglyLinkedList[int]
		want     int
	}{
		{
			name: "empty list",
			listFunc: func() *SinglyLinkedList[int] {
				return NewSinglyLinkedList[int]()
			},
			want: 0,
		},
		{
			name: "single node",
			listFunc: func() *SinglyLinkedList[int] {
				list := NewSinglyLinkedList[int]()
				list.Append(1)
				return list
			},
			want: 1,
		},
		{
			name: "multiple nodes",
			listFunc: func() *SinglyLinkedList[int] {
				list := NewSinglyLinkedList[int]()
				list.Append(1)
				list.Append(2)
				list.Append(3)
				return list
			},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := tt.listFunc()
			got := list.Length()
			assert.Equal(t, tt.want, got, "Length() = %v, want %v", got, tt.want)
		})
	}
}

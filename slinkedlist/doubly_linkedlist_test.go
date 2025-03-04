package slinkedlist

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//go test -v -run=TestAppend slinkedlist/doubly_linkedlist*

// TestAppend_EmptyValues_NoChange tests it appending an empty slice to the list.
func TestAppend_EmptyValues_NoChange(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append()
	if list.head != nil || list.tail != nil {
		t.Errorf("Expected list to remain empty, but it was modified")
	}
}

// TestAppend_SingleValue_AddsToEmptyList tests appending a single value to an empty list.
func TestAppend_SingleValue_AddsToEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	if list.head == nil || list.tail == nil {
		t.Errorf("Expected list to have a head and tail, but it was empty")
	}
	if list.head.val != 1 || list.tail.val != 1 {
		t.Errorf("Expected value 1 to be the only element in the list")
	}
}

// TestAppend_MultipleValues_AddsToEmptyList tests appending multiple values to an empty list.
func TestAppend_MultipleValues_AddsToEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1, 2, 3)
	if list.head == nil || list.tail == nil {
		t.Errorf("Expected list to have a head and tail, but it was empty")
	}
	if list.head.val != 1 || list.tail.val != 3 {
		t.Errorf("Expected list to start with 1 and end with 3")
	}
}

// TestAppend_SingleValue_AddsToNonEmptyList tests appending a single value to a non-empty list.
func TestAppend_SingleValue_AddsToNonEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	if list.head == nil || list.tail == nil {
		t.Errorf("Expected list to have a head and tail, but it was empty")
	}
	if list.head.val != 1 || list.tail.val != 2 {
		t.Errorf("Expected list to start with 1 and end with 2")
	}
}

// TestAppend_MultipleValues_AddsToNonEmptyList tests appending multiple values to a non-empty list.
func TestAppend_MultipleValues_AddsToNonEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2, 3)
	if list.head == nil || list.tail == nil {
		t.Errorf("Expected list to have a head and tail, but it was empty")
	}
	if list.head.val != 1 || list.tail.val != 3 {
		t.Errorf("Expected list to start with 1 and end with 3")
	}
}

// TestInsertBefore_NegativeIndex_ReturnsError tests the InsertBefore method when the index is negative.
func TestInsertBefore_NegativeIndex_ReturnsError(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	err := list.InsertBefore(-1, 10)
	assert.Error(t, err, "index must be zero or a positive number")
}

// TestInsertBefore_EmptyList_ReturnsError tests the InsertBefore method when the list is empty.
func TestInsertBefore_EmptyList_ReturnsError(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	err := list.InsertBefore(0, 10)
	assert.Error(t, err, "index out of range")
}

// TestInsertBefore_IndexZero_InsertsAtHead tests the InsertBefore method when inserting at index zero.
func TestInsertBefore_IndexZero_InsertsAtHead(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	err := list.InsertBefore(0, 10)
	assert.NoError(t, err)
	assert.Equal(t, 10, list.head.val)
	assert.Equal(t, 1, list.head.next.val)
	assert.Equal(t, 1, list.tail.val)
}

// TestInsertBefore_ValidIndex_InsertsNode tests the InsertBefore method when inserting at a valid index.
func TestInsertBefore_ValidIndex_InsertsNode(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.InsertBefore(2, 10)
	assert.NoError(t, err)
	assert.Equal(t, "(int)(1) <-> (int)(2) <-> (int)(10) <-> (int)(3)", list.String())
}

// TestInsertBefore_OutOfRangeIndex_ReturnsError tests the InsertBefore method when inserting at an out-
func TestInsertBefore_OutOfRangeIndex_ReturnsError(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	err := list.InsertBefore(2, 10)
	assert.Error(t, err, "index out of range")
}

// TestInsertBefore_InsertAtTail_ReturnsError tests the InsertBefore method when inserting at the tail of the list.
func TestInsertBefore_InsertAtTail_ReturnsError(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	err := list.InsertBefore(1, 10)
	assert.Error(t, err, "index out of range")
}

// TestInsertAfterOnEmptyList tests the InsertAfter method when the list is empty.
func TestInsertAfterOnEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	err := list.InsertAfter(0, 1)
	if err == nil {
		t.Errorf("Expected an error when inserting into an empty list, but got nil")
	}
}

// TestInsertAfterSingleNode tests the InsertAfter method when inserting a single node.
func TestInsertAfterSingleNode(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	err := list.InsertAfter(0, 2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if list.head.val != 1 || list.head.next.val != 2 {
		t.Errorf("Expected list to be [1, 2], but got head: %v, next: %v", list.head.val, list.head.next.val)
	}
}

// TestInsertAfterMultiNode tests the InsertAfter method when inserting multiple nodes.
func TestInsertAfterMultiNode(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.InsertAfter(1, 4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// 验证链表状态
	if list.head.next.val != 2 || list.head.next.next.val != 4 {
		t.Errorf("Expected list to be [1, 2, 4, 3], but got next: %v, next.next: %v", list.head.next.val, list.head.next.next.val)
	}
}

// TestInsertAfterEnd tests the InsertAfter method when inserting at the end of the list.
func TestInsertAfterEnd(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAfter(1, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// 验证链表状态
	if list.tail.val != 3 {
		t.Errorf("Expected list to be [1, 2, 3], but got tail: %v", list.tail.val)
	}
}

// TestInsertAfterInvalidIndex tests the InsertAfter method when inserting at an invalid index.
func TestInsertAfterInvalidIndex(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAfter(2, 3)
	if err == nil {
		t.Errorf("Expected an error when inserting at an invalid index, but got nil")
	}
}

// TestInsertAfterNegativeIndex tests the InsertAfter method when inserting at a negative index.
func TestInsertAfterNegativeIndex(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAfter(-1, 3)
	if err == nil {
		t.Errorf("Expected an error when inserting at a negative index, but got nil")
	}
}

// TestInsertAfterMultiple tests the InsertAfter method when inserting multiple nodes.
func TestInsertAfterMultiple(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	err := list.InsertAfter(0, 4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = list.InsertAfter(2, 5)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if list.head.next.val != 4 || list.head.next.next.val != 2 {
		t.Errorf("Expected list to be [1, 4, 2, 5, 3], but got next: %v, next.next: %v", list.head.next.val, list.head.next.next.val)
	}
}

// TestInsertAfterHead tests the InsertAfter method when inserting at the head.
func TestInsertAfterHead(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	err := list.InsertAfter(0, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if list.head.val != 1 || list.head.next.val != 3 {
		t.Errorf("Expected list to be [1, 3, 2], but got head: %v, next: %v", list.head.val, list.head.next.val)
	}
}

// TestRemoveEmptyList tests the Remove method on an empty list.
func TestRemoveEmptyList(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Remove(0)
	if list.head != nil {
		t.Errorf("Expected head to be nil, got %v", list.head)
	}
}

// TestRemoveNegativeIndex tests the Remove method with a negative index.
func TestRemoveNegativeIndex(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Remove(-1)
	if list.head.val != 1 {
		t.Errorf("Expected head value to be 1, got %v", list.head.val)
	}
}

// TestRemoveHead tests the Remove method to remove the head node.
func TestRemoveHead(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Remove(0)
	if list.head == nil || list.head.val != 2 {
		t.Errorf("Expected head value to be 2, got %v", list.head)
	}
	if list.head.prev != nil {
		t.Errorf("Expected head.prev to be nil, got %v", list.head.prev)
	}
	if list.tail == nil || list.tail.val != 2 {
		t.Errorf("Expected tail val to be 2, got %v", list.tail)
	}
}

// TestRemoveMiddle tests the Remove method to remove a middle node.
func TestRemoveMiddle(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Remove(1)
	if list.head == nil || list.head.val != 1 {
		t.Errorf("Expected head value to be 1, got %v", list.head)
	}
	if list.head.next == nil || list.head.next.val != 3 {
		t.Errorf("Expected second node value to be 3, got %v", list.head.next)
	}
	if list.head.next.prev == nil || list.head.next.prev.val != 1 {
		t.Errorf("Expected second node prev value to be 1, got %v", list.head.next.prev)
	}
	if list.tail == nil || list.tail.val != 3 {
		t.Errorf("Expected tail value to be 3, got %v", list.tail)
	}
}

// TestRemoveTail tests the Remove method to remove the tail node.
func TestRemoveTail(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Remove(1)
	if list.head.next != nil {
		t.Errorf("Expected tail to be nil, got %v", list.head.next)
	}
	if list.tail == nil || list.tail.val != 1 {
		t.Errorf("Expected tail value to be 1, got %v", list.tail)
	}
}

// TestRemoveIndexOutOfRange tests the Remove method with an index out of range.
func TestRemoveIndexOutOfRange(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Remove(1)
	if list.head == nil || list.head.val != 1 {
		t.Errorf("Expected head value to be 1, got %v", list.head)
	}
}

// TestIndexOf_EmptyList_ReturnsMinusOne_Doubly tests the IndexOf method on an empty list.
func TestIndexOf_EmptyList_ReturnsMinusOne_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	assert.Equal(t, -1, list.IndexOf(10))
}

// TestIndexOf_ValueExists_ReturnsCorrectIndex_Doubly tests the IndexOf method with a value that exists in the list.
func TestIndexOf_ValueExists_ReturnsCorrectIndex_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	assert.Equal(t, 1, list.IndexOf(20))
}

// TestIndexOf_ValueDoesNotExist_ReturnsMinusOne_Doubly tests the IndexOf method with a value that does not exist in the list.
func TestIndexOf_ValueDoesNotExist_ReturnsMinusOne_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)
	assert.Equal(t, -1, list.IndexOf(40))
}

// TestIndexOf_ValueExistsMultipleTimes_ReturnsFirstIndex_Doubly tests the IndexOf method with a value that exists multiple times in the list.
func TestIndexOf_ValueExistsMultipleTimes_ReturnsFirstIndex_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(20)
	list.Append(30)
	assert.Equal(t, 1, list.IndexOf(20))
}

// TestFind_EmptyList_ReturnsNil_Doubly tests the Find method on an empty list.
func TestFind_EmptyList_ReturnsNil_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	node := list.Find(0)
	assert.Nil(t, node, "Expected nil for empty list")
}

// TestFind_ValidIndex_ReturnsNode_Doubly tests the Find method with a valid index.
func TestFind_ValidIndex_ReturnsNode_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	node := list.Find(1)
	assert.NotNil(t, node, "Expected non-nil node")
	assert.Equal(t, 20, node.val, "Expected value at index 1 to be 20")
}

// TestFind_IndexOutOfRange_ReturnsNil_Doubly tests the Find method with an out-of-range index.
func TestFind_IndexOutOfRange_ReturnsNil_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	node := list.Find(2)
	assert.Nil(t, node, "Expected nil for out-of-range index")
}

// TestFind_NegativeIndex_ReturnsNil_Doubly tests the Find method with a negative index.
func TestFind_NegativeIndex_ReturnsNil_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	node := list.Find(-1)
	assert.Nil(t, node, "Expected nil for negative index")
}

// TestFind_ZeroIndex_ReturnsHead_Doubly tests the Find method with index zero.
func TestFind_ZeroIndex_ReturnsHead_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	node := list.Find(0)
	assert.NotNil(t, node, "Expected non-nil node")
	assert.Equal(t, 10, node.val, "Expected value at index 0 to be 10")
}

// TestUpdate_NodeExists_ValueUpdated
func TestUpdate_NodeExists_ValueUpdated_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)
	list.Append(30)

	err := list.Update(1, 25)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if list.head.next.val != 25 {
		t.Errorf("Expected value at index 1 to be 25, got %v", list.head.next.val)
	}
}

// TestUpdate_NodeDoesNotExist_ReturnsError
func TestUpdate_NodeDoesNotExist_ReturnsError_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(10)
	list.Append(20)

	err := list.Update(2, 30)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	} else {
		expectedError := fmt.Errorf("index out of range")
		if err.Error() != expectedError.Error() {
			t.Errorf("Expected error %v, got %v", expectedError, err)
		}
	}
}

// TestWalk_ForwardTraversal tests the Walk method for forward traversal.
func TestWalk_ForwardTraversal(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	expected := []int{1, 2, 3}
	actual := []int{}

	list.Walk(func(val int) {
		actual = append(actual, val)
	})

	if !assert.Equal(t, expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// TestWalk_BackwardTraversal tests the Walk method for backward traversal.
func TestWalk_BackwardTraversal(t *testing.T) {
	list := &DoublyLinkedList[int]{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	expected := []int{3, 2, 1}
	actual := []int{}

	list.Walk(func(val int) {
		actual = append(actual, val)
	}, true)

	if !assert.Equal(t, expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

// TestAppendSingleNode tests the Append method for a single node.
func TestAppendSingleNode(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	if list.head == nil || list.tail == nil {
		t.Errorf("Expected head and tail to be set after appending a node, got head: %v, tail: %v", list.head, list.tail)
	}
	if list.head.val != 1 || list.tail.val != 1 {
		t.Errorf("Expected head and tail values to be 1, got head.val: %v, tail.val: %v", list.head.val, list.tail.val)
	}
	if list.head.prev != nil || list.head.next != nil {
		t.Errorf("Expected head.prev and head.next to be nil, got head.prev: %v, head.next: %v", list.head.prev, list.head.next)
	}
}

// TestReverseEmptyList tests the Reverse method for an empty list.
func TestReverseEmptyList(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Reverse()
	if list.head != nil || list.tail != nil {
		t.Errorf("Expected head and tail to be nil after reversing an empty list, got head: %v, tail: %v", list.head, list.tail)
	}
}

// TestReverseSingleNode tests the Reverse method for a single node list.
func TestReverseSingleNode(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Reverse()
	if list.head.val != 1 || list.tail.val != 1 {
		t.Errorf("Expected head.val and tail.val to be 1 after reversing a single node list, got head.val: %v, tail.val: %v", list.head.val, list.tail.val)
	}
	if list.head.prev != nil || list.head.next != nil {
		t.Errorf("Expected head.prev and head.next to be nil after reversing a single node list, got head.prev: %v, head.next: %v", list.head.prev, list.head.next)
	}
}

// TestReverseMultiNode tests the Reverse method for a multi-node list.
func TestReverseMultiNode(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, []int{1, 2, 3}, list.ToSlice())
	list.Reverse()
	if list.head.val != 3 || list.tail.val != 1 {
		t.Errorf("Expected head.val to be 3 and tail.val to be 1 after reversing, got head.val: %v, tail.val: %v", list.head.val, list.tail.val)
	}
	if list.head.next.val != 2 || list.tail.prev.val != 2 {
		t.Errorf("Expected head.next.val to be 2 and tail.prev.val to be 2 after reversing, got head.next.val: %v, tail.prev.val: %v", list.head.next.val, list.tail.prev.val)
	}
}

// TestMerge_BothListsEmpty_NoChange tests the Merge method when both lists are empty.
func TestMerge_BothListsEmpty_NoChange(t *testing.T) {
	list1 := &DoublyLinkedList[int]{}
	list2 := &DoublyLinkedList[int]{}
	list1.Merge(list2)
	if list1.head != nil || list1.tail != nil {
		t.Errorf("Expected both head and tail to be nil, got head: %v, tail: %v", list1.head, list1.tail)
	}
}

// TestMerge_List1Empty_List2Merged tests the Merge method when list1 is empty.
func TestMerge_List1Empty_List2Merged(t *testing.T) {
	list1 := &DoublyLinkedList[int]{}
	list2 := &DoublyLinkedList[int]{}
	list2.Append(1)
	list2.Append(2)
	list1.Merge(list2)
	if list1.head != list2.head || list1.tail != list2.tail {
		t.Errorf("Expected list1 to be merged with list2, got %v", list1)
	}
}

// TestMerge_List2Empty_List1Unchanged tests the Merge method when list2 is empty.
func TestMerge_List2Empty_List1Unchanged(t *testing.T) {
	list1 := &DoublyLinkedList[int]{}
	list1.Append(1)
	list1.Append(2)
	list2 := &DoublyLinkedList[int]{}
	list1.Merge(list2)
	if list1.String() != "(int)(1) <-> (int)(2)" {
		t.Errorf("Expected list1 to remain unchanged, got %v", list1)
	}
}

// TestMerge_BothListsNonEmpty_MergedCorrectly tests the Merge method when both lists are non-empty.
func TestMerge_BothListsNonEmpty_MergedCorrectly(t *testing.T) {
	list1 := &DoublyLinkedList[int]{}
	list1.Append(1)
	list1.Append(2)
	list2 := &DoublyLinkedList[int]{}
	list2.Append(3)
	list2.Append(4)
	list1.Merge(list2)
	if list1.String() != "(int)(1) <-> (int)(2) <-> (int)(3) <-> (int)(4)" {
		t.Errorf("Expected lists to be merged correctly, got %v", list1)
	}
}

// TestString_EmptyList_ReturnsEmptyString tests the String method on an empty list.
func TestString_EmptyList_ReturnsEmptyString_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	assert.Equal(t, "", list.String())
}

// TestString_SingleNode_ReturnsNodeValue tests the String method on a list with a single node.
func TestString_SingleNode_ReturnsNodeValue_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	assert.Equal(t, "(int)(1)", list.String())
}

// TestString_MultipleNodes_ReturnsCorrectString tests the String method on a list with multiple nodes.
func TestString_MultipleNodes_ReturnsCorrectString_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	assert.Equal(t, "(int)(1) <-> (int)(2) <-> (int)(3)", list.String())
}

// TestString_ListWithNilValues_HandlesNilValues tests the String method on a list with nil values.
func TestString_ListWithNilValues_HandlesNilValues(t *testing.T) {
	list := NewDoublyLinkedList[*int]()
	var a, b int = 1, 2
	list.Append(&a)
	list.Append(nil)
	list.Append(&b)
	assert.Equal(t, "(*int)(1) <-> (*int)(nil) <-> (*int)(2)", list.String())
}

// TestString_MultiLevelPointers_ReturnsCorrectString tests the String method on a list with multi-level pointers.
func TestString_MultiLevelPointers_ReturnsCorrectString_Doubly(t *testing.T) {
	// 创建多层指针
	var _1 = 1
	var a = new(*int)
	var b = new(*int)
	*b = &_1

	list := NewDoublyLinkedList[**int]()
	list.Append(a)
	list.Append(b)

	expected := fmt.Sprintf("(**int)(nil) <-> (**int)(1)")

	assert.Equal(t, expected, list.String())
}

// TestCountOf_EmptyList_ShouldReturnZero tests the CountOf method on an empty list.
func TestCountOf_EmptyList_ShouldReturnZero_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	assert.Equal(t, 0, list.CountOf(1))
}

// TestCountOf_SingleNodeMatchingValue_ShouldReturnOne tests the CountOf method with a single node matching the value.
func TestCountOf_SingleNodeMatchingValue_ShouldReturnOne_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	assert.Equal(t, 1, list.CountOf(1))
}

// TestCountOf_SingleNodeNonMatchingValue_ShouldReturnZero tests the CountOf method with a single node not matching the value.
func TestCountOf_SingleNodeNonMatchingValue_ShouldReturnZero_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(2)
	assert.Equal(t, 0, list.CountOf(1))
}

// TestCountOf_MultipleNodesAllMatching_ShouldReturnCount tests the CountOf method with all nodes matching the value.
func TestCountOf_MultipleNodesAllMatching_ShouldReturnCount_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(1)
	list.Append(1)
	assert.Equal(t, 3, list.CountOf(1))
}

// TestCountOf_MultipleNodesNoneMatching_ShouldReturnZero tests the CountOf method with no nodes matching the value.
func TestCountOf_MultipleNodesNoneMatching_ShouldReturnZero_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(2)
	list.Append(3)
	list.Append(4)
	assert.Equal(t, 0, list.CountOf(1))
}

// TestCountOf_MultipleNodesSomeMatching_ShouldReturnCount tests the CountOf method with some nodes matching the value.
func TestCountOf_MultipleNodesSomeMatching_ShouldReturnCount_Doubly(t *testing.T) {
	list := NewDoublyLinkedList[int]()
	list.Append(1)
	list.Append(2)
	list.Append(1)
	list.Append(3)
	list.Append(1)
	assert.Equal(t, 3, list.CountOf(1))
}

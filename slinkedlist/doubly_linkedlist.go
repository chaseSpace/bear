package slinkedlist

import (
	"fmt"
	"github.com/chaseSpace/bear/butil"
)

type DoublyNode[T comparable] struct {
	val  T
	prev *DoublyNode[T]
	next *DoublyNode[T]
}

type DoublyLinkedList[T comparable] struct {
	head *DoublyNode[T]
	tail *DoublyNode[T]
}

// NewDoublyLinkedList creates a new doubly linked list.
func NewDoublyLinkedList[T comparable]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

// Append adds one or more values to the end of the linked list.
func (list *DoublyLinkedList[T]) Append(val ...T) {
	if len(val) == 0 {
		return
	}
	curr := &DoublyNode[T]{val: val[0], prev: nil, next: nil}
	first := curr
	for i := 1; i < len(val); i++ {
		curr.next = &DoublyNode[T]{val: val[i], prev: curr, next: nil}
		curr = curr.next
	}
	if list.head == nil { // empty list
		list.head = first
		list.tail = curr
		return
	}
	// replace curr with curr
	first.prev = list.tail
	list.tail.next = first
	list.tail = curr
	return
}

// InsertBefore inserts a new node with the specified value before the node at the specified index.
func (list *DoublyLinkedList[T]) InsertBefore(index int, val T) error {
	if index < 0 {
		return fmt.Errorf("index must be zero or a positive number")
	}
	if list.head == nil { // operation on the empty list is prohibited
		return fmt.Errorf("index out of range")
	}
	newNode := &DoublyNode[T]{val: val, next: nil}
	if index == 0 { // insert before head
		newNode.next = list.head
		list.head = newNode
		return nil
	}

	// find the node that is before the node points to index
	current := list.head
	i := 1 // start from the second node
	for ; i < index; i++ {
		current = current.next
		if current == nil { // index > length
			return fmt.Errorf("index out of range")
		}
	}
	if current.next == nil { // index = length(point to tail)
		return fmt.Errorf("index out of range")
	}
	// do connect current <-> newNode
	newNode.prev = current
	newNode.next = current.next
	current.next = newNode
	// because insertion into(replace) the tail is not allowed here, there is no need to change the tail node
	return nil
}

// InsertAfter inserts a new node with the specified value after the node at the specified index.
func (list *DoublyLinkedList[T]) InsertAfter(index int, val T) error {
	if index < 0 {
		return fmt.Errorf("index must be zero or a positive number")
	}
	if list.head == nil { // operation on the empty list is prohibited
		return fmt.Errorf("index out of range")
	}
	newNode := &DoublyNode[T]{val: val, next: nil}

	current := list.head
	// find the node points to index
	for i := 0; i < index; i++ {
		current = current.next
		if current == nil {
			return fmt.Errorf("index out of range")
		}
	}
	// insertion after the tail node is possible
	if current.next == nil {
		list.tail = newNode
	}
	// do connect current <-> newNode
	newNode.prev = current
	newNode.next = current.next
	current.next = newNode
	return nil
}

// Remove removes the node at the specified index.
func (list *DoublyLinkedList[T]) Remove(index int) {
	if list.head == nil || index < 0 {
		return
	}

	if index == 0 {
		list.head = list.head.next
		if list.head != nil {
			list.head.prev = nil
		} else {
			list.tail = nil
		}
		return
	}
	// find one this is before the node points to index
	current := list.head
	for i := 1; current != nil && i < index; i++ {
		current = current.next
	}

	// if current is nil or current.next is nil, the index is out of range
	if current == nil || current.next == nil {
		return
	}

	// remove the node at index
	current.next = current.next.next
	if current.next != nil {
		current.next.prev = current
	} else {
		list.tail = current
	}
}

// IndexOf returns the index of the first occurrence of the specified value in the linked list.
func (list *DoublyLinkedList[T]) IndexOf(val T) int {
	current := list.head
	for i := 0; current != nil; i++ {
		if current.val == val {
			return i
		}
		current = current.next
	}
	return -1
}

// Find returns the node at the specified index.
func (list *DoublyLinkedList[T]) Find(index int) *DoublyNode[T] {
	current := list.head
	for i := 0; current != nil; i++ {
		if i == index {
			return current
		}
		current = current.next
	}
	return nil
}

// Update updates the value of the node at the specified index.
func (list *DoublyLinkedList[T]) Update(index int, newVal T) error {
	node := list.Find(index)
	if node != nil {
		node.val = newVal
		return nil
	} else {
		return fmt.Errorf("index out of range")
	}
}

// Walk applies a function to each node in the linked list.
func (list *DoublyLinkedList[T]) Walk(f func(T), reverse ...bool) {
	if len(reverse) > 0 && reverse[0] {
		current := list.tail
		for current != nil {
			f(current.val)
			current = current.prev
		}
		return
	}
	current := list.head
	for current != nil {
		f(current.val)
		current = current.next
	}
}

// Reverse reverses the linked list.
func (list *DoublyLinkedList[T]) Reverse() {
	if list.head == nil || list.head.next == nil {
		return
	}
	var curr, prev, next *DoublyNode[T]
	for curr = list.head; ; {
		// remember prev/next
		prev = curr.prev
		next = curr.next

		// switch next and prev
		curr.prev = next
		curr.next = prev

		if next == nil {
			break
		}
		curr = next
	}
	list.tail = list.head
	list.head = curr
}

// Merge merges the current linked list with another linked list.
func (list *DoublyLinkedList[T]) Merge(other *DoublyLinkedList[T]) {
	if other == nil || other.head == nil {
		return
	}
	if list.head == nil {
		list.head = other.head
		list.tail = other.tail
		return
	}
	// connect 1.tail -> 2.head
	list.tail.next = other.head
	other.head.prev = list.tail

	list.tail = other.tail
}

// ToSlice converts the linked list to a slice.
func (list *DoublyLinkedList[T]) ToSlice() []T {
	var arr []T
	current := list.head
	for current != nil {
		arr = append(arr, current.val)
		current = current.next
	}
	return arr
}

// Length returns the length of the linked list.
func (list *DoublyLinkedList[T]) Length() int {
	length := 0
	current := list.head
	for current != nil {
		length++
		current = current.next
	}
	return length
}

// IsEmpty checks if the linked list is empty.
func (list *DoublyLinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

// String returns a string representation of the linked list.
// For each node, the string representation includes the type and value of the node.
func (list *DoublyLinkedList[T]) String() string {
	var result string
	current := list.head
	for current != nil {
		result += fmt.Sprintf("%s", butil.PrintReadableTypeValue(current.val))
		if current.next != nil {
			result += " <-> "
		}
		current = current.next
	}
	return result
}

// CountOf returns the count occurrences of a specific value in the linked list.
func (list *DoublyLinkedList[T]) CountOf(val T) int {
	var count int
	current := list.head
	for current != nil {
		if current.val == val {
			count++
		}
		current = current.next
	}
	return count
}

package slinkedlist

import "fmt"

type SinglyNode[T comparable] struct {
	Val  T
	Next *SinglyNode[T]
}

type SinglyLinkedList[T comparable] struct {
	Head *SinglyNode[T]
}

// NewSinglyLinkedList creates a new singly linked list.
func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

// Append adds one or more values to the end of the linked list.
func (list *SinglyLinkedList[T]) Append(val ...T) {
	if len(val) == 0 {
		return
	}
	newNode := &SinglyNode[T]{Val: val[0], Next: nil}
	first := newNode
	for i := 1; i < len(val); i++ {
		newNode.Next = &SinglyNode[T]{Val: val[i], Next: nil}
		newNode = newNode.Next
	}

	if list.Head == nil {
		list.Head = first
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = first
}

// InsertAt inserts a new node at the specified index.
func (list *SinglyLinkedList[T]) InsertAt(index int, val T) error {
	if index < 0 {
		return fmt.Errorf("index must be zero or a positive number")
	}
	newNode := &SinglyNode[T]{Val: val, Next: nil}
	if index == 0 {
		newNode.Next = list.Head
		list.Head = newNode
		return nil
	}
	current := list.Head
	for i := 1; current != nil && i < index; i++ {
		current = current.Next
	}
	if current == nil {
		return fmt.Errorf("index out of range")
	}
	newNode.Next = current.Next
	current.Next = newNode
	return nil
}

// Remove removes the node at the specified index.
func (list *SinglyLinkedList[T]) Remove(index int) {
	if list.Head == nil || index < 0 {
		return
	}

	if index == 0 {
		list.Head = list.Head.Next
		return
	}

	current := list.Head
	for i := 1; current != nil && i < index; i++ {
		current = current.Next
	}

	if current == nil || current.Next == nil {
		return
	}

	current.Next = current.Next.Next
}

// IndexOf returns the index of the first occurrence of the specified value in the linked list.
func (list *SinglyLinkedList[T]) IndexOf(val T) int {
	current := list.Head
	for i := 0; current != nil; i++ {
		if current.Val == val {
			return i
		}
		current = current.Next
	}
	return -1
}

// Find returns the node at the specified index.
func (list *SinglyLinkedList[T]) Find(index int) *SinglyNode[T] {
	current := list.Head
	for i := 0; current != nil; i++ {
		if i == index {
			return current
		}
		current = current.Next
	}
	return nil
}

// Update updates the value of the node at the specified index.
func (list *SinglyLinkedList[T]) Update(index int, newVal T) {
	node := list.Find(index)
	if node != nil {
		node.Val = newVal
	}
}

// Walk applies a function to each node in the linked list.
func (list *SinglyLinkedList[T]) Walk(f func(T)) {
	current := list.Head
	for current != nil {
		f(current.Val)
		current = current.Next
	}
}

// Reverse reverses the linked list.
func (list *SinglyLinkedList[T]) Reverse() {
	if list.Head == nil || list.Head.Next == nil {
		return
	}
	var prev *SinglyNode[T]
	var current = list.Head
	var next *SinglyNode[T]

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	list.Head = prev
}

// Merge merges the current linked list with another linked list.
func (list *SinglyLinkedList[T]) Merge(other *SinglyLinkedList[T]) {
	if other == nil {
		return
	}
	if list.Head == nil {
		list.Head = other.Head
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = other.Head
}

// ToSlice converts the linked list to a slice.
func (list *SinglyLinkedList[T]) ToSlice() []T {
	var arr []T
	current := list.Head
	for current != nil {
		arr = append(arr, current.Val)
		current = current.Next
	}
	return arr
}

// Length returns the length of the linked list.
func (list *SinglyLinkedList[T]) Length() int {
	length := 0
	current := list.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}

// IsEmpty checks if the linked list is empty.
func (list *SinglyLinkedList[T]) IsEmpty() bool {
	return list.Head == nil
}

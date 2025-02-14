package slinkedlist

import "fmt"

type SinglyNode[T comparable] struct {
	val  T
	next *SinglyNode[T]
}

type SinglyLinkedList[T comparable] struct {
	head, tail *SinglyNode[T]
}

// NewSinglyLinkedList creates a new singly linked list.
func NewSinglyLinkedList[T comparable]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (list *SinglyLinkedList[T]) setTail() {
	if list.tail == nil && list.head != nil {
		list.tail = list.head
	}
}

// Append adds one or more values to the end of the linked list.
func (list *SinglyLinkedList[T]) Append(val ...T) {
	if len(val) == 0 {
		return
	}
	newNode := &SinglyNode[T]{val: val[0], next: nil}
	first := newNode
	for i := 1; i < len(val); i++ {
		newNode.next = &SinglyNode[T]{val: val[i], next: nil}
		newNode = newNode.next
	}

	if list.head == nil { // empty list
		list.head = first
		list.tail = newNode
		return
	}
	list.tail.next = first
	list.tail = newNode
}

// InsertBefore inserts a new node with the specified value before the node at the specified index.
func (list *SinglyLinkedList[T]) InsertBefore(index int, val T) error {
	if index < 0 {
		return fmt.Errorf("index must be zero or a positive number")
	}
	if list.head == nil { // operation on the empty list is prohibited
		return fmt.Errorf("index out of range")
	}
	newNode := &SinglyNode[T]{val: val, next: nil}
	if index == 0 { // insert before head
		newNode.next = list.head
		list.head = newNode
		list.setTail()
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
	if current.next == nil { // index = length
		return fmt.Errorf("index out of range")
	}
	// do insert operation
	newNode.next = current.next
	current.next = newNode
	// because insertion into the tail is not allowed here, there is no need to change the tail node
	return nil
}

// InsertAfter inserts a new node with the specified value after the node at the specified index.
func (list *SinglyLinkedList[T]) InsertAfter(index int, val T) error {
	if index < 0 {
		return fmt.Errorf("index must be zero or a positive number")
	}
	if list.head == nil { // operation on the empty list is prohibited
		return fmt.Errorf("index out of range")
	}
	newNode := &SinglyNode[T]{val: val, next: nil}

	current := list.head
	// find the node points to index
	for i := 0; i < index; i++ {
		current = current.next
		if current == nil {
			return fmt.Errorf("index out of range")
		}
	}
	// do insert operation
	newNode.next = current.next
	current.next = newNode
	return nil
}

// Remove removes the node at the specified index.
func (list *SinglyLinkedList[T]) Remove(index int) {
	if list.head == nil || index < 0 {
		return
	}

	if index == 0 {
		list.head = list.head.next
		return
	}

	current := list.head
	for i := 1; current != nil && i < index; i++ {
		current = current.next
	}

	if current == nil || current.next == nil {
		return
	}

	current.next = current.next.next
}

// IndexOf returns the index of the first occurrence of the specified value in the linked list.
func (list *SinglyLinkedList[T]) IndexOf(val T) int {
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
func (list *SinglyLinkedList[T]) Find(index int) *SinglyNode[T] {
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
func (list *SinglyLinkedList[T]) Update(index int, newVal T) {
	node := list.Find(index)
	if node != nil {
		node.val = newVal
	}
}

// Walk applies a function to each node in the linked list.
func (list *SinglyLinkedList[T]) Walk(f func(T)) {
	current := list.head
	for current != nil {
		f(current.val)
		current = current.next
	}
}

// Reverse reverses the linked list.
func (list *SinglyLinkedList[T]) Reverse() {
	if list.head == nil || list.head.next == nil {
		return
	}
	var prev *SinglyNode[T]
	var current = list.head
	var next *SinglyNode[T]

	for current != nil {
		next = current.next
		current.next = prev
		prev = current
		current = next
	}

	list.head = prev
}

// Merge merges the current linked list with another linked list.
func (list *SinglyLinkedList[T]) Merge(other *SinglyLinkedList[T]) {
	if other == nil {
		return
	}
	if list.head == nil {
		list.head = other.head
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = other.head
}

// ToSlice converts the linked list to a slice.
func (list *SinglyLinkedList[T]) ToSlice() []T {
	var arr []T
	current := list.head
	for current != nil {
		arr = append(arr, current.val)
		current = current.next
	}
	return arr
}

// Length returns the length of the linked list.
func (list *SinglyLinkedList[T]) Length() int {
	length := 0
	current := list.head
	for current != nil {
		length++
		current = current.next
	}
	return length
}

// IsEmpty checks if the linked list is empty.
func (list *SinglyLinkedList[T]) IsEmpty() bool {
	return list.head == nil
}

package sslice

import (
	"fmt"
	"math/rand"
	"strings"
)

type Slice[T comparable] struct {
	data   []T
	sortFn func([]T, ...bool)
}

// New creates a new ptr of Slice type.
func New[T comparable](items ...T) *Slice[T] {
	if len(items) == 0 {
		items = []T{}
	}
	return &Slice[T]{data: items}
}

// Append appends data to the slice. It returns the Slice itself.
func (s *Slice[T]) Append(items ...T) *Slice[T] {
	s.data = append(s.data, items...)
	return s
}

func (s *Slice[T]) Clone() *Slice[T] {
	copied := make([]T, len(s.data))
	copy(copied, s.data)
	return &Slice[T]{data: copied}
}

// Filter filters the slice. It returns the Slice itself.
func (s *Slice[T]) Filter(f func(T) bool) *Slice[T] {
	var filtered []T
	for _, item := range s.data {
		if f(item) {
			filtered = append(filtered, item)
		}
	}
	s.data = filtered
	return s
}

// Map maps the slice. It returns the Slice itself.
func (s *Slice[T]) Map(f func(T) T) *Slice[T] {
	for i, item := range s.data {
		s.data[i] = f(item)
	}
	return s
}

// Unique removes duplicate items from the slice. It returns the Slice itself.
func (s *Slice[T]) Unique() *Slice[T] {
	var duplicates = make(map[T]struct{})
	var unique []T
	for _, item := range s.data {
		if _, ok := duplicates[item]; !ok {
			unique = append(unique, item)
			duplicates[item] = struct{}{}
		}
	}
	s.data = unique
	return s
}

// Reverse reverses the slice. It returns the Slice itself.
func (s *Slice[T]) Reverse() *Slice[T] {
	for i, j := 0, len(s.data)-1; i < j; i, j = i+1, j-1 {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	}
	return s
}

// Shuffle shuffles the slice. It returns the Slice itself.
func (s *Slice[T]) Shuffle() *Slice[T] {
	rand.Shuffle(len(s.data), func(i, j int) {
		s.data[i], s.data[j] = s.data[j], s.data[i]
	})
	return s
}

// PopLeft pops the rightmost element in Slice.
func (s *Slice[T]) PopLeft() *Slice[T] {
	if len(s.data) == 0 {
		return s
	}
	s.data = s.data[1:]
	return s
}

// PopRight pops the rightmost element in Slice.
func (s *Slice[T]) PopRight() *Slice[T] {
	if len(s.data) == 0 {
		return s
	}
	s.data = s.data[:len(s.data)-1]
	return s
}

// ------------------ split line ------------------------
// - Below are non-chain methods.

// Slice returns a copy of the slice.
func (s *Slice[T]) Slice() (copied []T) {
	copied = make([]T, len(s.data))
	copy(copied, s.data)
	return
}

// Len returns the length of the slice.
func (s *Slice[T]) Len() int {
	return len(s.data)
}

// Contains returns true if the slice contains the item.
func (s *Slice[T]) Contains(item T) bool {
	for _, i := range s.data {
		if i == item {
			return true
		}
	}
	return false
}

// Reduce reduces the slice to a single value.
func (s *Slice[T]) Reduce(f func(x, y T) T) T {
	if len(s.data) == 0 {
		var zero T
		return zero
	}
	result := s.data[0]
	for _, item := range s.data[1:] {
		result = f(result, item)
	}
	return result
}

// Equal returns true if the slice is equal to the other Slice.
func (s *Slice[T]) Equal(other *Slice[T]) bool {
	if len(s.data) != len(other.data) {
		return false
	}
	for i, item := range s.data {
		if item != other.data[i] {
			return false
		}
	}
	return true
}

// IndexOf returns the index of the item in the slice. If the item is not found, it returns -1.
func (s *Slice[T]) IndexOf(item T) int {
	for i, v := range s.data {
		if v == item {
			return i
		}
	}
	return -1
}

// Get returns the item at the given index.
func (s *Slice[T]) Get(index int) T {
	return s.data[index]
}

// Join joins the elements in Slice by the given separator.
func (s *Slice[T]) Join(sep string) string {
	var ss []string
	for _, item := range s.data {
		ss = append(ss, fmt.Sprintf("%v", item))
	}
	return strings.Join(ss, sep)
}

// IsEmpty returns true if the underlying slice is empty.
func (s *Slice[T]) IsEmpty() bool {
	return len(s.data) == 0
}

package sslice

import "github.com/chaseSpace/bear/constraints"

type ComputableSlice[T constraints.Computable] struct {
	slice *Slice[T]
}

func NewComputableSlice[T constraints.Computable](data ...T) *ComputableSlice[T] {
	return &ComputableSlice[T]{slice: New(data...)}
}

// Append appends data to the end of ComputableSlice.
func (s *ComputableSlice[T]) Append(data ...T) *ComputableSlice[T] {
	s.slice.Append(data...)
	return s
}

// Clone returns a copy of ComputableSlice.
func (s *ComputableSlice[T]) Clone() *ComputableSlice[T] {
	return &ComputableSlice[T]{slice: s.slice.Clone()}
}

// Filter filters the elements in ComputableSlice by the given function.
func (s *ComputableSlice[T]) Filter(f func(T) bool) *ComputableSlice[T] {
	s.slice.Filter(f)
	return s
}

// Map maps the elements in ComputableSlice by the given function.
func (s *ComputableSlice[T]) Map(f func(T) T) *ComputableSlice[T] {
	s.slice.Map(f)
	return s
}

// Unique removes duplicate elements in ComputableSlice.
func (s *ComputableSlice[T]) Unique() *ComputableSlice[T] {
	s.slice.Unique()
	return s
}

// Reverse reverses the elements in ComputableSlice.
func (s *ComputableSlice[T]) Reverse() *ComputableSlice[T] {
	s.slice.Reverse()
	return s
}

// Shuffle shuffles the elements in ComputableSlice.
func (s *ComputableSlice[T]) Shuffle() *ComputableSlice[T] {
	s.slice.Shuffle()
	return s
}

// PopLeft pops the leftmost element in ComputableSlice.
func (s *ComputableSlice[T]) PopLeft() *ComputableSlice[T] {
	s.slice.PopLeft()
	return s
}

// PopRight pops the rightmost element in ComputableSlice.
func (s *ComputableSlice[T]) PopRight() *ComputableSlice[T] {
	s.slice.PopRight()
	return s
}

// ------------------ split line ------------------------
// - Below are non-chain methods.

// Sum returns the sum of all elements in the slice.
func (s *ComputableSlice[T]) Sum() T {
	var sum T
	for _, item := range s.slice.data {
		sum += item
	}
	return sum
}

// Slice returns a copy of the elements in ComputableSlice.
func (s *ComputableSlice[T]) Slice() (copied []T) {
	return s.slice.Slice()
}

// Len returns the length of ComputableSlice.
func (s *ComputableSlice[T]) Len() int {
	return s.slice.Len()
}

// Contains returns true if the element is in ComputableSlice.
func (s *ComputableSlice[T]) Contains(item T) bool {
	return s.slice.Contains(item)
}

// Reduce reduces the elements in ComputableSlice by the given function.
func (s *ComputableSlice[T]) Reduce(f func(x, y T) T) T {
	return s.slice.Reduce(f)
}

// Equal returns true if the elements in ComputableSlice are equal to the elements in others.
func (s *ComputableSlice[T]) Equal(other *ComputableSlice[T]) bool {
	return s.slice.Equal(other.slice)
}

// IndexOf returns the index of the element in ComputableSlice.
func (s *ComputableSlice[T]) IndexOf(item T) int {
	return s.slice.IndexOf(item)
}

// Join joins the elements in ComputableSlice by the given separator.
func (s *ComputableSlice[T]) Join(sep string) string {
	return s.slice.Join(sep)
}

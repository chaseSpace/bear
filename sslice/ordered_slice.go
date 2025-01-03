package sslice

import (
	"github.com/chaseSpace/bear/constraints"
	"sort"
)

type OrderedSlice[T constraints.Ordered] struct {
	slice *Slice[T]
}

func NewOrderedSlice[T constraints.Ordered](data ...T) *OrderedSlice[T] {
	return &OrderedSlice[T]{slice: New[T](data...)}
}

// Sort sorts the elements in OrderedSlice in ascending order by default. If true is passed,
// it sorts in descending order.
func (s *OrderedSlice[T]) Sort(desc ...bool) *OrderedSlice[T] {
	if len(desc) > 0 && desc[0] { // desc sort
		sort.Slice(s.slice.data, func(i, j int) bool {
			return s.slice.data[i] > s.slice.data[j]
		})
		return s
	}
	// asc sort
	sort.Slice(s.slice.data, func(i, j int) bool {
		return s.slice.data[i] < s.slice.data[j]
	})
	return s
}

// Append appends data to the end of OrderedSlice.
func (s *OrderedSlice[T]) Append(data ...T) *OrderedSlice[T] {
	s.slice.Append(data...)
	return s
}

// Clone returns a copy of OrderedSlice.
func (s *OrderedSlice[T]) Clone() *OrderedSlice[T] {
	return &OrderedSlice[T]{slice: s.slice.Clone()}
}

// Filter filters the elements in OrderedSlice by the given function.
func (s *OrderedSlice[T]) Filter(f func(T) bool) *OrderedSlice[T] {
	s.slice.Filter(f)
	return s
}

// Map maps the elements in OrderedSlice by the given function.
func (s *OrderedSlice[T]) Map(f func(T) T) *OrderedSlice[T] {
	s.slice.Map(f)
	return s
}

// Unique removes duplicate elements in OrderedSlice.
func (s *OrderedSlice[T]) Unique() *OrderedSlice[T] {
	s.slice.Unique()
	return s
}

// Reverse reverses the elements in OrderedSlice.
func (s *OrderedSlice[T]) Reverse() *OrderedSlice[T] {
	s.slice.Reverse()
	return s
}

// Shuffle shuffles the elements in OrderedSlice.
func (s *OrderedSlice[T]) Shuffle() *OrderedSlice[T] {
	s.slice.Shuffle()
	return s
}

// PopLeft pops the leftmost element in OrderedSlice.
func (s *OrderedSlice[T]) PopLeft() *OrderedSlice[T] {
	s.slice.PopLeft()
	return s
}

// PopRight pops the rightmost element in OrderedSlice.
func (s *OrderedSlice[T]) PopRight() *OrderedSlice[T] {
	s.slice.PopRight()
	return s
}

// ------------------ split line ------------------------
// - Below are non-chain methods.

// Slice returns a copy of the elements in OrderedSlice.
func (s *OrderedSlice[T]) Slice() (copied []T) {
	return s.slice.Slice()
}

// Len returns the length of OrderedSlice.
func (s *OrderedSlice[T]) Len() int {
	return s.slice.Len()
}

// Contains returns true if the element is in OrderedSlice.
func (s *OrderedSlice[T]) Contains(item T) bool {
	return s.slice.Contains(item)
}

// Reduce reduces the elements in OrderedSlice by the given function.
func (s *OrderedSlice[T]) Reduce(f func(x, y T) T) T {
	return s.slice.Reduce(f)
}

// Equal returns true if the elements in OrderedSlice are equal to the elements in others.
func (s *OrderedSlice[T]) Equal(other *OrderedSlice[T]) bool {
	return s.slice.Equal(other.slice)
}

// IndexOf returns the index of the element in OrderedSlice.
func (s *OrderedSlice[T]) IndexOf(item T) int {
	return s.slice.IndexOf(item)
}

// Join joins the elements in OrderedSlice by the given separator.
func (s *OrderedSlice[T]) Join(sep string) string {
	return s.slice.Join(sep)
}

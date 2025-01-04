package sset

import (
	"fmt"
	"strings"
)

type Set[T comparable] struct {
	data map[T]struct{}
}

// New returns a new Set
func New[T comparable](items ...T) *Set[T] {
	var data = make(map[T]struct{})
	for _, item := range items {
		data[item] = struct{}{}
	}
	return &Set[T]{data: data}
}

// ---------------------- Chained Methods ----------------------

// Add adds data to Set
func (s *Set[T]) Add(data ...T) *Set[T] {
	for _, item := range data {
		s.data[item] = struct{}{}
	}
	return s
}

// Clone returns a copy of Set
func (s *Set[T]) Clone() *Set[T] {
	var data = make(map[T]struct{})
	for k, v := range s.data {
		data[k] = v
	}
	return &Set[T]{data: data}
}

// Filter filters Set by f
func (s *Set[T]) Filter(f func(T) bool) *Set[T] {
	for k := range s.data {
		if !f(k) {
			delete(s.data, k)
		}
	}
	return s
}

// Clear clears Set
func (s *Set[T]) Clear() *Set[T] {
	s.data = make(map[T]struct{})
	return s
}

// Delete deletes data from Set
func (s *Set[T]) Delete(data ...T) *Set[T] {
	for _, item := range data {
		delete(s.data, item)
	}
	return s
}

// ForEach iterates Set by f
func (s *Set[T]) ForEach(f func(T)) *Set[T] {
	for k := range s.data {
		f(k)
	}
	return s
}

// Map maps Set by f
func (s *Set[T]) Map(f func(T) T) *Set[T] {
	var data = make(map[T]struct{})
	for k := range s.data {
		data[f(k)] = struct{}{}
		delete(s.data, k)
	}
	s.data = data
	return s
}

// Merge merges Set to self with other
func (s *Set[T]) Merge(other *Set[T]) *Set[T] {
	for k := range other.data {
		s.data[k] = struct{}{}
	}
	return s
}

// Intersect return a new Set, which is the intersection of Set and other
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	var newSet = New[T]()
	for k := range s.data {
		if other.Has(k) {
			newSet.Add(k)
		}
	}
	return newSet
}

// Union returns a new Set which is the union of Set and others
func (s *Set[T]) Union(others ...*Set[T]) *Set[T] {
	var newSet = s.Clone()
	for _, other := range others {
		for k := range other.data {
			newSet.Add(k)
		}
	}
	return newSet
}

// IsSubsetOf checks if Set is a subset of other
func (s *Set[T]) IsSubsetOf(other *Set[T]) bool {
	for k := range s.data {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

// Diff returns a new Set which is the difference set from `set` to `others`.
// Which means, all the items in `newSet` are in `set` but not in `others`.
func (s *Set[T]) Diff(other *Set[T]) *Set[T] {
	var newSet = New[T]()
	for k := range s.data {
		if _, ok := other.data[k]; !ok {
			newSet.data[k] = struct{}{}
		}
	}
	return newSet
}

// ---------------------- Non-Chained Methods ----------------------

// Slice returns a slice of Set
func (s *Set[T]) Slice() (copied []T) {
	for k := range s.data {
		copied = append(copied, k)
	}
	return copied
}

// Size returns the size of Set
func (s *Set[T]) Size() int {
	return len(s.data)
}

// Has checks if Set contains item
func (s *Set[T]) Has(item T) bool {
	_, ok := s.data[item]
	return ok
}

// Equal checks if Set is equal to other
func (s *Set[T]) Equal(other *Set[T]) bool {
	if s.Size() != other.Size() {
		return false
	}
	for k := range s.data {
		if !other.Has(k) {
			return false
		}
	}
	return true
}

// Join joins Set by sep
func (s *Set[T]) Join(sep string) string {
	var sb strings.Builder
	for k := range s.data {
		sb.WriteString(fmt.Sprintf("%v", k))
		sb.WriteString(sep)
	}
	return strings.TrimRight(sb.String(), sep)
}

// IsEmpty checks if Set is empty
func (s *Set[T]) IsEmpty() bool {
	return len(s.data) == 0
}

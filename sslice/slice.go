package sslice

type Slice[T comparable] struct {
	data []T
}

func New[T comparable](data ...T) *Slice[T] {
	return &Slice[T]{data: data}
}

// Append appends data to the slice. It returns the slice itself.
func (s *Slice[T]) Append(data ...T) *Slice[T] {
	s.data = append(s.data, data...)
	return s
}

func (s *Slice[T]) Clone() *Slice[T] {
	copied := make([]T, len(s.data))
	copy(copied, s.data)
	return &Slice[T]{data: copied}
}

// Filter filters the slice. It returns the slice itself.
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

// Map maps the slice. It returns the slice itself.
func (s *Slice[T]) Map(f func(T) T) *Slice[T] {
	for i, item := range s.data {
		s.data[i] = f(item)
	}
	return s
}

func (s *Slice[T]) Contains(item T) bool {
	for _, i := range s.data {
		if i == item {
			return true
		}
	}
	return false
}

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

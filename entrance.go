package bear

import (
	"github.com/chaseSpace/bear/constraints"
	"github.com/chaseSpace/bear/sset"
	"github.com/chaseSpace/bear/sslice"
)

// NewSlice creates a new instance of Slice.
func NewSlice[T comparable](data ...T) *sslice.Slice[T] {
	return sslice.New(data...)
}

// NewComputableSlice creates a new instance of ComputableSlice.
func NewComputableSlice[T constraints.Computable](data ...T) *sslice.ComputableSlice[T] {
	return sslice.NewComputableSlice(data...)
}

// NewOrderedSlice creates a new instance of OrderedSlice.
func NewOrderedSlice[T constraints.Ordered](data ...T) *sslice.OrderedSlice[T] {
	return sslice.NewOrderedSlice(data...)
}

// NewSet creates a new instance of Set.
func NewSet[T comparable](data ...T) *sset.Set[T] {
	return sset.New(data...)
}

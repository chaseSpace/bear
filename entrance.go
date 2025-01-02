package sabre

import "github.com/chasespace/sabre/sslice"

func NewSlice[T comparable](data ...T) *sslice.Slice[T] {
	return sslice.New(data...)
}

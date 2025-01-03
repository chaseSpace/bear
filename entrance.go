package bear

import "github.com/chaseSpace/bear/sslice"

func NewSlice[T comparable](data ...T) *sslice.Slice[T] {
	return sslice.New(data...)
}

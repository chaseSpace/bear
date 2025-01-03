# bear

[![Go Report Card](https://goreportcard.com/badge/github.com/chasespace/bear)](https://goreportcard.com/report/github.com/chaseSpace/bear)
[![Go Reference](https://pkg.go.dev/badge/github.com/chasespace/bear.svg)](https://pkg.go.dev/github.com/chaseSpace/bear)

bear, focusing on **Data Structure Processing** (Using Generic) in golang.

**NOTE**: All APIs is **not** concurrency-safe. We're doing less work like most of stdlib APIs, so you also should use it like
use stdlib APIs.

Min Go version: 1.18

## Download

```shell
go get -u github.com/chaseSpace/bear
```

## Import

```
import "github.com/chaseSpace/bear"
```

## Quick Start

```go
package main

import (
	"fmt"
	"github.com/chaseSpace/bear"
)

func main() {
	// Use common methods of Slice
	var s1 = bear.NewSlice("1", "2", "2", "3")
	s1.Unique().Reverse().Map(func(i string) string { return i + "x" })
	fmt.Println(s1.Slice()) // 3x 2x 1x

	// Use Sort() of OrderedSlice
	var s2 = bear.NewOrderedSlice(1, 2, 2, 3)
	var joined = s2.Unique().Sort(true).Join(".")
	fmt.Println(joined) // 3.2.1

	// Use Sum() of ComputableSlice
	var s3 = bear.NewComputableSlice(2.0, 3.0, 4.0)
	var sum = s3.Append(5.0).Sum()
	fmt.Println("Sum:", sum) // 14
}

```

For now, only support slice operations.

### Slice API Documentation

The `Slice` type provides a convenient interface for common slice operations.

| Method                   | Description                                                                               |
|--------------------------|-------------------------------------------------------------------------------------------|
| `Append(data ...T)`      | Appends new elements to the slice and returns the updated slice.                          |
| `Clone()`                | Returns a new slice that is a copy of the original slice.                                 |
| `Filter(f func(T) bool)` | Filters the slice based on the provided predicate function and returns the slice.         |
| `Map(f func(T) T)`       | Applies a function to each element of the slice and returns the mapped slice.             |
| `Unique()`               | Removes duplicate items from the slice and returns the slice.                             |
| `Reverse()`              | Reverses the order of elements in the slice and returns the slice.                        |
| `Shuffle()`              | Randomly shuffles the elements in the slice and returns the slice.                        |
| `PopLeft()`              | PopLeft pops the leftmost element in Slice.                                               |
| `PopRight()`             | PopRight pops the rightmost element in Slice.                                             |
| `Sort(desc ...bool)`     | [**OrderedSlice**] Sort sorts the elements in OrderedSlice in ascending order by default. |

#### Slice Non-Chain Methods

These methods do not return a pointer to the `Slice` type, hence they do not support method chaining.

| Method                        | Description                                                             |
|-------------------------------|-------------------------------------------------------------------------|
| `Slice() []T`                 | Returns a copy of the slice as a standard Go slice.                     |
| `Len() int`                   | Returns the length of the slice.                                        |
| `Contains(item T) bool`       | Checks if the slice contains a specific item and returns a boolean.     |
| `Reduce(f func(x, y T) T) T`  | Reduces the slice to a single value by applying a function.             |
| `Equal(other *Slice[T]) bool` | Compares the slice with another slice and returns a boolean.            |
| `IndexOf(item T) int`         | Returns the index of a specific item or -1 if not found.                |
| `Get(index int) T`            | Returns the item at the given index.                                    |
| `Sum()`                       | [**ComputableSlice**] Sum returns the sum of all elements in the slice. |

## License

MIT License.
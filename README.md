# bear

[![Go Report Card](https://goreportcard.com/badge/github.com/chasespace/bear)](https://goreportcard.com/report/github.com/chaseSpace/bear)
[![Go Reference](https://pkg.go.dev/badge/github.com/chasespace/bear.svg)](https://pkg.go.dev/github.com/chaseSpace/bear)

bear, focusing on **Data Structure Processing** (Using Generic) in golang.

**NOTE**: All APIs is **not** concurrency-safe. We're doing less work like most of stdlib APIs, so you also should use
it like
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
	var sum = s3.Append(5.0).PopLeft().Append(6.0).Avg()
	fmt.Println("Sum:", sum) // 4.55
	// For ComputableSlice, you can also call Max()/Min()/Avg().
}

```

For now, only support slice operations.

### Slice API Documentation

TheSlicetype provides a convenient interface for common slice operations.

| Method     | Description                                                                                               |
|------------|-----------------------------------------------------------------------------------------------------------|
| Append()   | Appends new elements to the slice and returns the updated slice.                                          |
| Clone()    | Returns a new slice that is a copy of the original slice.                                                 |
| Filter()   | Filters the slice based on the provided predicate function and returns the slice.                         |
| Map()      | Applies a function to each element of the slice and returns the mapped slice.                             |
| Unique()   | Removes duplicate items from the slice and returns the slice.                                             |
| Reverse()  | Reverses the order of elements in the slice and returns the slice.                                        |
| Shuffle()  | Randomly shuffles the elements in the slice and returns the slice.                                        |
| PopLeft()  | PopLeft pops the leftmost element in Slice.                                                               |
| PopRight() | PopRight pops the rightmost element in Slice.                                                             |
| Sort()     | [**OrderedSlice/ComputableSlice**] Sort sorts the elements in OrderedSlice in ascending order by default. |

#### Slice Non-Chain Methods

These methods do not return a pointer to theSlicetype, hence they do not support method chaining.

| Method     | Description                                                                           |
|------------|---------------------------------------------------------------------------------------|
| Slice()    | Returns a copy of the slice as a standard Go slice.                                   |
| Len()      | Returns the length of the slice.                                                      |
| Contains() | Checks if the slice contains a specific item and returns a boolean.                   |
| Reduce()   | Reduces the slice to a single value by applying a function.                           |
| Equal()    | Compares the slice with another slice and returns a boolean.                          |
| IndexOf()  | Returns the index of a specific item or -1 if not found.                              |
| Get()      | Returns the item at the given index.                                                  |
| Sum()      | [**ComputableSlice**] Sum returns the sum of all elements in the ComputableSlice.     |
| Max()      | [**ComputableSlice**] Max returns the maximum value in the ComputableSlice.           |
| Min()      | [**ComputableSlice**] Min returns the minimum value in the ComputableSlice.           |
| Avg()      | [**ComputableSlice**] Avg returns the average of all elements in the ComputableSlice. |
| IsEmpty()  | Checks if the underlying slice is empty.                                              |

### Set API Documentation

TheSettype provides a convenient interface for common set operations.

| Method       | Description                                                             |
|--------------|-------------------------------------------------------------------------|
| Add()        | adds data to Set.                                                       |
| Clone()      | Clone returns a copy of Set.                                            |
| Filter()     | Filter filters Set by f.                                                |
| Clear()      | Clear clears Set.                                                       |
| Delete()     | Delete deletes data from Set.                                           |
| ForEach()    | ForEach iterates Set by f.                                              |
| Map()        | Map maps Set by f.                                                      |
| Merge()      | Merge merges Set to self with other.                                    |
| Intersect()  | Intersect return a new Set, which is the intersection of Set and other. |
| Union()      | Union returns a new Set which is the union set ofsetandothers`.         |
| IsSubsetOf() | IsSubsetOf checks if Set is a subset of other.                          |
| Diff()       | Diff returns a new Set which is the difference set fromsettoothers`.    |

#### Set Non-Chain Methods

These methods do not return a pointer to theSettype, hence they do not support method chaining.

| Method    | Description                                                         |
|-----------|---------------------------------------------------------------------|
| Slice()   | Returns a copy of the set as a standard Go slice.                   |
| Size()    | Returns the length of the set.                                      |
| Has()     | Checks if the set contains a specific item and returns a boolean.   |
| Equal()   | Compares the set with another set and returns a boolean.            |
| Join()    | Joins the set elements into a string using the specified separator. |
| IsEmpty() | Checks if the set is empty and returns a boolean.                   |

## License

MIT License.
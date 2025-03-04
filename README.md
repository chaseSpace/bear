# bear

[![Go Report Card](https://goreportcard.com/badge/github.com/chasespace/bear)](https://goreportcard.com/report/github.com/chaseSpace/bear)
[![Go Reference](https://pkg.go.dev/badge/github.com/chasespace/bear.svg)](https://pkg.go.dev/github.com/chaseSpace/bear)

bear, focusing on **Data Structure Processing** (Using Generic) in golang.

**NOTE**: All APIs are **not** concurrency-safe. We're doing less work like most of the stdlib APIs, so you also should
use it like use stdlib APIs.

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

### 1. Slice API Documentation

The Slice type provides a convenient interface for common slice operations.

| Method     | Description                                                                                          |
|------------|------------------------------------------------------------------------------------------------------|
| `Append`   | Appends new elements to the slice and returns the updated slice.                                     |
| `Clone`    | Returns a new slice that is a copy of the original slice.                                            |
| `Filter`   | Filter removes elements that match f function from underlying slice.                                 |
| `Map`      | Applies a function to each element of the slice and returns the mapped slice.                        |
| `Unique`   | Removes duplicate items from the slice and returns the slice.                                        |
| `Reverse`  | Reverses the order of elements in the slice and returns the slice.                                   |
| `Shuffle`  | Randomly shuffles the elements in the slice and returns the slice.                                   |
| `PopLeft`  | PopLeft pops the leftmost element in Slice.                                                          |
| `PopRight` | PopRight pops the rightmost element in Slice.                                                        |
| `Sort`     | [**OrderedSlice/ComputableSlice**] Sorts the elements in OrderedSlice in ascending order by default. |

#### Slice Non-Chain Methods

These methods do not return a pointer to theSlicetype, hence they do not support method chaining.

| Method     | Description                                                                           |
|------------|---------------------------------------------------------------------------------------|
| `Slice`    | Returns a copy of the slice as a standard Go slice.                                   |
| `Len`      | Returns the length of the slice.                                                      |
| `Contains` | Checks if the slice contains a specific item and returns a boolean.                   |
| `Reduce`   | Reduces the slice to a single value by applying a function.                           |
| `Equal`    | Compares the slice with another slice and returns a boolean.                          |
| `IndexOf`  | Returns the index of a specific item or -1 if not found.                              |
| `Get`      | Returns the item at the given index.                                                  |
| `Sum`      | [**ComputableSlice**] Sum returns the sum of all elements in the ComputableSlice.     |
| `Max`      | [**ComputableSlice**] Max returns the maximum value in the ComputableSlice.           |
| `Min`      | [**ComputableSlice**] Min returns the minimum value in the ComputableSlice.           |
| `Avg`      | [**ComputableSlice**] Avg returns the average of all elements in the ComputableSlice. |
| `IsEmpty`  | Checks if the underlying slice is empty.                                              |

### 2. Set API Documentation

The Set type provides a convenient interface for common set operations.

| Method       | Description                                                             |
|--------------|-------------------------------------------------------------------------|
| `Add`        | Adds data to Set.                                                       |
| `Clone`      | Returns a copy of Set.                                                  |
| `Filter`     | Removes elements that match the function `f` from the underlying slice. |
| `Clear`      | Clears Set.                                                             |
| `Delete`     | Deletes data from Set.                                                  |
| `ForEach`    | Iterates Set by applying function `f`.                                  |
| `Map`        | Maps Set by applying function `f`.                                      |
| `Merge`      | Merges Set to self with another Set.                                    |
| `Intersect`  | Returns a new Set, which is the intersection of Set and another Set.    |
| `Union`      | Returns a new Set which is the union of Set and others.                 |
| `IsSubsetOf` | Checks if Set is a subset of another Set.                               |
| `Diff`       | Returns a new Set which is the difference set from Set to others.       |

#### Set Non-Chain Methods

These methods do not return a pointer to the Set type, hence they do not support method chaining.

| Method    | Description                                                         |
|-----------|---------------------------------------------------------------------|
| `Slice`   | Returns a copy of the set as a standard Go slice.                   |
| `Size`    | Returns the length of the set.                                      |
| `Has`     | Checks if the set contains a specific item and returns a boolean.   |
| `Equal`   | Compares the set with another set and returns a boolean.            |
| `Join`    | Joins the set elements into a string using the specified separator. |
| `IsEmpty` | Checks if the set is empty and returns a boolean.                   |

### 3. SinglyLinkedList API Documentation

The SinglyLinkedList type provides a convenient interface for common **singly** linked list operations.

| Method         | Description                                                                          |
|----------------|--------------------------------------------------------------------------------------|
| `Append`       | Adds one or more values to the end of the linked list.                               |
| `InsertBefore` | Inserts a new node with the specified value before the node at the specified index.  |
| `InsertAfter`  | Inserts a new node with the specified value after the node at the specified index.   |
| `Remove`       | Removes the node at the specified index.                                             |
| `IndexOf`      | Returns the index of the first occurrence of the specified value in the linked list. |
| `Find`         | Returns the node at the specified index.                                             |
| `Update`       | Updates the value of the node at the specified index.                                |
| `Walk`         | Applies a function to each node in the linked list.                                  |
| `Reverse`      | Reverses the linked list.                                                            |
| `Merge`        | Merges the current linked list with another linked list.                             |
| `ToSlice`      | Converts all elements from the linked list to a slice.                               |
| `Length`       | Returns the length of the linked list.                                               |
| `IsEmpty`      | Checks if the linked list is empty.                                                  |
| `String`       | Returns a string representation of the linked list.                                  |
| `CountOf`      | Count occurrences of a specific value in the linked list.                            |

> [!NOTE]
> This type does not support method chaining.

### 4. DoublyLinkedList API Documentation

The DoublyLinkedList type provides a convenient interface for common **doubly** linked list operations.

| Method         | Description                                                                          |
|----------------|--------------------------------------------------------------------------------------|
| `Append`       | Adds one or more values to the end of the linked list.                               |
| `InsertBefore` | Inserts a new node with the specified value before the node at the specified index.  |
| `InsertAfter`  | Inserts a new node with the specified value after the node at the specified index.   |
| `Remove`       | Removes the node at the specified index.                                             |
| `IndexOf`      | Returns the index of the first occurrence of the specified value in the linked list. |
| `Find`         | Returns the node at the specified index.                                             |
| `Update`       | Updates the value of the node at the specified index.                                |
| `Walk`         | Applies a function to each node in the linked list.                                  |
| `Reverse`      | Reverses the linked list.                                                            |
| `Merge`        | Merges the current linked list with another linked list.                             |
| `ToSlice`      | Converts all elements from the linked list to a slice.                               |
| `Length`       | Returns the length of the linked list.                                               |
| `IsEmpty`      | Checks if the linked list is empty.                                                  |
| `String`       | Returns a string representation of the linked list.                                  |
| `CountOf`      | Count occurrences of a specific value in the linked list.                            |

> [!NOTE]
> This type does not support method chaining.

## License

MIT License.
# bear

[![Go Report Card](https://goreportcard.com/badge/github.com/chasespace/bear)](https://goreportcard.com/report/github.com/chaseSpace/bear)

bear, Focusing on **Data Structure Processing** (Using Generic) in golang.

Min Go version: 1.18

## Download

```shell
go get -u github.com/chaseSpace/bear
```

## Import

```
import "github.com/chaseSpace/bear"
```

## Examples

For now, only support slice operations.

To create new slice:

```go
slice := bear.NewSlice(1, 2, 3, 4, 5) // there could be any "comparable" type for element
fmt.Println("Created slice:", slice.Slice()) // [1 2 3 4 5]
```

To append data to slice:

```go
slice.Append(6, 7, 8)
fmt.Println("Appended slice:", slice.Slice()) // [1 2 3 4 5 6 7 8]
```

To clone slice:

```go
clonedSlice := slice.Clone()
fmt.Println("Cloned slice:", clonedSlice.Slice()) // [1 2 3 4 5]
```

To keep even numbers in slice:

```go
filteredSlice := slice.Filter(func (item int) bool {
return item%2 == 0
})
fmt.Println("Filtered slice:", filteredSlice.Slice()) // [2 4]
```

To map per element`*`2 in slice:

```go
mappedSlice := slice.Map(func (item int) int {
return item * 2
})
fmt.Println("Mapped slice:", mappedSlice.Slice()) // [2 4 6 8 10]
```

To check if slice contains `3`:

```go
containsThree := slice.Contains(3)
fmt.Println("Contains 3:", containsThree) // true
```

To get after-duplicates slice:

```go
slice := bear.NewSlice(1, 2, 2, 4, 4)
uniqueSlice := slice.Unique()
fmt.Println("Unique slice:", uniqueSlice.Slice()) // [1 2 4]
```

To reduce slice to total:

```go
total := slice.Reduce(func (x, y int) int {
return x + y
})
fmt.Println("Total:", total) // 15
```

## License

MIT License.
# Slices in Go

Slices are dynamic, flexible views into arrays. They're one of the most important and commonly used data structures in Go.

## Table of Contents

- [What are Slices?](#what-are-slices)
- [Slice Internals](#slice-internals)
- [Declaration and Initialization](#declaration-and-initialization)
- [Slice Operations](#slice-operations)
- [Capacity and Length](#capacity-and-length)
- [Appending Elements](#appending-elements)
- [Copying Slices](#copying-slices)
- [Slicing Slices](#slicing-slices)
- [Multi-dimensional Slices](#multi-dimensional-slices)
- [Slice as Function Parameters](#slice-as-function-parameters)
- [Common Patterns](#common-patterns)
- [Performance Considerations](#performance-considerations)
- [Best Practices](#best-practices)

## What are Slices?

A slice is:

- **Dynamic** - Can grow and shrink
- **Reference type** - Points to an underlying array
- **Flexible** - Size not part of the type
- **Three components** - Pointer, length, and capacity

```go
// Slice type (size not specified)
var s []int

// Array type (size required)
var a [5]int
```

## Slice Internals

A slice is a descriptor containing:

1. **Pointer** to the underlying array
2. **Length** - number of elements in the slice
3. **Capacity** - number of elements in underlying array (from slice start)

```go
type slice struct {
    ptr *ElementType  // Pointer to array
    len int           // Length
    cap int           // Capacity
}
```

### Visual Representation

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := arr[1:4]  // [2, 3, 4]

// slice points to arr, starting at index 1
// len(slice) = 3
// cap(slice) = 4 (from index 1 to end of array)
```

## Declaration and Initialization

### Nil Slice

The zero value of a slice is `nil`.

```go
var s []int
fmt.Println(s == nil)  // true
fmt.Println(len(s))    // 0
fmt.Println(cap(s))    // 0

// Nil slice is safe to use with append
s = append(s, 1, 2, 3)
```

### Empty Slice

An empty slice is not nil but has zero length.

```go
// Empty slice (not nil)
s := []int{}
fmt.Println(s == nil)  // false
fmt.Println(len(s))    // 0
fmt.Println(cap(s))    // 0

// Also empty slice
s := make([]int, 0)
```

### Slice Literal

```go
// Initialize with values
numbers := []int{1, 2, 3, 4, 5}
names := []string{"Alice", "Bob", "Charlie"}

// Composite literal with structs
type Person struct {
    Name string
    Age  int
}

people := []Person{
    {"Alice", 30},
    {"Bob", 25},
}
```

### Using make()

`make()` creates a slice with specified length and optional capacity.

```go
// make([]Type, length, capacity)

// Length 5, capacity 5 (zero values)
s1 := make([]int, 5)       // [0 0 0 0 0]

// Length 3, capacity 10
s2 := make([]int, 3, 10)   // [0 0 0]

// Length 0, capacity 5 (empty but with capacity)
s3 := make([]int, 0, 5)    // []
```

### From Array

```go
arr := [5]int{1, 2, 3, 4, 5}

// Slice from array
slice := arr[:]       // [1 2 3 4 5]
slice = arr[1:4]      // [2 3 4]
slice = arr[:3]       // [1 2 3]
slice = arr[2:]       // [3 4 5]
```

## Slice Operations

### Accessing Elements

```go
s := []int{10, 20, 30, 40, 50}

// Reading
first := s[0]   // 10
last := s[4]    // 50

// Writing
s[0] = 100      // [100 20 30 40 50]

// Bounds checking (panic if out of range)
// value := s[10]  // Panic: index out of range
```

### Iteration

```go
s := []int{1, 2, 3, 4, 5}

// Range with index and value
for i, v := range s {
    fmt.Printf("s[%d] = %d\n", i, v)
}

// Value only
for _, v := range s {
    fmt.Println(v)
}

// Index only
for i := range s {
    fmt.Println(i)
}

// Traditional for loop
for i := 0; i < len(s); i++ {
    fmt.Println(s[i])
}
```

### Modifying During Iteration

```go
s := []int{1, 2, 3, 4, 5}

// Modify slice elements
for i := range s {
    s[i] *= 2  // Double each element
}
// s is now [2 4 6 8 10]

// Note: Range copies values
for i, v := range s {
    v *= 2         // Modifies copy, not original!
    s[i] *= 2      // This modifies original
}
```

## Capacity and Length

### Understanding Length and Capacity

```go
s := make([]int, 3, 5)

len(s)  // 3 - number of elements
cap(s)  // 5 - capacity of underlying array

// Visual: [0 0 0 _ _]
//         ^^^^^     - length
//         ^^^^^^^^^  - capacity
```

### Capacity Changes with Slicing

```go
s := []int{1, 2, 3, 4, 5}
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=5 cap=5

s2 := s[2:]
fmt.Printf("len=%d cap=%d\n", len(s2), cap(s2))  // len=3 cap=3

s3 := s[:2]
fmt.Printf("len=%d cap=%d\n", len(s3), cap(s3))  // len=2 cap=5
```

### Full Slice Expression

Control the capacity with three-index slicing: `a[low:high:max]`

```go
s := []int{1, 2, 3, 4, 5}

// Normal slicing
s1 := s[1:3]
fmt.Printf("len=%d cap=%d\n", len(s1), cap(s1))  // len=2 cap=4

// Full slice expression (limit capacity)
s2 := s[1:3:3]
fmt.Printf("len=%d cap=%d\n", len(s2), cap(s2))  // len=2 cap=2
```

## Appending Elements

### Basic Append

```go
var s []int

s = append(s, 1)           // [1]
s = append(s, 2, 3, 4)     // [1 2 3 4]

// Append slice to slice
s2 := []int{5, 6, 7}
s = append(s, s2...)       // [1 2 3 4 5 6 7]
```

### Append and Capacity Growth

When capacity is exceeded, Go allocates a new array:

```go
s := make([]int, 0, 2)
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=0 cap=2

s = append(s, 1)
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=1 cap=2

s = append(s, 2)
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=2 cap=2

s = append(s, 3)  // Capacity exceeded, reallocate
fmt.Printf("len=%d cap=%d\n", len(s), cap(s))  // len=3 cap=4 (doubled)
```

### Growth Strategy

Go typically doubles capacity when growing (up to 1024, then grows by ~25%).

```go
func demonstrateGrowth() {
    var s []int
    prevCap := 0

    for i := 0; i < 20; i++ {
        s = append(s, i)
        if cap(s) != prevCap {
            fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
            prevCap = cap(s)
        }
    }
}
// Output:
// len=1 cap=1
// len=2 cap=2
// len=3 cap=4
// len=5 cap=8
// len=9 cap=16
// len=17 cap=32
```

### Append vs Manual Growth

```go
// Bad: Inefficient, multiple reallocations
var s []int
for i := 0; i < 1000; i++ {
    s = append(s, i)
}

// Good: Pre-allocate capacity
s := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    s = append(s, i)
}

// Alternative: Pre-allocate and index
s := make([]int, 1000)
for i := 0; i < 1000; i++ {
    s[i] = i
}
```

## Copying Slices

### Using copy()

`copy(dst, src)` copies elements from src to dst, returns number copied.

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, len(src))

n := copy(dst, src)
fmt.Println(n)    // 5
fmt.Println(dst)  // [1 2 3 4 5]

// dst is independent of src
src[0] = 100
fmt.Println(dst)  // [1 2 3 4 5] - unchanged
```

### Partial Copy

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)

n := copy(dst, src)  // Copies min(len(dst), len(src))
fmt.Println(n)       // 3
fmt.Println(dst)     // [1 2 3]
```

### Copy Overlapping Slices

```go
s := []int{1, 2, 3, 4, 5}

// Shift elements (overlapping copy is safe)
copy(s[1:], s[:4])  // Copy [1 2 3 4] to positions [1 2 3 4]
fmt.Println(s)       // [1 1 2 3 4]
```

### Copy vs Append

```go
// copy(): Fixed destination size
src := []int{1, 2, 3}
dst := make([]int, 2)
copy(dst, src)      // Copies 2 elements

// append(): Grows destination
dst = append(dst, src...)  // Appends all elements
```

## Slicing Slices

### Basic Slicing

```go
s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

s[2:5]   // [2 3 4]
s[:4]    // [0 1 2 3]
s[6:]    // [6 7 8 9]
s[:]     // [0 1 2 3 4 5 6 7 8 9] (all elements)
```

### Slices Share Underlying Array

```go
original := []int{1, 2, 3, 4, 5}
slice1 := original[1:4]  // [2 3 4]
slice2 := original[2:5]  // [3 4 5]

// Modify through slice1
slice1[0] = 100  // Changes original[1]

fmt.Println(original)  // [1 100 3 4 5]
fmt.Println(slice1)    // [100 3 4]
fmt.Println(slice2)    // [3 4 5] - also affected
```

### Creating Independent Copy

```go
original := []int{1, 2, 3, 4, 5}

// Shared: sub-slice (not independent)
shared := original[1:4]

// Independent: use copy
independent := make([]int, 3)
copy(independent, original[1:4])

shared[0] = 100
fmt.Println(original)      // [1 100 3 4 5] - affected
fmt.Println(shared)        // [100 3 4]
fmt.Println(independent)   // [2 3 4] - independent
```

## Multi-dimensional Slices

### 2D Slices

```go
// Create 2D slice (slice of slices)
rows, cols := 3, 4
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols)
}

// Initialize with values
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access elements
value := matrix[0][1]  // 2
matrix[1][1] = 99      // Modify element

// Iterate
for i, row := range matrix {
    for j, val := range row {
        fmt.Printf("matrix[%d][%d] = %d\n", i, j, val)
    }
}
```

### Jagged Slices

Unlike arrays, slices can have varying lengths (jagged/ragged arrays).

```go
jagged := [][]int{
    {1},
    {2, 3},
    {4, 5, 6},
    {7, 8, 9, 10},
}

// Each row can have different length
for i, row := range jagged {
    fmt.Printf("Row %d has %d elements\n", i, len(row))
}
```

### 3D Slices

```go
// Create 3D slice
x, y, z := 2, 3, 4
cube := make([][][]int, x)
for i := range cube {
    cube[i] = make([][]int, y)
    for j := range cube[i] {
        cube[i][j] = make([]int, z)
    }
}

// Access: cube[x][y][z]
cube[0][1][2] = 42
```

## Slice as Function Parameters

### Slices are Reference Types

```go
func modify(s []int) {
    s[0] = 100  // Modifies original slice's underlying array
}

func main() {
    numbers := []int{1, 2, 3, 4, 5}
    modify(numbers)
    fmt.Println(numbers)  // [100 2 3 4 5] - modified
}
```

### Append in Function (Gotcha!)

```go
func appendValue(s []int, val int) {
    s = append(s, val)  // Modifies local slice header, not caller's
}

func main() {
    numbers := []int{1, 2, 3}
    appendValue(numbers, 4)
    fmt.Println(numbers)  // [1 2 3] - unchanged!
}

// Correct: Return new slice or use pointer
func appendValue(s []int, val int) []int {
    return append(s, val)
}

func main() {
    numbers := []int{1, 2, 3}
    numbers = appendValue(numbers, 4)
    fmt.Println(numbers)  // [1 2 3 4] - correct
}
```

### Returning Slices

```go
func createSlice(n int) []int {
    s := make([]int, n)
    for i := range s {
        s[i] = i
    }
    return s
}

func filterEven(numbers []int) []int {
    var result []int
    for _, n := range numbers {
        if n%2 == 0 {
            result = append(result, n)
        }
    }
    return result
}
```

## Common Patterns

### Filter

```go
func filter(s []int, predicate func(int) bool) []int {
    var result []int
    for _, v := range s {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5, 6}
evens := filter(numbers, func(n int) bool {
    return n%2 == 0
})  // [2 4 6]
```

### Map

```go
func mapSlice(s []int, transform func(int) int) []int {
    result := make([]int, len(s))
    for i, v := range s {
        result[i] = transform(v)
    }
    return result
}

// Usage
numbers := []int{1, 2, 3, 4, 5}
doubled := mapSlice(numbers, func(n int) int {
    return n * 2
})  // [2 4 6 8 10]
```

### Reduce

```go
func reduce(s []int, initial int, fn func(int, int) int) int {
    result := initial
    for _, v := range s {
        result = fn(result, v)
    }
    return result
}

// Usage: Sum
numbers := []int{1, 2, 3, 4, 5}
sum := reduce(numbers, 0, func(acc, n int) int {
    return acc + n
})  // 15
```

### Remove Element

```go
// Remove element at index
func remove(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}

// Remove element by value (first occurrence)
func removeValue(s []int, value int) []int {
    for i, v := range s {
        if v == value {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}

// Remove all occurrences of value
func removeAll(s []int, value int) []int {
    var result []int
    for _, v := range s {
        if v != value {
            result = append(result, v)
        }
    }
    return result
}
```

### Insert Element

```go
// Insert at index
func insert(s []int, index, value int) []int {
    s = append(s, 0)              // Expand slice
    copy(s[index+1:], s[index:])  // Shift elements
    s[index] = value              // Insert value
    return s
}

// Usage
numbers := []int{1, 2, 4, 5}
numbers = insert(numbers, 2, 3)  // [1 2 3 4 5]
```

### Reverse

```go
func reverse(s []int) {
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

// Return new reversed slice
func reversed(s []int) []int {
    result := make([]int, len(s))
    for i, v := range s {
        result[len(s)-1-i] = v
    }
    return result
}
```

### Unique Elements

```go
func unique(s []int) []int {
    seen := make(map[int]bool)
    var result []int

    for _, v := range s {
        if !seen[v] {
            seen[v] = true
            result = append(result, v)
        }
    }

    return result
}

// Usage
numbers := []int{1, 2, 2, 3, 3, 3, 4}
uniq := unique(numbers)  // [1 2 3 4]
```

### Contains

```go
func contains(s []int, value int) bool {
    for _, v := range s {
        if v == value {
            return true
        }
    }
    return false
}
```

### Stack Operations

```go
type Stack []int

func (s *Stack) Push(v int) {
    *s = append(*s, v)
}

func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return 0, false
    }
    index := len(*s) - 1
    value := (*s)[index]
    *s = (*s)[:index]
    return value, true
}

func (s *Stack) Peek() (int, bool) {
    if len(*s) == 0 {
        return 0, false
    }
    return (*s)[len(*s)-1], true
}
```

### Queue Operations

```go
type Queue []int

func (q *Queue) Enqueue(v int) {
    *q = append(*q, v)
}

func (q *Queue) Dequeue() (int, bool) {
    if len(*q) == 0 {
        return 0, false
    }
    value := (*q)[0]
    *q = (*q)[1:]
    return value, true
}
```

## Performance Considerations

### Pre-allocate Capacity

```go
// Bad: Multiple reallocations
var s []int
for i := 0; i < 10000; i++ {
    s = append(s, i)
}

// Good: Pre-allocate known size
s := make([]int, 0, 10000)
for i := 0; i < 10000; i++ {
    s = append(s, i)
}

// Best: When length is known
s := make([]int, 10000)
for i := 0; i < 10000; i++ {
    s[i] = i
}
```

### Avoid Memory Leaks

```go
// Potential memory leak: Large underlying array retained
func processFirstFive(data []int) []int {
    return data[:5]  // Keeps reference to entire data array
}

// Better: Copy to new slice
func processFirstFive(data []int) []int {
    result := make([]int, 5)
    copy(result, data[:5])
    return result  // Old array can be garbage collected
}
```

### Efficient Deletion

```go
// Order doesn't matter: Swap with last
func deleteFast(s []int, index int) []int {
    s[index] = s[len(s)-1]
    return s[:len(s)-1]
}

// Order matters: Shift elements
func deleteOrdered(s []int, index int) []int {
    return append(s[:index], s[index+1:]...)
}
```

### Reuse Slices

```go
// Bad: Creates new slice each time
for {
    buffer := make([]byte, 1024)
    // Use buffer
}

// Good: Reuse buffer
buffer := make([]byte, 1024)
for {
    // Reuse buffer
    buffer = buffer[:0]  // Reset length, keep capacity
    // Use buffer
}
```

## Best Practices

### 1. Prefer Slices Over Arrays

```go
// Preferred
func process(data []int) {}

// Avoid (unless you specifically need array semantics)
func process(data [100]int) {}
```

### 2. Check for nil Before Use

```go
func safePrint(s []int) {
    if s == nil {
        fmt.Println("nil slice")
        return
    }
    for _, v := range s {
        fmt.Println(v)
    }
}
```

### 3. Always Capture append Result

```go
// Wrong: append result not captured
numbers := []int{1, 2, 3}
append(numbers, 4)  // Bug: result ignored

// Correct
numbers = append(numbers, 4)
```

### 4. Use Full Slice Expression to Prevent Leaks

```go
// May retain large underlying array
small := large[:10]

// Limits capacity, allows garbage collection of large
small := large[:10:10]
```

### 5. Initialize With make for Known Size

```go
// Good: Pre-allocate
items := make([]Item, 0, expectedSize)

// Or if length is known
items := make([]Item, actualSize)
```

### 6. Don't Compare Slices Directly

```go
// Wrong: Won't compile
// if slice1 == slice2 {}

// Correct: Compare elements manually
func equal(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}

// Or use reflect.DeepEqual
import "reflect"
equal := reflect.DeepEqual(slice1, slice2)
```

### 7. Be Careful with Slice of Pointers

```go
// Pointers in slice keep objects alive
var items []*Item
for _, data := range largeDataset {
    item := &Item{data: data}
    items = append(items, item)
}
// All Items remain in memory

// Better: Use values or clear when done
items = nil  // Allow garbage collection
```

### 8. Document Slice Mutations

```go
// Good: Clear documentation
// ProcessItems modifies the input slice in-place.
func ProcessItems(items []Item) {
    // Modify items
}

// Good: Returns new slice
// FilterItems returns a new slice containing filtered items.
// The input slice is not modified.
func FilterItems(items []Item, condition func(Item) bool) []Item {
    // Return new slice
}
```

## Summary

- Slices are **dynamic**, **reference types** with pointer, length, and capacity
- **Nil slice** is valid and safe for most operations
- Use **make()** to pre-allocate capacity for better performance
- **append()** may reallocate; always capture the result
- **copy()** creates independent slices
- Slices **share underlying arrays** - beware of modifications
- **Pre-allocate capacity** when size is known or predictable
- Slices are **passed by reference** - modifications affect the caller
- Use **full slice expression** `[low:high:max]` to limit capacity
- Most Go code should use **slices**, not arrays

# Arrays in Go

Arrays are fixed-size sequences of elements of the same type. They form the foundation for understanding slices and are value types in Go.

## Table of Contents

- [What are Arrays?](#what-are-arrays)
- [Declaration and Initialization](#declaration-and-initialization)
- [Array Properties](#array-properties)
- [Accessing Elements](#accessing-elements)
- [Array Operations](#array-operations)
- [Multi-dimensional Arrays](#multi-dimensional-arrays)
- [Array as Function Parameters](#array-as-function-parameters)
- [Array Comparison](#array-comparison)
- [Common Patterns](#common-patterns)
- [Arrays vs Slices](#arrays-vs-slices)
- [Best Practices](#best-practices)

## What are Arrays?

An array in Go is:

- **Fixed-size** - Size is determined at compile time
- **Value type** - Arrays are copied when assigned or passed to functions
- **Homogeneous** - All elements must be the same type
- **Zero-indexed** - First element is at index 0
- **Contiguous** - Elements stored in consecutive memory locations

### Array Type

The array's size is part of its type:

```go
var arr1 [5]int     // Array of 5 integers
var arr2 [10]int    // Different type from arr1!
// arr1 = arr2      // Compile error: type mismatch
```

## Declaration and Initialization

### Zero Value

An array's zero value is an array where all elements are their type's zero value.

```go
var arr [5]int      // [0 0 0 0 0]
var strs [3]string  // ["" "" ""]
var bools [2]bool   // [false false]
```

### Declaration with Explicit Initialization

```go
// Method 1: List all elements
var numbers [5]int = [5]int{1, 2, 3, 4, 5}

// Method 2: Short declaration
primes := [5]int{2, 3, 5, 7, 11}

// Method 3: Let compiler count (...)
colors := [...]string{"red", "green", "blue"}  // Length is 3
```

### Partial Initialization

```go
// Initialize first few elements, rest are zero values
arr := [5]int{1, 2, 3}  // [1 2 3 0 0]

// Initialize specific indices
arr := [5]int{0: 10, 2: 20, 4: 30}  // [10 0 20 0 30]

// Mix both
arr := [5]int{1, 2, 4: 40}  // [1 2 0 0 40]
```

### Array of Structs

```go
type Point struct {
    X, Y int
}

// Array of structs
points := [3]Point{
    {X: 0, Y: 0},
    {X: 1, Y: 2},
    {X: 3, Y: 4},
}

// With field names
points := [3]Point{
    {X: 0, Y: 0},
    Point{X: 1, Y: 2},  // Type name optional after first
    {3, 4},             // Field names optional
}
```

## Array Properties

### Length

Use `len()` to get the array length (known at compile time).

```go
arr := [5]int{1, 2, 3, 4, 5}
length := len(arr)  // 5

// Length is part of the type
var a [5]int
var b [10]int
fmt.Printf("%T\n", a)  // [5]int
fmt.Printf("%T\n", b)  // [10]int
```

### Size in Memory

```go
import "unsafe"

arr := [5]int64{}
size := unsafe.Sizeof(arr)  // 40 bytes (5 * 8 bytes)
```

## Accessing Elements

### Index Access

```go
arr := [5]int{10, 20, 30, 40, 50}

// Reading elements
first := arr[0]   // 10
last := arr[4]    // 50

// Writing elements
arr[0] = 100      // [100 20 30 40 50]
arr[4] = 500      // [100 20 30 40 500]
```

### Bounds Checking

Go performs automatic bounds checking at runtime.

```go
arr := [3]int{1, 2, 3}

value := arr[2]   // OK
// value := arr[3]  // Runtime panic: index out of range
```

### Slicing Arrays

You can create a slice from an array (doesn't copy, creates a reference).

```go
arr := [5]int{1, 2, 3, 4, 5}

slice := arr[1:4]   // [2 3 4] - elements at indices 1, 2, 3
slice = arr[:3]     // [1 2 3] - first 3 elements
slice = arr[2:]     // [3 4 5] - from index 2 to end
slice = arr[:]      // [1 2 3 4 5] - all elements
```

## Array Operations

### Iteration

#### For Loop with Index

```go
arr := [5]int{1, 2, 3, 4, 5}

for i := 0; i < len(arr); i++ {
    fmt.Printf("arr[%d] = %d\n", i, arr[i])
}
```

#### Range Loop

```go
arr := [5]int{1, 2, 3, 4, 5}

// With index and value
for i, v := range arr {
    fmt.Printf("Index %d: Value %d\n", i, v)
}

// Value only
for _, v := range arr {
    fmt.Println(v)
}

// Index only
for i := range arr {
    fmt.Println(i)
}
```

### Copying Arrays

Arrays are value types, so assignment copies all elements.

```go
arr1 := [3]int{1, 2, 3}
arr2 := arr1              // Copy entire array

arr2[0] = 100
fmt.Println(arr1)         // [1 2 3] - unchanged
fmt.Println(arr2)         // [100 2 3]
```

### Manual Copy with Loop

```go
src := [5]int{1, 2, 3, 4, 5}
var dst [5]int

for i, v := range src {
    dst[i] = v
}
```

### Filling Array

```go
// Fill with specific value
var arr [10]int
for i := range arr {
    arr[i] = 42
}
```

### Finding Elements

```go
func contains(arr [5]int, target int) bool {
    for _, v := range arr {
        if v == target {
            return true
        }
    }
    return false
}

// Usage
numbers := [5]int{1, 2, 3, 4, 5}
found := contains(numbers, 3)  // true
```

### Finding Index

```go
func indexOf(arr [5]int, target int) int {
    for i, v := range arr {
        if v == target {
            return i
        }
    }
    return -1  // Not found
}

// Usage
numbers := [5]int{10, 20, 30, 40, 50}
index := indexOf(numbers, 30)  // 2
```

### Sum and Average

```go
func sum(arr [5]int) int {
    total := 0
    for _, v := range arr {
        total += v
    }
    return total
}

func average(arr [5]int) float64 {
    return float64(sum(arr)) / float64(len(arr))
}

// Usage
numbers := [5]int{10, 20, 30, 40, 50}
total := sum(numbers)      // 150
avg := average(numbers)    // 30.0
```

### Min and Max

```go
func min(arr [5]int) int {
    minVal := arr[0]
    for _, v := range arr[1:] {
        if v < minVal {
            minVal = v
        }
    }
    return minVal
}

func max(arr [5]int) int {
    maxVal := arr[0]
    for _, v := range arr[1:] {
        if v > maxVal {
            maxVal = v
        }
    }
    return maxVal
}
```

### Reverse Array

```go
func reverse(arr [5]int) [5]int {
    var reversed [5]int
    for i, v := range arr {
        reversed[len(arr)-1-i] = v
    }
    return reversed
}

// In-place reversal (modifies original)
func reverseInPlace(arr *[5]int) {
    for i := 0; i < len(arr)/2; i++ {
        j := len(arr) - 1 - i
        arr[i], arr[j] = arr[j], arr[i]
    }
}
```

## Multi-dimensional Arrays

### 2D Arrays

```go
// Declaration
var matrix [3][3]int

// Initialization
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Access elements
value := matrix[0][0]  // 1
matrix[1][1] = 99      // Set center to 99

// Iteration
for i := 0; i < len(matrix); i++ {
    for j := 0; j < len(matrix[i]); j++ {
        fmt.Printf("%d ", matrix[i][j])
    }
    fmt.Println()
}

// Range iteration
for i, row := range matrix {
    for j, val := range row {
        fmt.Printf("matrix[%d][%d] = %d\n", i, j, val)
    }
}
```

### 3D Arrays

```go
// 3D array: cube of 2x2x2
var cube [2][2][2]int = [2][2][2]int{
    {
        {1, 2},
        {3, 4},
    },
    {
        {5, 6},
        {7, 8},
    },
}

// Access: cube[x][y][z]
value := cube[0][1][0]  // 3
```

### Jagged Arrays (via Slices)

True jagged arrays aren't possible with fixed arrays, use slices:

```go
// Array of slices (not truly jagged array)
var jagged [3][]int = [3][]int{
    {1, 2},
    {3, 4, 5},
    {6},
}
```

## Array as Function Parameters

### Pass by Value

Arrays are passed by value (copied), which can be inefficient for large arrays.

```go
func modifyArray(arr [5]int) {
    arr[0] = 100  // Modifies the copy, not original
}

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    modifyArray(numbers)
    fmt.Println(numbers)  // [1 2 3 4 5] - unchanged
}
```

### Pass by Pointer (Recommended)

Pass array pointer to avoid copying and allow modifications.

```go
func modifyArray(arr *[5]int) {
    arr[0] = 100  // Modifies original array
}

func main() {
    numbers := [5]int{1, 2, 3, 4, 5}
    modifyArray(&numbers)
    fmt.Println(numbers)  // [100 2 3 4 5] - modified
}
```

### Return Array

```go
func createArray() [5]int {
    return [5]int{1, 2, 3, 4, 5}
}

func doubleArray(arr [5]int) [5]int {
    var result [5]int
    for i, v := range arr {
        result[i] = v * 2
    }
    return result
}
```

### Variadic Functions with Arrays

Arrays can't be passed to variadic functions directly, but can be unpacked:

```go
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Convert array to variadic arguments
arr := [5]int{1, 2, 3, 4, 5}
total := sum(arr[:]...)  // Unpack slice from array
```

## Array Comparison

Arrays of the same type can be compared with `==` and `!=`.

```go
arr1 := [3]int{1, 2, 3}
arr2 := [3]int{1, 2, 3}
arr3 := [3]int{4, 5, 6}

fmt.Println(arr1 == arr2)  // true
fmt.Println(arr1 == arr3)  // false
fmt.Println(arr1 != arr3)  // true
```

### Arrays of Different Sizes

Cannot compare arrays of different sizes (different types).

```go
arr1 := [3]int{1, 2, 3}
arr2 := [5]int{1, 2, 3, 4, 5}
// fmt.Println(arr1 == arr2)  // Compile error: type mismatch
```

### Arrays with Non-comparable Elements

Arrays containing slices, maps, or functions cannot be compared.

```go
// OK: Arrays of comparable types
arr1 := [2]int{1, 2}
arr2 := [2]int{1, 2}
equal := arr1 == arr2

// ERROR: Arrays of non-comparable types
// arr3 := [2][]int{{1}, {2}}
// arr4 := [2][]int{{1}, {2}}
// equal := arr3 == arr4  // Compile error
```

## Common Patterns

### Fixed-Size Buffer

```go
type RingBuffer struct {
    buffer [10]int
    index  int
}

func (rb *RingBuffer) Add(value int) {
    rb.buffer[rb.index%len(rb.buffer)] = value
    rb.index++
}
```

### Lookup Table

```go
// Days in each month (non-leap year)
var daysInMonth [12]int = [12]int{
    31, 28, 31, 30, 31, 30,
    31, 31, 30, 31, 30, 31,
}

func getDaysInMonth(month int) int {
    if month < 1 || month > 12 {
        return 0
    }
    return daysInMonth[month-1]
}
```

### Character Counting

```go
func countLetters(s string) [26]int {
    var counts [26]int
    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            counts[r-'a']++
        } else if r >= 'A' && r <= 'Z' {
            counts[r-'A']++
        }
    }
    return counts
}

// Usage
text := "Hello World"
counts := countLetters(text)
fmt.Printf("Letter 'l' appears %d times\n", counts['l'-'a'])
```

### Matrix Operations

```go
func addMatrices(a, b [3][3]int) [3][3]int {
    var result [3][3]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            result[i][j] = a[i][j] + b[i][j]
        }
    }
    return result
}

func multiplyMatrices(a, b [3][3]int) [3][3]int {
    var result [3][3]int
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            for k := 0; k < 3; k++ {
                result[i][j] += a[i][k] * b[k][j]
            }
        }
    }
    return result
}
```

## Arrays vs Slices

### Key Differences

| Feature          | Arrays                    | Slices                  |
| ---------------- | ------------------------- | ----------------------- |
| Size             | Fixed at compile time     | Dynamic                 |
| Type             | Size is part of type      | Size not part of type   |
| Memory           | Value type (copied)       | Reference type (shared) |
| Zero value       | Array of zero values      | `nil`                   |
| Comparison       | Can be compared with `==` | Cannot be compared      |
| Function passing | Copied (expensive)        | Reference (cheap)       |

### When to Use Arrays

```go
// Use arrays when:

// 1. Size is known and fixed
var ipv4Addr [4]byte = [4]byte{192, 168, 1, 1}

// 2. You want value semantics
type Vec3 [3]float64

// 3. Small, fixed-size collections
var rgba [4]uint8  // RGBA color

// 4. As backing storage for slices
var buffer [1024]byte
slice := buffer[:]
```

### When to Use Slices

```go
// Use slices when:

// 1. Size is unknown or variable
var items []Item

// 2. You need to append elements
items = append(items, newItem)

// 3. Passing to/from functions
func process(data []int) {}

// 4. Most cases! Slices are more flexible
```

## Best Practices

### 1. Prefer Slices Over Arrays

```go
// Less common: Array
func processArray(arr [100]int) {}

// Preferred: Slice
func processSlice(s []int) {}
```

### 2. Use Arrays for Fixed-Size Data

```go
// Good: Array for fixed-size
type IPv4 [4]byte
type UUID [16]byte
type SHA256 [32]byte

// Good: Known fixed size at compile time
var weekdays [7]string = [7]string{
    "Monday", "Tuesday", "Wednesday", "Thursday",
    "Friday", "Saturday", "Sunday",
}
```

### 3. Pass Array Pointers to Functions

```go
// Bad: Copies entire array
func process(arr [1000]int) {}

// Good: Passes pointer (8 bytes)
func process(arr *[1000]int) {}

// Better: Use slice
func process(arr []int) {}
```

### 4. Use [...] for Compile-Time Sizing

```go
// Good: Compiler counts elements
primes := [...]int{2, 3, 5, 7, 11, 13}

// Less maintainable: Manual count
primes := [6]int{2, 3, 5, 7, 11, 13}
```

### 5. Initialize Arrays Clearly

```go
// Good: Clear initialization
config := [3]string{
    "host=localhost",
    "port=8080",
    "debug=true",
}

// Good: Specific index initialization
flags := [10]bool{
    0: true,   // First flag enabled
    9: true,   // Last flag enabled
}
```

### 6. Document Array Sizes

```go
// Good: Constants for array sizes
const (
    MaxConnections = 100
    BufferSize     = 4096
)

var connections [MaxConnections]*Connection
var buffer [BufferSize]byte
```

### 7. Use Range for Iteration

```go
// Preferred: Safe, idiomatic
for i, v := range arr {
    // Process i and v
}

// Avoid: Easy to make off-by-one errors
for i := 0; i < len(arr); i++ {
    v := arr[i]
    // Process i and v
}
```

### 8. Be Careful with Multi-dimensional Arrays

```go
// Clear: Formatted initialization
matrix := [3][3]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// Consider: Using slices for large/dynamic matrices
matrix := make([][]int, rows)
for i := range matrix {
    matrix[i] = make([]int, cols)
}
```

## Summary

- Arrays have **fixed size** determined at compile time
- Arrays are **value types** - copied on assignment/passing
- Array **size is part of the type**: `[5]int` ≠ `[10]int`
- Arrays can be compared with `==` if element type is comparable
- Use **range loops** for safe iteration
- **Pass pointers** to avoid copying large arrays
- **Prefer slices** for most use cases - arrays are for specific scenarios
- Multi-dimensional arrays are possible but slices are more flexible
- Arrays are great for fixed-size data like IP addresses, colors, small buffers

# Integers in Go

Integers are whole numbers without decimal points. Go provides multiple integer types with different sizes and sign characteristics.

## Table of Contents

- [Integer Types](#integer-types)
- [Special Integer Types](#special-integer-types)
- [Declaration and Initialization](#declaration-and-initialization)
- [Arithmetic Operations](#arithmetic-operations)
- [Bitwise Operations](#bitwise-operations)
- [Comparison Operations](#comparison-operations)
- [Type Conversion](#type-conversion)
- [Common Functions](#common-functions)
- [Overflow and Underflow](#overflow-and-underflow)
- [Best Practices](#best-practices)

## Integer Types

Go provides both signed and unsigned integer types:

### Signed Integers

Signed integers can represent both positive and negative numbers.

| Type | Size | Range |
|------|------|-------|
| `int8` | 8 bits | -128 to 127 |
| `int16` | 16 bits | -32,768 to 32,767 |
| `int32` | 32 bits | -2,147,483,648 to 2,147,483,647 |
| `int64` | 64 bits | -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807 |
| `int` | Platform dependent | 32 or 64 bits (depends on system architecture) |

### Unsigned Integers

Unsigned integers can only represent non-negative numbers (0 and positive).

| Type | Size | Range |
|------|------|-------|
| `uint8` | 8 bits | 0 to 255 |
| `uint16` | 16 bits | 0 to 65,535 |
| `uint32` | 32 bits | 0 to 4,294,967,295 |
| `uint64` | 64 bits | 0 to 18,446,744,073,709,551,615 |
| `uint` | Platform dependent | 32 or 64 bits (depends on system architecture) |

## Special Integer Types

### byte

`byte` is an alias for `uint8`. It's commonly used for raw binary data.

```go
var b byte = 255
var data []byte = []byte("Hello")
```

### rune

`rune` is an alias for `int32`. It represents a Unicode code point.

```go
var r rune = 'A'  // Unicode code point 65
var emoji rune = '😀'  // Unicode code point 128512
```

### uintptr

`uintptr` is an unsigned integer large enough to store a pointer value. Used in low-level programming.

```go
var ptr uintptr
```

## Declaration and Initialization

### Zero Value

Uninitialized integers have a zero value of `0`.

```go
var x int        // x = 0
var y int32      // y = 0
var z uint       // z = 0
```

### Explicit Declaration

```go
var age int = 25
var count int32 = 1000
var size uint64 = 1024
```

### Short Declaration

```go
x := 42          // Type inferred as int
y := int64(100)  // Explicitly int64
```

### Multiple Declaration

```go
var a, b, c int = 1, 2, 3
x, y, z := 10, 20, 30
```

### Using Different Number Bases

```go
decimal := 42        // Base 10
binary := 0b101010   // Base 2 (Binary) = 42
octal := 0o52        // Base 8 (Octal) = 42
hex := 0x2A          // Base 16 (Hexadecimal) = 42
```

## Arithmetic Operations

### Basic Arithmetic

```go
a := 10
b := 3

sum := a + b        // 13 (Addition)
diff := a - b       // 7 (Subtraction)
product := a * b    // 30 (Multiplication)
quotient := a / b   // 3 (Integer division)
remainder := a % b  // 1 (Modulus)
```

### Increment and Decrement

```go
x := 5
x++  // x = 6 (Increment)
x--  // x = 5 (Decrement)

// Note: ++x and --x are NOT valid in Go
```

### Compound Assignment

```go
x := 10
x += 5   // x = 15 (x = x + 5)
x -= 3   // x = 12 (x = x - 3)
x *= 2   // x = 24 (x = x * 2)
x /= 4   // x = 6 (x = x / 4)
x %= 4   // x = 2 (x = x % 4)
```

### Unary Operations

```go
x := 10
y := -x   // y = -10 (Negation)
z := +x   // z = 10 (Unary plus, rarely used)
```

## Bitwise Operations

Bitwise operations work on individual bits of integers.

### Bitwise Operators

```go
a := 12  // 1100 in binary
b := 10  // 1010 in binary

// AND: 1100 & 1010 = 1000 = 8
and := a & b  // 8

// OR: 1100 | 1010 = 1110 = 14
or := a | b   // 14

// XOR: 1100 ^ 1010 = 0110 = 6
xor := a ^ b  // 6

// AND NOT: 1100 &^ 1010 = 0100 = 4
andNot := a &^ b  // 4

// NOT (complement)
not := ^a  // -13 (inverts all bits)
```

### Bit Shift Operations

```go
x := 8  // 1000 in binary

// Left shift (multiply by 2^n)
leftShift := x << 2   // 32 (1000 << 2 = 100000)

// Right shift (divide by 2^n)
rightShift := x >> 1  // 4 (1000 >> 1 = 0100)
```

### Practical Bitwise Examples

```go
// Check if a number is even
isEven := (x & 1) == 0

// Check if a number is odd
isOdd := (x & 1) == 1

// Multiply by 2
doubled := x << 1

// Divide by 2
halved := x >> 1

// Get nth bit (0-indexed from right)
nthBit := (x >> n) & 1

// Set nth bit to 1
setBit := x | (1 << n)

// Clear nth bit (set to 0)
clearBit := x & ^(1 << n)

// Toggle nth bit
toggleBit := x ^ (1 << n)
```

## Comparison Operations

```go
a := 10
b := 20

equal := a == b      // false
notEqual := a != b   // true
less := a < b        // true
greater := a > b     // false
lessEq := a <= b     // true
greaterEq := a >= b  // false
```

## Type Conversion

Go requires explicit type conversion between different integer types.

### Converting Between Integer Types

```go
var x int32 = 100
var y int64 = int64(x)  // Convert int32 to int64
var z int = int(x)      // Convert int32 to int

// Narrowing conversion (may lose data)
var big int64 = 1000000
var small int16 = int16(big)  // Truncates if value is too large
```

### Converting to/from Strings

```go
import (
    "fmt"
    "strconv"
)

// Integer to string
num := 42
str1 := strconv.Itoa(num)              // "42"
str2 := strconv.FormatInt(int64(num), 10)  // "42" (base 10)
str3 := fmt.Sprintf("%d", num)         // "42"

// String to integer
str := "123"
num1, err := strconv.Atoi(str)         // 123
num2, err := strconv.ParseInt(str, 10, 64)  // 123 (base 10, 64-bit)
```

## Common Functions

Go's `math` package provides useful functions for integers.

```go
import "math"

// Absolute value
abs := int(math.Abs(float64(-42)))  // 42

// Maximum and minimum
max := math.Max(float64(a), float64(b))
min := math.Min(float64(a), float64(b))

// Power
power := int(math.Pow(2, 10))  // 1024 (2^10)
```

### Custom Helper Functions

```go
// Max for integers
func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// Min for integers
func Min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// Abs for integers
func Abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
```

## Overflow and Underflow

Integer overflow/underflow occurs when a value exceeds the type's range.

```go
var x int8 = 127
x++  // x = -128 (overflow wraps around)

var y uint8 = 0
y--  // y = 255 (underflow wraps around)
```

### Checking for Overflow

```go
import "math"

func AddWithOverflowCheck(a, b int) (int, bool) {
    if b > 0 && a > math.MaxInt-b {
        return 0, true  // Overflow would occur
    }
    if b < 0 && a < math.MinInt-b {
        return 0, true  // Underflow would occur
    }
    return a + b, false
}
```

## Best Practices

1. **Use `int` by default** - It's optimized for the platform and sufficient for most cases

```go
// Preferred
count := 42

// Only use sized types when necessary
var smallCount int8 = 10  // If you specifically need 8-bit storage
```

2. **Be careful with unsigned integers** - They can cause subtle bugs

```go
// Bad: Can cause underflow
var x uint = 5
var y uint = 10
diff := x - y  // Wraps to large positive number!

// Good: Use signed integers for arithmetic
var x int = 5
var y int = 10
diff := x - y  // -5 (correct)
```

3. **Always check for overflow** in critical applications

```go
result, overflow := AddWithOverflowCheck(a, b)
if overflow {
    // Handle error
}
```

4. **Use appropriate types for specific domains**

```go
// For byte data
var data []byte = []byte("Hello")

// For Unicode characters
var char rune = 'A'

// For file sizes (always positive)
var fileSize uint64 = 1024 * 1024
```

5. **Be explicit with type conversions**

```go
// Bad: Implicit conversion not allowed
var x int32 = 100
var y int64 = x  // Compile error

// Good: Explicit conversion
var x int32 = 100
var y int64 = int64(x)
```

6. **Use constants for fixed values**

```go
const (
    MaxRetries = 3
    BufferSize = 1024
    MinAge     = 18
)
```

## Summary

- Go provides multiple integer types with different sizes and sign characteristics
- Use `int` for general-purpose integers
- Use unsigned types (`uint`) only when you need non-negative values
- Be aware of overflow and underflow behavior
- Always use explicit type conversion
- Bitwise operations are powerful for low-level programming
- Use `byte` for binary data and `rune` for Unicode characters

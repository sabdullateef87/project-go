# Floating-Point Numbers in Go

Floating-point numbers represent decimal values in Go. They're essential for scientific calculations, financial computations, and any application requiring fractional precision.

## Table of Contents

- [Floating-Point Types](#floating-point-types)
- [Declaration and Initialization](#declaration-and-initialization)
- [Arithmetic Operations](#arithmetic-operations)
- [Comparison Operations](#comparison-operations)
- [Special Values](#special-values)
- [Precision and Rounding](#precision-and-rounding)
- [Type Conversion](#type-conversion)
- [Math Package Functions](#math-package-functions)
- [Formatting and Parsing](#formatting-and-parsing)
- [Common Pitfalls](#common-pitfalls)
- [Best Practices](#best-practices)

## Floating-Point Types

Go provides two floating-point types:

| Type      | Size    | Precision          | Range                      |
| --------- | ------- | ------------------ | -------------------------- |
| `float32` | 32 bits | ~7 decimal digits  | ±1.18×10⁻³⁸ to ±3.4×10³⁸   |
| `float64` | 64 bits | ~15 decimal digits | ±2.23×10⁻³⁰⁸ to ±1.8×10³⁰⁸ |

**Note:** `float64` is the default and recommended type for most floating-point operations.

## Declaration and Initialization

### Zero Value

The zero value for floating-point types is `0.0`.

```go
var x float32  // x = 0.0
var y float64  // y = 0.0
```

### Explicit Declaration

```go
var temperature float64 = 98.6
var price float32 = 19.99
var pi float64 = 3.14159265359
```

### Short Declaration

```go
height := 5.9         // Type inferred as float64
width := float32(10.5)  // Explicitly float32
```

### Scientific Notation

```go
avogadro := 6.02214076e23     // 6.02214076 × 10²³
electron := 9.10938356e-31    // 9.10938356 × 10⁻³¹
planck := 6.62607015E-34      // 6.62607015 × 10⁻³⁴
```

### Multiple Declaration

```go
var length, width, height float64 = 10.5, 20.3, 15.7
x, y, z := 1.1, 2.2, 3.3
```

## Arithmetic Operations

### Basic Arithmetic

```go
a := 10.5
b := 3.2

sum := a + b         // 13.7 (Addition)
diff := a - b        // 7.3 (Subtraction)
product := a * b     // 33.6 (Multiplication)
quotient := a / b    // 3.28125 (Division)

// Note: Modulus (%) is NOT available for floating-point numbers
```

### Compound Assignment

```go
x := 10.5
x += 5.2   // x = 15.7
x -= 3.1   // x = 12.6
x *= 2.0   // x = 25.2
x /= 4.0   // x = 6.3
```

### Unary Operations

```go
x := 10.5
y := -x   // y = -10.5 (Negation)
z := +x   // z = 10.5 (Unary plus)
```

### Mixed Integer and Float Operations

```go
// Must convert explicitly
intVal := 10
floatVal := 3.5

// Bad: Type mismatch error
// result := intVal + floatVal

// Good: Explicit conversion
result := float64(intVal) + floatVal  // 13.5
```

## Comparison Operations

### Standard Comparisons

```go
a := 10.5
b := 20.3

equal := a == b      // false
notEqual := a != b   // true
less := a < b        // true
greater := a > b     // false
lessEq := a <= b     // true
greaterEq := a >= b  // false
```

### Floating-Point Comparison Pitfall

Due to precision issues, direct equality comparison can be unreliable.

```go
// Problematic comparison
a := 0.1 + 0.2
b := 0.3
result := (a == b)  // May be false due to floating-point precision!

// Better approach: Use epsilon comparison
func almostEqual(a, b, epsilon float64) bool {
    return math.Abs(a - b) < epsilon
}

const epsilon = 1e-9
result := almostEqual(a, b, epsilon)  // More reliable
```

## Special Values

Go supports special floating-point values from IEEE 754 standard.

### Infinity

```go
import "math"

positiveInf := math.Inf(1)   // +Inf
negativeInf := math.Inf(-1)  // -Inf

// Check for infinity
isInf := math.IsInf(positiveInf, 0)   // true (any infinity)
isPosInf := math.IsInf(positiveInf, 1)  // true (positive only)
isNegInf := math.IsInf(negativeInf, -1) // true (negative only)

// Operations with infinity
result := 1.0 / 0.0  // +Inf (division by zero)
```

### NaN (Not a Number)

```go
import "math"

nan := math.NaN()

// Check for NaN
isNaN := math.IsNaN(nan)  // true

// NaN properties
result1 := nan == nan  // false (NaN is never equal to anything, even itself!)
result2 := math.IsNaN(0.0 / 0.0)  // true
result3 := math.IsNaN(math.Sqrt(-1))  // true
```

### Checking for Special Values

```go
func isNormal(f float64) bool {
    return !math.IsNaN(f) && !math.IsInf(f, 0)
}

// Usage
x := 42.5
if isNormal(x) {
    // Safe to use x in calculations
}
```

## Precision and Rounding

### Understanding Precision

```go
var f32 float32 = 1.23456789
var f64 float64 = 1.23456789012345678

fmt.Printf("float32: %.10f\n", f32)  // 1.2345678806 (precision lost)
fmt.Printf("float64: %.20f\n", f64)  // 1.23456789012345678000
```

### Rounding Functions

```go
import "math"

x := 3.7
y := -2.8

// Ceiling (round up)
ceil := math.Ceil(x)   // 4.0
ceilNeg := math.Ceil(y)  // -2.0

// Floor (round down)
floor := math.Floor(x)   // 3.0
floorNeg := math.Floor(y)  // -3.0

// Truncate (remove decimal part)
trunc := math.Trunc(x)   // 3.0
truncNeg := math.Trunc(y)  // -2.0

// Round to nearest integer
round := math.Round(x)   // 4.0
roundNeg := math.Round(y)  // -3.0

// Round to nearest even (banker's rounding)
roundEven := math.RoundToEven(2.5)  // 2.0
roundEven2 := math.RoundToEven(3.5)  // 4.0
```

### Custom Rounding to Decimal Places

```go
// Round to n decimal places
func roundToDecimal(num float64, precision int) float64 {
    shift := math.Pow(10, float64(precision))
    return math.Round(num * shift) / shift
}

// Usage
value := 3.14159265
rounded := roundToDecimal(value, 2)  // 3.14
```

## Type Conversion

### Between Float Types

```go
var f32 float32 = 3.14
var f64 float64 = float64(f32)  // float32 to float64

var big float64 = 3.141592653589793
var small float32 = float32(big)  // float64 to float32 (may lose precision)
```

### Float to Integer

```go
f := 42.7

// Simple conversion (truncates)
i := int(f)  // 42

// Round before converting
rounded := int(math.Round(f))  // 43
ceiling := int(math.Ceil(f))   // 43
floored := int(math.Floor(f))  // 42
```

### Integer to Float

```go
i := 42
f := float64(i)  // 42.0
```

### String Conversions

```go
import (
    "fmt"
    "strconv"
)

// Float to string
num := 3.14159

str1 := fmt.Sprintf("%f", num)         // "3.141590"
str2 := fmt.Sprintf("%.2f", num)       // "3.14"
str3 := strconv.FormatFloat(num, 'f', 2, 64)  // "3.14"
str4 := strconv.FormatFloat(num, 'e', -1, 64) // "3.14159e+00" (scientific)

// String to float
str := "3.14159"
parsed, err := strconv.ParseFloat(str, 64)
if err != nil {
    // Handle error
}
```

## Math Package Functions

Go's `math` package provides extensive floating-point functions.

### Power and Exponential

```go
import "math"

// Power
pow := math.Pow(2, 10)        // 1024.0 (2¹⁰)
sqrt := math.Sqrt(16)          // 4.0
cbrt := math.Cbrt(27)          // 3.0

// Exponential
exp := math.Exp(1)             // 2.718281828... (e¹)
exp2 := math.Exp2(10)          // 1024.0 (2¹⁰)

// Logarithms
ln := math.Log(math.E)         // 1.0 (natural log)
log10 := math.Log10(100)       // 2.0
log2 := math.Log2(1024)        // 10.0
```

### Trigonometric Functions

```go
import "math"

angle := math.Pi / 4  // 45 degrees in radians

// Basic trigonometric
sin := math.Sin(angle)    // 0.707...
cos := math.Cos(angle)    // 0.707...
tan := math.Tan(angle)    // 1.0

// Inverse trigonometric
asin := math.Asin(0.5)    // 0.523... (30 degrees)
acos := math.Acos(0.5)    // 1.047... (60 degrees)
atan := math.Atan(1.0)    // 0.785... (45 degrees)
atan2 := math.Atan2(1, 1) // 0.785... (angle of point (1,1))

// Hyperbolic
sinh := math.Sinh(1)
cosh := math.Cosh(1)
tanh := math.Tanh(1)
```

### Absolute and Sign

```go
import "math"

// Absolute value
abs := math.Abs(-42.5)     // 42.5

// Copy sign
copysign := math.Copysign(10, -1)  // -10.0 (magnitude of 10, sign of -1)

// Sign bit
signbit := math.Signbit(-42.5)  // true (negative number)
```

### Min, Max, and Modulus

```go
import "math"

// Maximum and minimum
max := math.Max(10.5, 20.3)  // 20.3
min := math.Min(10.5, 20.3)  // 10.5

// Floating-point remainder (modulus)
mod := math.Mod(10.5, 3.0)   // 1.5
```

### Constants

```go
import "math"

pi := math.Pi              // 3.141592653589793
e := math.E                // 2.718281828459045
phi := math.Phi            // 1.618033988749895 (golden ratio)
sqrt2 := math.Sqrt2        // 1.414213562373095
ln2 := math.Ln2            // 0.693147180559945
log2e := math.Log2E        // 1.442695040888963

maxFloat64 := math.MaxFloat64  // 1.797693134862315708e+308
minFloat64 := math.SmallestNonzeroFloat64  // 4.940656458412465441e-324
```

## Formatting and Parsing

### Printf Formatting

```go
import "fmt"

num := 3.14159265

// Basic float
fmt.Printf("%f\n", num)      // 3.141593

// With precision
fmt.Printf("%.2f\n", num)    // 3.14
fmt.Printf("%.4f\n", num)    // 3.1416

// Scientific notation
fmt.Printf("%e\n", num)      // 3.141593e+00
fmt.Printf("%.2e\n", num)    // 3.14e+00

// Compact representation
fmt.Printf("%g\n", num)      // 3.14159265

// With width and alignment
fmt.Printf("%10.2f\n", num)  //       3.14
fmt.Printf("%-10.2f\n", num) // 3.14
fmt.Printf("%010.2f\n", num) // 0000003.14
```

### Parsing Strings

```go
import "strconv"

// Parse float64
str := "3.14159"
f64, err := strconv.ParseFloat(str, 64)
if err != nil {
    // Handle error
}

// Parse float32
f32, err := strconv.ParseFloat(str, 32)
if err != nil {
    // Handle error
}
result := float32(f32)

// Handle special values
inf, _ := strconv.ParseFloat("Inf", 64)      // +Inf
negInf, _ := strconv.ParseFloat("-Inf", 64)  // -Inf
nan, _ := strconv.ParseFloat("NaN", 64)      // NaN
```

## Common Pitfalls

### 1. Precision Loss

```go
// Problem
var sum float32 = 0.0
for i := 0; i < 1000000; i++ {
    sum += 0.0001
}
// sum may not equal 100.0 due to accumulated precision errors

// Solution: Use float64 for better precision
var sum float64 = 0.0
for i := 0; i < 1000000; i++ {
    sum += 0.0001
}
```

### 2. Equality Comparison

```go
// Problem
a := 0.1 + 0.2
b := 0.3
if a == b {  // May be false!
    // ...
}

// Solution: Use epsilon comparison
const epsilon = 1e-9
if math.Abs(a - b) < epsilon {
    // ...
}
```

### 3. Division by Zero

```go
// Returns infinity, not a panic
result := 1.0 / 0.0  // +Inf

// Check before division
func safeDivide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}
```

### 4. Integer Division

```go
// Problem: Integer division loses decimal part
result := 5 / 2  // 2 (not 2.5)

// Solution: Convert to float first
result := float64(5) / float64(2)  // 2.5
```

### 5. NaN Propagation

```go
// NaN spreads through calculations
x := math.NaN()
y := x + 10      // NaN
z := y * 2       // NaN

// Always check for NaN in critical calculations
if !math.IsNaN(x) {
    // Safe to use x
}
```

## Best Practices

### 1. Use float64 by Default

```go
// Preferred
distance := 42.5

// Only use float32 when memory is critical
var positions []float32  // For large arrays where precision is acceptable
```

### 2. Avoid Direct Equality Comparisons

```go
// Bad
if x == y {
    // ...
}

// Good
const epsilon = 1e-9
if math.Abs(x - y) < epsilon {
    // ...
}

// Helper function
func almostEqual(a, b, epsilon float64) bool {
    return math.Abs(a-b) < epsilon
}
```

### 3. Use Math Package Constants

```go
// Bad
pi := 3.14

// Good
pi := math.Pi
```

### 4. Handle Special Values

```go
func safeCalculation(x float64) (float64, error) {
    if math.IsNaN(x) {
        return 0, fmt.Errorf("input is NaN")
    }
    if math.IsInf(x, 0) {
        return 0, fmt.Errorf("input is infinite")
    }
    // Proceed with calculation
    return math.Sqrt(x), nil
}
```

### 5. Be Aware of Precision Limits

```go
// For financial calculations, consider using integers with fixed decimal places
// or the decimal package from third-party libraries

// Example: Store cents as integers
priceInCents := 1999  // $19.99
priceInDollars := float64(priceInCents) / 100.0
```

### 6. Round Appropriately for Display

```go
// For monetary values
func formatMoney(amount float64) string {
    rounded := math.Round(amount * 100) / 100
    return fmt.Sprintf("$%.2f", rounded)
}
```

### 7. Use Scientific Notation for Very Large/Small Numbers

```go
// Preferred
avogadro := 6.02214076e23
planck := 6.62607015e-34

// Less readable
avogadro := 602214076000000000000000.0
```

## Summary

- Go provides `float32` and `float64` for decimal numbers
- `float64` is the default and recommended for most use cases
- Floating-point arithmetic has precision limitations
- Never use direct equality comparison for floats; use epsilon comparison
- Be aware of special values: `Inf` and `NaN`
- The `math` package provides comprehensive mathematical functions
- Always check for edge cases (division by zero, NaN, infinity)
- For financial calculations, consider alternatives to floating-point

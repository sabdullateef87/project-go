# Complex Numbers in Go

Complex numbers are numbers with a real and imaginary part, commonly used in scientific computing, signal processing, and engineering applications.

## Table of Contents

- [Complex Number Types](#complex-number-types)
- [Declaration and Initialization](#declaration-and-initialization)
- [Arithmetic Operations](#arithmetic-operations)
- [Accessing Real and Imaginary Parts](#accessing-real-and-imaginary-parts)
- [Complex Number Functions](#complex-number-functions)
- [Type Conversion](#type-conversion)
- [Polar Representation](#polar-representation)
- [Practical Examples](#practical-examples)
- [Best Practices](#best-practices)

## Complex Number Types

Go provides two complex number types:

| Type         | Size     | Components           | Precision                        |
| ------------ | -------- | -------------------- | -------------------------------- |
| `complex64`  | 64 bits  | Two `float32` values | ~7 decimal digits per component  |
| `complex128` | 128 bits | Two `float64` values | ~15 decimal digits per component |

**Mathematical Representation:** A complex number is written as $a + bi$ where:

- $a$ is the real part
- $b$ is the imaginary part
- $i$ is the imaginary unit where $i^2 = -1$

## Declaration and Initialization

### Zero Value

The zero value for complex types is `(0+0i)`.

```go
var c1 complex64   // c1 = (0+0i)
var c2 complex128  // c2 = (0+0i)
```

### Using Complex Literals

```go
// Direct literal notation
c1 := 3 + 4i              // Type inferred as complex128
c2 := 1.5 + 2.7i          // complex128
c3 := -3 - 4i             // complex128

// Explicit type
var c4 complex64 = 5 + 6i
var c5 complex128 = 2.5 + 3.8i
```

### Using the complex() Function

The `complex()` built-in function creates a complex number from real and imaginary parts.

```go
// complex128 (default)
c1 := complex(3, 4)           // (3+4i)
c2 := complex(1.5, 2.7)       // (1.5+2.7i)

// complex64 (explicit)
real := float32(3.0)
imag := float32(4.0)
c3 := complex(real, imag)     // complex64: (3+4i)
```

### Pure Real or Imaginary Numbers

```go
// Pure real number (imaginary part is 0)
realOnly := complex(5, 0)     // (5+0i)
realOnly := 5 + 0i            // Same as above

// Pure imaginary number (real part is 0)
imagOnly := complex(0, 3)     // (0+3i)
imagOnly := 3i                // Same as above
```

### Multiple Declaration

```go
var a, b, c complex128 = 1+2i, 3+4i, 5+6i
x, y, z := 1+1i, 2+2i, 3+3i
```

## Arithmetic Operations

Complex numbers support standard arithmetic operations.

### Addition

Mathematical: $(a + bi) + (c + di) = (a + c) + (b + d)i$

```go
c1 := 3 + 4i
c2 := 1 + 2i
sum := c1 + c2  // (4+6i)
```

### Subtraction

Mathematical: $(a + bi) - (c + di) = (a - c) + (b - d)i$

```go
c1 := 5 + 6i
c2 := 2 + 3i
diff := c1 - c2  // (3+3i)
```

### Multiplication

Mathematical: $(a + bi) \times (c + di) = (ac - bd) + (ad + bc)i$

```go
c1 := 2 + 3i
c2 := 4 + 5i
product := c1 * c2  // (-7+22i)
// Calculation: (2*4 - 3*5) + (2*5 + 3*4)i = -7 + 22i
```

### Division

Mathematical: $\frac{a + bi}{c + di} = \frac{(ac + bd) + (bc - ad)i}{c^2 + d^2}$

```go
c1 := 10 + 5i
c2 := 2 + 1i
quotient := c1 / c2  // (5+0i)
```

### Compound Assignment

```go
c := 3 + 4i
c += 1 + 2i   // c = (4+6i)
c -= 1 + 1i   // c = (3+5i)
c *= 2 + 0i   // c = (6+10i)
c /= 2 + 0i   // c = (3+5i)
```

### Multiplication with Scalars

```go
c := 3 + 4i
scaled := c * 2       // (6+8i) - multiply by real number
scaled := c * (2+0i)  // Same as above
```

### Negation

```go
c := 3 + 4i
neg := -c  // (-3-4i)
```

## Accessing Real and Imaginary Parts

Use the `real()` and `imag()` built-in functions to extract components.

```go
c := 3 + 4i

// Extract real part
realPart := real(c)  // 3.0 (type: float64)

// Extract imaginary part
imagPart := imag(c)  // 4.0 (type: float64)

// For complex64
var c64 complex64 = 5 + 6i
r := real(c64)  // 5.0 (type: float32)
i := imag(c64)  // 6.0 (type: float32)
```

## Complex Number Functions

Go's `math/cmplx` package provides functions for complex numbers.

### Import the Package

```go
import (
    "math"
    "math/cmplx"
)
```

### Magnitude (Absolute Value)

The magnitude of $a + bi$ is $\sqrt{a^2 + b^2}$

```go
c := 3 + 4i
magnitude := cmplx.Abs(c)  // 5.0
// |3+4i| = √(3² + 4²) = √25 = 5
```

### Phase (Argument)

The phase is the angle $\theta$ in polar form: $r e^{i\theta}$

```go
c := 1 + 1i
phase := cmplx.Phase(c)  // 0.7853981633974483 (π/4 radians, 45°)

// Convert to degrees
degrees := phase * 180 / math.Pi  // 45.0
```

### Conjugate

The conjugate of $a + bi$ is $a - bi$

```go
c := 3 + 4i
conj := cmplx.Conj(c)  // (3-4i)
```

### Square Root

```go
c := -1 + 0i
sqrt := cmplx.Sqrt(c)  // (0+1i) - the imaginary unit i
```

### Power

```go
c := 2 + 0i
power := cmplx.Pow(c, 3)  // (8+0i) - 2³ = 8

// Complex exponent
c2 := 1 + 1i
result := cmplx.Pow(c2, 2+0i)  // Square of (1+1i)
```

### Exponential

Euler's formula: $e^{i\theta} = \cos(\theta) + i\sin(\theta)$

```go
c := complex(0, math.Pi)  // iπ
exp := cmplx.Exp(c)       // (-1+0i) - Euler's identity: e^(iπ) = -1
```

### Logarithm

```go
c := 1 + 0i
log := cmplx.Log(c)  // (0+0i) - ln(1) = 0

c2 := 0 + 1i
log2 := cmplx.Log(c2)  // (0+1.5707963267948966i) - ln(i) = iπ/2
```

### Trigonometric Functions

```go
c := 1 + 1i

// Basic trigonometric
sin := cmplx.Sin(c)
cos := cmplx.Cos(c)
tan := cmplx.Tan(c)

// Inverse trigonometric
asin := cmplx.Asin(c)
acos := cmplx.Acos(c)
atan := cmplx.Atan(c)

// Hyperbolic
sinh := cmplx.Sinh(c)
cosh := cmplx.Cosh(c)
tanh := cmplx.Tanh(c)

// Inverse hyperbolic
asinh := cmplx.Asinh(c)
acosh := cmplx.Acosh(c)
atanh := cmplx.Atanh(c)
```

### Special Values and Checks

```go
import "math/cmplx"

// Check for infinity
c := complex(math.Inf(1), 0)
isInf := cmplx.IsInf(c)  // true

// Check for NaN
c2 := complex(math.NaN(), 0)
isNaN := cmplx.IsNaN(c2)  // true

// Special constants
inf := cmplx.Inf()    // Complex infinity
nan := cmplx.NaN()    // Complex NaN
```

## Type Conversion

### Between Complex Types

```go
var c64 complex64 = 3 + 4i
var c128 complex128 = complex128(c64)  // complex64 to complex128

var big complex128 = 5.123456789 + 6.987654321i
var small complex64 = complex64(big)   // complex128 to complex64 (may lose precision)
```

### From Real Numbers

```go
// Float to complex
f := 3.14
c := complex(f, 0)  // (3.14+0i)

// Integer to complex
i := 42
c2 := complex(float64(i), 0)  // (42+0i)
```

### To Real Numbers

```go
c := 3 + 4i

// Extract real part
realPart := real(c)  // 3.0

// Get magnitude (absolute value)
magnitude := cmplx.Abs(c)  // 5.0
```

### String Conversions

```go
import (
    "fmt"
    "strconv"
)

// Complex to string
c := 3 + 4i
str := fmt.Sprintf("%v", c)    // "(3+4i)"
str2 := fmt.Sprintf("%.2f", c) // "(3.00+4.00i)"

// String to complex (requires custom parsing or libraries)
// Go's standard library doesn't have direct string to complex conversion
```

## Polar Representation

Complex numbers can be represented in polar form: $r e^{i\theta}$ where:

- $r$ is the magnitude (modulus)
- $\theta$ is the phase (argument)

### Conversion: Rectangular to Polar

```go
import (
    "math"
    "math/cmplx"
)

c := 3 + 4i

// Get polar coordinates
r := cmplx.Abs(c)      // Magnitude: 5.0
theta := cmplx.Phase(c) // Phase: 0.927 radians (≈53.13°)

// Convert phase to degrees
degrees := theta * 180 / math.Pi  // 53.13°
```

### Conversion: Polar to Rectangular

```go
import (
    "math"
    "math/cmplx"
)

// Given polar coordinates
r := 5.0
theta := math.Pi / 4  // 45 degrees

// Create complex number using Euler's formula
// z = r * e^(iθ) = r * (cos(θ) + i*sin(θ))
c := complex(r*math.Cos(theta), r*math.Sin(theta))
// Result: (3.535+3.535i)

// Alternative using cmplx package
c2 := cmplx.Rect(r, theta)  // Same result
```

### Polar Form Operations

```go
// Multiplication in polar form: multiply magnitudes, add phases
c1 := cmplx.Rect(2, math.Pi/4)   // r=2, θ=45°
c2 := cmplx.Rect(3, math.Pi/6)   // r=3, θ=30°

product := c1 * c2
// Polar: r=6, θ=75° (45° + 30°)
rProduct := cmplx.Abs(product)      // 6.0
thetaProduct := cmplx.Phase(product) // 1.309 rad (75°)

// Division in polar form: divide magnitudes, subtract phases
quotient := c1 / c2
// Polar: r=2/3, θ=15° (45° - 30°)
```

## Practical Examples

### Electrical Engineering: AC Circuit Analysis

```go
// Impedance calculation: Z = R + jωL - j/(ωC)
func calculateImpedance(R, L, C, omega float64) complex128 {
    resistance := complex(R, 0)
    inductance := complex(0, omega*L)
    capacitance := complex(0, -1/(omega*C))
    return resistance + inductance + capacitance
}

// Usage
Z := calculateImpedance(100, 0.1, 0.00001, 377)  // ω = 2πf, f=60Hz
fmt.Printf("Impedance: %.2f Ω\n", Z)
```

### Signal Processing: Fourier Transform Component

```go
import (
    "math"
    "math/cmplx"
)

// Calculate a single DFT coefficient
func dftCoefficient(signal []float64, k int) complex128 {
    N := len(signal)
    sum := complex(0, 0)

    for n := 0; n < N; n++ {
        angle := -2 * math.Pi * float64(k) * float64(n) / float64(N)
        exponential := cmplx.Exp(complex(0, angle))
        sum += complex(signal[n], 0) * exponential
    }

    return sum
}
```

### Quantum Computing: Qubit State

```go
// Qubit state: α|0⟩ + β|1⟩ where |α|² + |β|² = 1
type Qubit struct {
    alpha complex128  // Amplitude for |0⟩
    beta  complex128  // Amplitude for |1⟩
}

// Calculate probability of measuring |0⟩
func (q Qubit) Prob0() float64 {
    return real(q.alpha * cmplx.Conj(q.alpha))
}

// Calculate probability of measuring |1⟩
func (q Qubit) Prob1() float64 {
    return real(q.beta * cmplx.Conj(q.beta))
}

// Hadamard gate
func (q Qubit) Hadamard() Qubit {
    sqrt2 := complex(1/math.Sqrt(2), 0)
    return Qubit{
        alpha: sqrt2 * (q.alpha + q.beta),
        beta:  sqrt2 * (q.alpha - q.beta),
    }
}
```

### Mandelbrot Set

```go
// Check if a complex number is in the Mandelbrot set
func inMandelbrotSet(c complex128, maxIterations int) bool {
    z := complex(0, 0)

    for i := 0; i < maxIterations; i++ {
        z = z*z + c
        if cmplx.Abs(z) > 2 {
            return false
        }
    }

    return true
}

// Usage
c := complex(-0.5, 0.5)
if inMandelbrotSet(c, 100) {
    fmt.Println("Point is in Mandelbrot set")
}
```

### Root Finding

```go
// Find nth roots of a complex number
func nthRoots(c complex128, n int) []complex128 {
    r := cmplx.Abs(c)
    theta := cmplx.Phase(c)

    roots := make([]complex128, n)
    rRoot := math.Pow(r, 1.0/float64(n))

    for k := 0; k < n; k++ {
        angle := (theta + 2*math.Pi*float64(k)) / float64(n)
        roots[k] = cmplx.Rect(rRoot, angle)
    }

    return roots
}

// Find cube roots of -8
roots := nthRoots(-8+0i, 3)
// Results: 1+1.732i, -2+0i, 1-1.732i
```

## Best Practices

### 1. Use complex128 by Default

```go
// Preferred
z := 3 + 4i  // complex128

// Only use complex64 when memory is critical
var smallComplex complex64 = 1 + 2i
```

### 2. Use cmplx Package for Operations

```go
// Don't manually calculate magnitude
c := 3 + 4i
// Bad
magnitude := math.Sqrt(real(c)*real(c) + imag(c)*imag(c))

// Good
magnitude := cmplx.Abs(c)
```

### 3. Check for Special Values

```go
import "math/cmplx"

func safeOperation(c complex128) (complex128, error) {
    if cmplx.IsNaN(c) {
        return 0, fmt.Errorf("input is NaN")
    }
    if cmplx.IsInf(c) {
        return 0, fmt.Errorf("input is infinite")
    }
    // Perform operation
    return cmplx.Sqrt(c), nil
}
```

### 4. Use Polar Form for Multiplication/Division

```go
// When multiplying/dividing many complex numbers,
// polar form can be more efficient and numerically stable

c1 := 2 + 2i
c2 := 3 + 4i

// Convert to polar
r1, theta1 := cmplx.Polar(c1)
r2, theta2 := cmplx.Polar(c2)

// Multiply in polar form
rProduct := r1 * r2
thetaProduct := theta1 + theta2

// Convert back to rectangular
product := cmplx.Rect(rProduct, thetaProduct)
```

### 5. Document Physical Meanings

```go
// Good: Clear physical interpretation
type Impedance complex128  // Electrical impedance in ohms

func (z Impedance) Resistance() float64 {
    return real(complex128(z))
}

func (z Impedance) Reactance() float64 {
    return imag(complex128(z))
}
```

### 6. Be Careful with Equality Comparisons

```go
// Due to floating-point precision, direct comparison may fail
c1 := cmplx.Sqrt(-1 + 0i)
c2 := 0 + 1i

// Use epsilon comparison for real and imaginary parts separately
func almostEqualComplex(a, b complex128, epsilon float64) bool {
    return math.Abs(real(a)-real(b)) < epsilon &&
           math.Abs(imag(a)-imag(b)) < epsilon
}

const epsilon = 1e-9
if almostEqualComplex(c1, c2, epsilon) {
    // ...
}
```

### 7. Format Output Appropriately

```go
c := 3.14159 + 2.71828i

// Default format
fmt.Printf("%v\n", c)          // (3.14159+2.71828i)

// With precision
fmt.Printf("%.2f\n", c)        // (3.14+2.72i)

// Custom formatting
fmt.Printf("%.2f + %.2fi\n", real(c), imag(c))  // 3.14 + 2.72i
```

## Summary

- Go provides `complex64` and `complex128` for complex number arithmetic
- Complex numbers have a real part and an imaginary part
- Use the `complex()` function to create complex numbers from components
- Use `real()` and `imag()` to extract components
- The `math/cmplx` package provides comprehensive complex number functions
- Complex numbers can be represented in rectangular $(a + bi)$ or polar $(re^{i\theta})$ form
- Common applications include signal processing, electrical engineering, and quantum computing
- Always use `complex128` unless memory constraints require `complex64`

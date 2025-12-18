# Basic Data Types in Go

Go is a statically typed language with a rich set of built-in data types. This guide covers all the fundamental data types and their operations.

## Table of Contents

1. [Integers](integers.md) - Signed and unsigned integer types
2. [Floating-Point Numbers](floating-point.md) - Float32 and Float64
3. [Complex Numbers](complex-numbers.md) - Complex64 and Complex128
4. [Booleans](booleans.md) - Boolean type and logical operations
5. [Strings](strings.md) - String type and manipulations
6. [Constants](constants.md) - Constant declarations and iota

## Overview

### Numeric Types

Go provides several numeric types with different sizes and characteristics:

| Category | Types | Description |
|----------|-------|-------------|
| **Integers** | `int8`, `int16`, `int32`, `int64`, `int` | Signed integers |
| | `uint8`, `uint16`, `uint32`, `uint64`, `uint` | Unsigned integers |
| | `byte` (alias for `uint8`) | Byte values |
| | `rune` (alias for `int32`) | Unicode code points |
| **Floating-Point** | `float32`, `float64` | Decimal numbers |
| **Complex** | `complex64`, `complex128` | Complex numbers |

### Non-Numeric Types

| Type | Description |
|------|-------------|
| **Boolean** | `bool` - `true` or `false` |
| **String** | `string` - UTF-8 encoded text |

## Quick Reference

### Zero Values

Every type in Go has a zero value (the default value when declared without initialization):

```go
var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // "" (empty string)
var c complex128 // (0+0i)
```

### Type Conversion

Go requires explicit type conversion between different types:

```go
var i int = 42
var f float64 = float64(i)  // Convert int to float64
var u uint = uint(f)         // Convert float64 to uint
```

### Common Operations

All numeric types support standard arithmetic operations:

- Addition: `+`
- Subtraction: `-`
- Multiplication: `*`
- Division: `/`
- Modulus: `%` (integers only)

Comparison operators work across all types:

- Equal: `==`
- Not equal: `!=`
- Less than: `<`
- Greater than: `>`
- Less than or equal: `<=`
- Greater than or equal: `>=`

## Best Practices

1. **Use `int` for integers** unless you have a specific reason to use a sized type
2. **Use `float64`** as the default floating-point type (more precise than `float32`)
3. **Be explicit with type conversions** - Go doesn't allow implicit conversions
4. **Use constants** for values that don't change
5. **Choose the right type** for your use case to avoid unnecessary conversions

## Learning Path

For a comprehensive understanding of Go's basic data types, explore each topic in order:

1. Start with [Integers](integers.md) to understand whole numbers and their operations
2. Move to [Floating-Point Numbers](floating-point.md) for decimal arithmetic
3. Learn about [Complex Numbers](complex-numbers.md) if working with advanced mathematics
4. Understand [Booleans](booleans.md) for logical operations and control flow
5. Master [Strings](strings.md) for text manipulation
6. Finally, study [Constants](constants.md) for declaring immutable values

## Additional Resources

- [Go Language Specification - Types](https://go.dev/ref/spec#Types)
- [Effective Go - Data](https://go.dev/doc/effective_go#data)
- [Go by Example - Values](https://gobyexample.com/values)

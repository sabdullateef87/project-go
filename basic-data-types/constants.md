# Constants in Go

Constants are immutable values that are known at compile time and do not change during program execution. They're essential for defining fixed values and improving code readability.

## Table of Contents

- [What are Constants?](#what-are-constants)
- [Declaring Constants](#declaring-constants)
- [Typed vs Untyped Constants](#typed-vs-untyped-constants)
- [Constant Expressions](#constant-expressions)
- [The iota Identifier](#the-iota-identifier)
- [Constant Types](#constant-types)
- [Enumerated Constants](#enumerated-constants)
- [Best Practices](#best-practices)

## What are Constants?

Constants are values that:

- **Cannot be changed** after declaration
- Are **evaluated at compile time**
- Can be **numbers, characters, strings, or booleans**
- Do not occupy memory at runtime (replaced by their values during compilation)

### Constants vs Variables

| Feature     | Constants            | Variables        |
| ----------- | -------------------- | ---------------- |
| Mutability  | Immutable            | Mutable          |
| Evaluation  | Compile-time         | Runtime          |
| Storage     | No memory allocation | Memory allocated |
| Declaration | `const` keyword      | `var` keyword    |
| Type        | Can be untyped       | Always typed     |

## Declaring Constants

### Single Constant

```go
const Pi = 3.14159
const MaxUsers = 100
const AppName = "MyApp"
const IsDebug = false
```

### Multiple Constants

```go
// Method 1: Multiple lines
const (
    StatusOK       = 200
    StatusNotFound = 404
    StatusError    = 500
)

// Method 2: Single line
const Width, Height = 800, 600
```

### With Type Annotation

```go
const MaxRetries int = 3
const Timeout time.Duration = 30 * time.Second
const Rate float64 = 0.05
```

### Implicit Type

```go
// Type inferred from value
const Port = 8080              // int
const Message = "Hello"        // string
const Enabled = true           // bool
const Pi = 3.14159            // float64
```

## Typed vs Untyped Constants

### Untyped Constants

Untyped constants have a default type but can be used in any context where that type is compatible.

```go
const (
    // Untyped constants
    UntypedInt    = 42        // Default: int
    UntypedFloat  = 3.14      // Default: float64
    UntypedString = "hello"   // Default: string
    UntypedBool   = true      // Default: bool
)

// Can be used with different types
var i8 int8 = UntypedInt      // OK
var i64 int64 = UntypedInt    // OK
var f32 float32 = UntypedFloat // OK
var f64 float64 = UntypedFloat // OK
```

### Typed Constants

Typed constants have an explicit type and follow the same type rules as variables.

```go
const (
    // Typed constants
    TypedInt8  int8    = 42
    TypedInt64 int64   = 42
    TypedFloat float32 = 3.14
)

// Type must match exactly
var i int64 = TypedInt64         // OK
var j int = TypedInt64           // Error: cannot use int64 as int
var k int = int(TypedInt64)      // OK: explicit conversion
```

### Precision Benefits of Untyped Constants

Untyped constants can maintain arbitrary precision until assigned.

```go
const (
    // Untyped: maintains full precision
    UntypedPi = 3.14159265358979323846

    // Typed: limited by float64 precision
    TypedPi float64 = 3.14159265358979323846
)

// Untyped constants can be used in high-precision calculations
result := UntypedPi * UntypedPi * 100  // Full precision maintained

// Assigned to variable with appropriate precision
var precise float64 = UntypedPi
var less float32 = UntypedPi  // Converted to float32 precision
```

## Constant Expressions

Constants can be created using expressions with other constants and literals.

### Arithmetic Operations

```go
const (
    A = 10
    B = 20
    Sum = A + B        // 30
    Product = A * B    // 200
    Quotient = B / A   // 2
)
```

### Complex Expressions

```go
const (
    SecondsPerMinute = 60
    MinutesPerHour   = 60
    HoursPerDay      = 24

    SecondsPerHour = SecondsPerMinute * MinutesPerHour  // 3600
    SecondsPerDay  = SecondsPerHour * HoursPerDay       // 86400
)
```

### String Operations

```go
const (
    FirstName = "John"
    LastName  = "Doe"
    // String concatenation is allowed
    FullName = FirstName + " " + LastName  // "John Doe"
)
```

### Boolean Operations

```go
const (
    IsProduction = false
    IsDebug      = !IsProduction           // true
    EnableLogs   = IsDebug || IsProduction // true
)
```

### Built-in Functions

Only certain built-in functions can be used in constant expressions:

```go
const (
    Text = "Hello, World!"

    // len() - string and array length
    TextLength = len(Text)              // 13

    // cap() - array capacity
    ArrayCap = cap([5]int{})            // 5

    // real() and imag() - complex number parts
    ComplexReal = real(2 + 3i)          // 2
    ComplexImag = imag(2 + 3i)          // 3

    // complex() - create complex number
    ComplexNum = complex(1, 2)          // (1+2i)
)
```

### Limitations

Some operations are NOT allowed in constant expressions:

```go
// INVALID: These require runtime evaluation
// const Time = time.Now()               // Error
// const Random = rand.Intn(100)         // Error
// const UserInput = getUserInput()      // Error
// const SliceLen = len(mySlice)         // Error (slice length is runtime)
```

## The iota Identifier

`iota` is a special identifier used to create enumerated constants with incrementing values.

### Basic iota Usage

`iota` starts at 0 and increments by 1 for each constant in a `const` block.

```go
const (
    Sunday    = iota  // 0
    Monday            // 1
    Tuesday           // 2
    Wednesday         // 3
    Thursday          // 4
    Friday            // 5
    Saturday          // 6
)
```

### Starting from a Different Value

```go
const (
    January = iota + 1  // 1
    February            // 2
    March               // 3
    April               // 4
    May                 // 5
    June                // 6
    July                // 7
    August              // 8
    September           // 9
    October             // 10
    November            // 11
    December            // 12
)
```

### Skipping Values

```go
const (
    _           = iota  // 0 (skip using blank identifier)
    KB = 1 << (10 * iota) // 1 << 10 = 1024
    MB                     // 1 << 20 = 1048576
    GB                     // 1 << 30 = 1073741824
    TB                     // 1 << 40 = 1099511627776
)
```

### Custom Expressions with iota

```go
const (
    _  = iota             // 0 (skip)
    B  = 1 << (10 * iota) // 1 << 10 = 1024
    KB                    // 1 << 20
    MB                    // 1 << 30
    GB                    // 1 << 40
)

const (
    // Powers of 2
    Bit0 = 1 << iota  // 1 << 0 = 1
    Bit1              // 1 << 1 = 2
    Bit2              // 1 << 2 = 4
    Bit3              // 1 << 3 = 8
    Bit4              // 1 << 4 = 16
)

const (
    // File permissions (Unix style)
    Execute = 1 << iota  // 1
    Write                // 2
    Read                 // 4
)
```

### Multiple iota in Same Block

```go
const (
    A, B = iota, iota + 10  // 0, 10
    C, D                     // 1, 11
    E, F                     // 2, 12
)
```

### Resetting iota

Each new `const` block resets `iota` to 0.

```go
const (
    A = iota  // 0
    B         // 1
    C         // 2
)

const (
    X = iota  // 0 (reset)
    Y         // 1
    Z         // 2
)
```

### Real-World iota Examples

#### HTTP Status Codes

```go
const (
    StatusContinue           = 100
    StatusSwitchingProtocols = 101

    StatusOK                   = 200
    StatusCreated              = 201
    StatusAccepted             = 202

    StatusBadRequest           = 400
    StatusUnauthorized         = 401
    StatusForbidden            = 403
    StatusNotFound             = 404

    StatusInternalServerError  = 500
    StatusNotImplemented       = 501
    StatusBadGateway           = 502
)
```

#### Application States

```go
type State int

const (
    Stopped State = iota  // 0
    Starting              // 1
    Running               // 2
    Stopping              // 3
)

func (s State) String() string {
    switch s {
    case Stopped:
        return "Stopped"
    case Starting:
        return "Starting"
    case Running:
        return "Running"
    case Stopping:
        return "Stopping"
    default:
        return "Unknown"
    }
}
```

#### Bit Flags

```go
type Permission uint

const (
    Read Permission = 1 << iota  // 1 (001)
    Write                         // 2 (010)
    Execute                       // 4 (100)
)

// Combine permissions with bitwise OR
func (p Permission) Has(flag Permission) bool {
    return p&flag != 0
}

// Usage
var perms Permission = Read | Write  // 3 (011)
canRead := perms.Has(Read)           // true
canExecute := perms.Has(Execute)     // false
```

## Constant Types

### Numeric Constants

```go
const (
    // Integers
    MaxInt     = 1000
    MinValue   = -100
    HexValue   = 0xFF    // 255
    BinaryValue = 0b1010 // 10

    // Floats
    Pi         = 3.14159
    E          = 2.71828
    GoldenRatio = 1.618

    // Complex
    ComplexNum = 3 + 4i
)
```

### String Constants

```go
const (
    AppName    = "MyApp"
    Version    = "1.0.0"
    Author     = "John Doe"

    // Multi-line string
    HelpText = `
        Usage: myapp [options]
        Options:
            -h    Show help
            -v    Show version
    `
)
```

### Boolean Constants

```go
const (
    IsProduction = false
    EnableDebug  = true
    UseCache     = true
    AllowGuests  = false
)
```

### Character Constants (Runes)

```go
const (
    Newline   = '\n'
    Tab       = '\t'
    Backslash = '\\'
    LetterA   = 'A'
    Emoji     = '😀'
)
```

## Enumerated Constants

Enumerations are a common use case for constants with iota.

### Simple Enum

```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func (w Weekday) String() string {
    names := [...]string{
        "Sunday", "Monday", "Tuesday", "Wednesday",
        "Thursday", "Friday", "Saturday",
    }
    if w < Sunday || w > Saturday {
        return "Unknown"
    }
    return names[w]
}
```

### Enum with Validation

```go
type Priority int

const (
    Low Priority = iota
    Medium
    High
    Critical
)

func (p Priority) IsValid() bool {
    return p >= Low && p <= Critical
}

func (p Priority) String() string {
    switch p {
    case Low:
        return "Low"
    case Medium:
        return "Medium"
    case High:
        return "High"
    case Critical:
        return "Critical"
    default:
        return "Invalid"
    }
}
```

### Bit Flag Enum

```go
type Feature uint32

const (
    FeatureA Feature = 1 << iota  // 1
    FeatureB                       // 2
    FeatureC                       // 4
    FeatureD                       // 8
)

// Check if feature is enabled
func (f Feature) Has(flag Feature) bool {
    return f&flag != 0
}

// Enable a feature
func (f Feature) With(flag Feature) Feature {
    return f | flag
}

// Disable a feature
func (f Feature) Without(flag Feature) Feature {
    return f &^ flag
}

// Usage
var features Feature
features = features.With(FeatureA).With(FeatureC)  // Enable A and C
hasA := features.Has(FeatureA)                      // true
hasB := features.Has(FeatureB)                      // false
```

## Best Practices

### 1. Use Constants for Fixed Values

```go
// Good: Clear intent, easy to update
const (
    MaxConnections = 100
    Timeout        = 30 * time.Second
    APIVersion     = "v1"
)

// Bad: Magic numbers in code
if connections > 100 {  // What is 100?
    // ...
}

// Good: Named constant
if connections > MaxConnections {
    // ...
}
```

### 2. Group Related Constants

```go
const (
    // HTTP Methods
    MethodGET     = "GET"
    MethodPOST    = "POST"
    MethodPUT     = "PUT"
    MethodDELETE  = "DELETE"
)

const (
    // Configuration
    DefaultPort = 8080
    DefaultHost = "localhost"
    BufferSize  = 1024
)
```

### 3. Use iota for Enumerations

```go
// Good: Clear sequence with iota
type Status int

const (
    Pending Status = iota
    Approved
    Rejected
)

// Avoid: Manual numbering (error-prone)
const (
    Pending  = 0
    Approved = 1
    Rejected = 2
)
```

### 4. Add String Methods to Enums

```go
type Color int

const (
    Red Color = iota
    Green
    Blue
)

func (c Color) String() string {
    return [...]string{"Red", "Green", "Blue"}[c]
}

// Usage
color := Red
fmt.Println(color)  // "Red" (not "0")
```

### 5. Use Typed Constants for Type Safety

```go
type UserID int
type ProductID int

const (
    DefaultUserID UserID = 1
    DefaultProductID ProductID = 1
)

// This prevents accidental misuse
func GetUser(id UserID) {}
func GetProduct(id ProductID) {}

// GetUser(DefaultProductID)  // Compile error: type mismatch
```

### 6. Document Constants

```go
const (
    // MaxRetries is the maximum number of retry attempts
    // before giving up on a failed operation.
    MaxRetries = 3

    // Timeout specifies how long to wait for a response
    // before considering the request as failed.
    Timeout = 30 * time.Second

    // BufferSize determines the size of the read/write buffer
    // in bytes. Must be a power of 2.
    BufferSize = 4096
)
```

### 7. Use Untyped Constants for Flexibility

```go
// Good: Untyped constants work with any compatible type
const Pi = 3.14159

var f32 float32 = Pi  // OK
var f64 float64 = Pi  // OK

// Less flexible: Typed constant
const TypedPi float64 = 3.14159

var x float32 = TypedPi         // Error: type mismatch
var y float32 = float32(TypedPi) // OK but requires conversion
```

### 8. Avoid Changing Constant Values

```go
// Good: Truly constant value
const DatabaseName = "production_db"

// Bad: Value that might need to change
const CurrentUser = "john"  // This should be a variable!
```

### 9. Use const for Better Performance

```go
// Constants are compile-time, no runtime overhead
const MaxSize = 1000

// Better than:
var maxSize = 1000  // Variable allocated at runtime
```

### 10. Organize Constants by Visibility

```go
// Public constants (exported)
const (
    DefaultTimeout = 30
    MaxRetries     = 3
)

// Private constants (unexported)
const (
    bufferSize = 4096
    maxWorkers = 10
)
```

## Summary

- Constants are immutable values known at compile time
- Declared with the `const` keyword
- Can be typed or untyped (untyped are more flexible)
- Support basic operations: arithmetic, string concatenation, boolean
- `iota` is used to create enumerated constants
- `iota` resets to 0 in each new `const` block
- Constants don't occupy memory at runtime
- Use constants for fixed values, configurations, and enumerations
- Add `String()` methods to enum types for better debugging
- Group related constants together
- Document complex or important constants
- Prefer untyped constants for flexibility unless type safety is needed

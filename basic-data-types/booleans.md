# Booleans in Go

The boolean type represents truth values and is fundamental to control flow and logical operations in Go programs.

## Table of Contents

- [Boolean Type](#boolean-type)
- [Declaration and Initialization](#declaration-and-initialization)
- [Boolean Values](#boolean-values)
- [Logical Operators](#logical-operators)
- [Comparison Operators](#comparison-operators)
- [Boolean Expressions](#boolean-expressions)
- [Short-Circuit Evaluation](#short-circuit-evaluation)
- [Type Conversion](#type-conversion)
- [Boolean in Control Flow](#boolean-in-control-flow)
- [Common Patterns](#common-patterns)
- [Best Practices](#best-practices)

## Boolean Type

Go has a built-in boolean type called `bool`.

| Type   | Size   | Values            |
| ------ | ------ | ----------------- |
| `bool` | 1 byte | `true` or `false` |

**Note:** Although a boolean only needs 1 bit of information, Go uses 1 byte for alignment and performance reasons.

## Declaration and Initialization

### Zero Value

The zero value of a boolean is `false`.

```go
var b bool  // b = false
```

### Explicit Declaration

```go
var isValid bool = true
var hasPermission bool = false
var isEnabled bool = true
```

### Short Declaration

```go
isActive := true
isReady := false
```

### Multiple Declaration

```go
var a, b, c bool = true, false, true
x, y, z := true, true, false
```

## Boolean Values

Go has only two boolean values:

```go
const (
    trueValue  = true
    falseValue = false
)
```

### Important Notes

- Boolean values are **not** integers
- `true` is **not** 1, and `false` is **not** 0
- Go does not allow implicit conversion between bool and other types

```go
// These are INVALID in Go
// var x int = true      // Compile error
// var b bool = 1        // Compile error
// if 1 { }              // Compile error
```

## Logical Operators

Go provides three logical operators for boolean values.

### AND Operator (&&)

Returns `true` only if both operands are `true`.

| A     | B     | A && B |
| ----- | ----- | ------ |
| false | false | false  |
| false | true  | false  |
| true  | false | false  |
| true  | true  | true   |

```go
a := true
b := false

result := a && b  // false
result = true && true  // true
```

### OR Operator (||)

Returns `true` if at least one operand is `true`.

| A     | B     | A \|\| B |
| ----- | ----- | -------- |
| false | false | false    |
| false | true  | true     |
| true  | false | true     |
| true  | true  | true     |

```go
a := true
b := false

result := a || b  // true
result = false || false  // false
```

### NOT Operator (!)

Inverts the boolean value.

| A     | !A    |
| ----- | ----- |
| false | true  |
| true  | false |

```go
a := true
result := !a  // false

b := false
result = !b  // true
```

### Combining Logical Operators

```go
a := true
b := false
c := true

// Complex expressions
result := a && b || c       // true  (false || true)
result = a && (b || c)      // true  (true && true)
result = !a || b && c       // false (!true || false)
result = !(a || b) && c     // false (false && true)
```

### Operator Precedence

From highest to lowest:

1. `!` (NOT)
2. `&&` (AND)
3. `||` (OR)

```go
a := true
b := false
c := true

// Without parentheses
result := !a || b && c  // Evaluated as: (!a) || (b && c)

// Explicit parentheses for clarity (recommended)
result = (!a) || (b && c)
```

## Comparison Operators

Comparison operators return boolean values.

### Equality Operators

```go
a := 10
b := 20

equal := a == b      // false (equal to)
notEqual := a != b   // true (not equal to)
```

### Relational Operators

```go
a := 10
b := 20

less := a < b           // true (less than)
greater := a > b        // false (greater than)
lessOrEqual := a <= b   // true (less than or equal)
greaterOrEqual := a >= b // false (greater than or equal)
```

### Comparing Booleans

```go
a := true
b := false

equal := a == b      // false
notEqual := a != b   // true

// Boolean comparison is less common than logical operators
// Prefer logical operators for combining booleans
```

## Boolean Expressions

Boolean expressions evaluate to `true` or `false`.

### Simple Expressions

```go
age := 25
isAdult := age >= 18  // true

temperature := 75
isComfortable := temperature >= 68 && temperature <= 78  // true
```

### Complex Expressions

```go
score := 85
isPassing := score >= 60
isExcellent := score >= 90
isGood := score >= 75 && score < 90

result := isPassing && (isExcellent || isGood)  // true
```

### Function Return Values

```go
func isEven(n int) bool {
    return n%2 == 0
}

func isPositive(n int) bool {
    return n > 0
}

func isValid(value string) bool {
    return len(value) > 0 && len(value) <= 100
}

// Usage
if isEven(10) {
    // ...
}

valid := isValid("hello")  // true
```

## Short-Circuit Evaluation

Go uses short-circuit evaluation for `&&` and `||` operators.

### AND Short-Circuit (&&)

If the left operand is `false`, the right operand is **not evaluated**.

```go
func expensiveCheck() bool {
    fmt.Println("Expensive check called")
    return true
}

// expensiveCheck() is NOT called because false && ... is always false
result := false && expensiveCheck()
// Output: (nothing printed)
```

### OR Short-Circuit (||)

If the left operand is `true`, the right operand is **not evaluated**.

```go
func expensiveCheck() bool {
    fmt.Println("Expensive check called")
    return false
}

// expensiveCheck() is NOT called because true || ... is always true
result := true || expensiveCheck()
// Output: (nothing printed)
```

### Practical Use Cases

```go
// Avoid nil pointer dereference
var ptr *string
if ptr != nil && *ptr == "value" {  // Safe: *ptr only evaluated if ptr != nil
    // ...
}

// Avoid division by zero
divisor := 0
if divisor != 0 && 100/divisor > 5 {  // Safe: division only if divisor != 0
    // ...
}

// Efficient validation
func validateUser(user *User) bool {
    // Check cheap conditions first
    return user != nil &&
           user.Age >= 18 &&
           expensiveEmailValidation(user.Email)
}
```

## Type Conversion

### Boolean to Integer

Go does **not** provide built-in conversion, but you can implement it:

```go
func boolToInt(b bool) int {
    if b {
        return 1
    }
    return 0
}

// Usage
value := boolToInt(true)  // 1
value = boolToInt(false)  // 0
```

### Integer to Boolean

```go
func intToBool(i int) bool {
    return i != 0
}

// Usage
flag := intToBool(1)   // true
flag = intToBool(0)    // false
flag = intToBool(42)   // true
```

### Boolean to String

```go
import (
    "fmt"
    "strconv"
)

// Method 1: strconv package
str := strconv.FormatBool(true)  // "true"
str = strconv.FormatBool(false)  // "false"

// Method 2: fmt.Sprintf
str = fmt.Sprintf("%t", true)    // "true"
str = fmt.Sprintf("%v", false)   // "false"
```

### String to Boolean

```go
import "strconv"

// ParseBool accepts: "1", "t", "T", "true", "TRUE", "True"
//                     "0", "f", "F", "false", "FALSE", "False"
b, err := strconv.ParseBool("true")   // true, nil
b, err = strconv.ParseBool("1")       // true, nil
b, err = strconv.ParseBool("false")   // false, nil
b, err = strconv.ParseBool("0")       // false, nil
b, err = strconv.ParseBool("invalid") // false, error
```

## Boolean in Control Flow

Booleans are primarily used in control flow statements.

### If Statements

```go
isValid := true

if isValid {
    fmt.Println("Valid")
}

if !isValid {
    fmt.Println("Invalid")
}

// With else
if isValid {
    fmt.Println("Valid")
} else {
    fmt.Println("Invalid")
}

// Short statement in if
if isAuthenticated := checkAuth(); isAuthenticated {
    // ...
}
```

### For Loops

```go
// Infinite loop
for true {
    // ...
    break
}

// Equivalent (preferred way)
for {
    // ...
    break
}

// Conditional loop
keepRunning := true
for keepRunning {
    // ...
    if someCondition {
        keepRunning = false
    }
}

// Traditional for loop with boolean
for i := 0; i < 10; i++ {  // i < 10 returns bool
    // ...
}
```

### Switch Statements

```go
// Switch with boolean expressions
score := 85

switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "F"
}

// Switch on boolean value
isActive := true
switch isActive {
case true:
    fmt.Println("Active")
case false:
    fmt.Println("Inactive")
}
```

## Common Patterns

### Flag Pattern

```go
// Command-line flags
var (
    verbose bool
    debug   bool
    dryRun  bool
)

// Feature flags
type FeatureFlags struct {
    EnableNewUI      bool
    EnableBetaFeature bool
    EnableLogging    bool
}

func (f FeatureFlags) IsNewUIEnabled() bool {
    return f.EnableNewUI
}
```

### State Tracking

```go
type Connection struct {
    isOpen      bool
    isAuthenticated bool
    isEncrypted bool
}

func (c *Connection) Open() error {
    if c.isOpen {
        return errors.New("already open")
    }
    // Open connection
    c.isOpen = true
    return nil
}

func (c *Connection) CanTransmit() bool {
    return c.isOpen && c.isAuthenticated && c.isEncrypted
}
```

### Validation Pattern

```go
func validateInput(input string) (bool, error) {
    if len(input) == 0 {
        return false, errors.New("input is empty")
    }
    if len(input) > 100 {
        return false, errors.New("input too long")
    }
    return true, nil
}

// Multiple validation checks
func isValidUser(user *User) bool {
    return user != nil &&
           user.Age >= 18 &&
           len(user.Name) > 0 &&
           isValidEmail(user.Email)
}
```

### Toggle Pattern

```go
type Setting struct {
    enabled bool
}

func (s *Setting) Toggle() {
    s.enabled = !s.enabled
}

func (s *Setting) Enable() {
    s.enabled = true
}

func (s *Setting) Disable() {
    s.enabled = false
}

func (s *Setting) IsEnabled() bool {
    return s.enabled
}
```

### Guard Clauses

```go
func processData(data []byte) error {
    // Guard clauses for early exit
    if data == nil {
        return errors.New("data is nil")
    }

    if len(data) == 0 {
        return errors.New("data is empty")
    }

    if !isValidFormat(data) {
        return errors.New("invalid format")
    }

    // Main logic here
    return nil
}
```

### Option Pattern

```go
type Server struct {
    host    string
    port    int
    verbose bool
    debug   bool
}

type Option func(*Server)

func WithVerbose(verbose bool) Option {
    return func(s *Server) {
        s.verbose = verbose
    }
}

func WithDebug(debug bool) Option {
    return func(s *Server) {
        s.debug = debug
    }
}

// Usage
server := NewServer(
    WithVerbose(true),
    WithDebug(false),
)
```

## Best Practices

### 1. Use Descriptive Names

```go
// Bad
var f bool
var x bool

// Good
var isValid bool
var hasPermission bool
var canExecute bool
var shouldRetry bool
var wasSuccessful bool
```

### 2. Prefer Positive Names

```go
// Less clear
var notReady bool = false
if !notReady {  // Double negative
    // ...
}

// Better
var isReady bool = true
if isReady {
    // ...
}
```

### 3. Use Boolean Functions

```go
// Good: Clear intent
func isEmpty(s string) bool {
    return len(s) == 0
}

func isAuthenticated(user *User) bool {
    return user != nil && user.Token != ""
}

// Usage
if isEmpty(input) {
    // ...
}
```

### 4. Avoid Redundant Comparisons

```go
// Bad
if isValid == true {
    // ...
}

if isActive == false {
    // ...
}

// Good
if isValid {
    // ...
}

if !isActive {
    // ...
}
```

### 5. Use Short-Circuit Evaluation Wisely

```go
// Expensive checks last
if cheapCheck() && expensiveCheck() {
    // expensiveCheck() only called if cheapCheck() is true
}

// Safe null checks
if user != nil && user.IsActive() {
    // IsActive() only called if user is not nil
}
```

### 6. Avoid Complex Boolean Expressions

```go
// Bad: Hard to read
if !(!a || b) && (c || d) && !(e && f) {
    // ...
}

// Good: Extract to functions
func shouldProceed(a, b, c, d, e, f bool) bool {
    condition1 := a && !b
    condition2 := c || d
    condition3 := !(e && f)
    return condition1 && condition2 && condition3
}

if shouldProceed(a, b, c, d, e, f) {
    // ...
}
```

### 7. Use Constants for Boolean Flags

```go
const (
    Enabled  = true
    Disabled = false
)

func NewServer(logging bool) *Server {
    return &Server{
        loggingEnabled: logging,
    }
}

// Usage
server := NewServer(Enabled)
```

### 8. Table-Driven Tests

```go
func TestIsEven(t *testing.T) {
    tests := []struct {
        name     string
        input    int
        expected bool
    }{
        {"zero", 0, true},
        {"positive even", 4, true},
        {"positive odd", 5, false},
        {"negative even", -6, true},
        {"negative odd", -7, false},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := isEven(tt.input)
            if result != tt.expected {
                t.Errorf("got %v, want %v", result, tt.expected)
            }
        })
    }
}
```

### 9. Return Early with Boolean Guards

```go
func canProcess(user *User, data []byte) bool {
    if user == nil {
        return false
    }

    if !user.IsAuthenticated() {
        return false
    }

    if len(data) == 0 {
        return false
    }

    return true
}
```

### 10. Use Boolean Methods for Readability

```go
type Order struct {
    status string
    paid   bool
    shipped bool
}

func (o *Order) IsComplete() bool {
    return o.paid && o.shipped
}

func (o *Order) CanCancel() bool {
    return !o.shipped
}

func (o *Order) NeedsPayment() bool {
    return !o.paid
}

// Usage
if order.IsComplete() {
    // ...
}
```

## Summary

- Go's boolean type is `bool` with values `true` and `false`
- The zero value is `false`
- Logical operators: `&&` (AND), `||` (OR), `!` (NOT)
- Go uses short-circuit evaluation for efficiency
- Booleans are not integers and require explicit conversion
- Use descriptive, positive names for boolean variables
- Avoid redundant comparisons (`== true`, `== false`)
- Extract complex boolean logic into well-named functions
- Booleans are essential for control flow (if, for, switch)
- Use guard clauses for cleaner code

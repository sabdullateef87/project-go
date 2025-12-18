# Strings in Go

Strings are one of the most commonly used data types in Go, representing sequences of characters encoded in UTF-8.

## Table of Contents

- [String Type](#string-type)
- [Declaration and Initialization](#declaration-and-initialization)
- [String Literals](#string-literals)
- [String Properties](#string-properties)
- [Accessing Characters](#accessing-characters)
- [String Concatenation](#string-concatenation)
- [String Comparison](#string-comparison)
- [String Iteration](#string-iteration)
- [Common String Operations](#common-string-operations)
- [String Conversion](#string-conversion)
- [String Formatting](#string-formatting)
- [String Builder](#string-builder)
- [Strings Package](#strings-package)
- [Unicode and Runes](#unicode-and-runes)
- [String Immutability](#string-immutability)
- [Best Practices](#best-practices)

## String Type

Strings in Go are:

- **Immutable** - Once created, cannot be modified
- **UTF-8 encoded** - Support for international characters
- **Reference types** - Stored as a pointer to underlying byte array with length

| Type     | Size     | Description                          |
| -------- | -------- | ------------------------------------ |
| `string` | Variable | Sequence of UTF-8 encoded characters |

## Declaration and Initialization

### Zero Value

The zero value of a string is an empty string `""`.

```go
var s string  // s = ""
```

### Explicit Declaration

```go
var name string = "Alice"
var message string = "Hello, World!"
var empty string = ""
```

### Short Declaration

```go
greeting := "Hello"
text := "Go programming"
```

### Multiple Declaration

```go
var first, last, middle string = "John", "Doe", "Michael"
name, age, city := "Alice", "25", "New York"
```

## String Literals

### Interpreted String Literals (Double Quotes)

Use double quotes for strings with escape sequences.

```go
s := "Hello, World!"
s = "Line 1\nLine 2"  // Newline
s = "Tab\tseparated"  // Tab
s = "Quote: \"text\""  // Escaped quote
s = "Backslash: \\"   // Escaped backslash
```

### Common Escape Sequences

| Escape       | Meaning                       |
| ------------ | ----------------------------- |
| `\n`         | Newline                       |
| `\t`         | Tab                           |
| `\r`         | Carriage return               |
| `\\`         | Backslash                     |
| `\"`         | Double quote                  |
| `\'`         | Single quote                  |
| `\x00`       | Hex byte (e.g., `\x41` = 'A') |
| `\u0000`     | Unicode character (16-bit)    |
| `\U00000000` | Unicode character (32-bit)    |

```go
// Unicode examples
star := "\u2605"        // ★
emoji := "\U0001F600"   // 😀
hex := "\x48\x65\x6C\x6C\x6F"  // Hello
```

### Raw String Literals (Backticks)

Use backticks for raw strings (no escape sequences, can span multiple lines).

```go
// Single line
path := `C:\Users\Alice\Documents`

// Multi-line
html := `
<html>
    <body>
        <h1>Hello</h1>
    </body>
</html>
`

// Regular expression (no need to escape backslashes)
regex := `\d{3}-\d{2}-\d{4}`

// JSON
json := `{"name": "Alice", "age": 30}`
```

## String Properties

### Length

```go
s := "Hello"
length := len(s)  // 5 (number of bytes, not characters)

// UTF-8 string with multi-byte characters
s2 := "Hello, 世界"
byteLen := len(s2)  // 13 bytes (not 9 characters)

// To count characters (runes)
import "unicode/utf8"
runeCount := utf8.RuneCountInString(s2)  // 9 characters
```

### Empty String Check

```go
s := ""

// Method 1: Compare with empty string
isEmpty := s == ""

// Method 2: Check length
isEmpty = len(s) == 0  // Preferred (more efficient)
```

## Accessing Characters

### Indexing (Byte Access)

```go
s := "Hello"

// Access individual bytes
firstByte := s[0]    // 'H' (byte value: 72)
secondByte := s[1]   // 'e' (byte value: 101)

// Get substring by byte range
sub := s[1:4]  // "ell" (bytes from index 1 to 3)
sub = s[:3]    // "Hel" (first 3 bytes)
sub = s[2:]    // "llo" (from index 2 to end)
sub = s[:]     // "Hello" (copy entire string)
```

### Important Note on UTF-8

Indexing gives you bytes, not characters. For multi-byte UTF-8 characters:

```go
s := "Hello, 世界"

// WRONG: This accesses bytes, not characters
// s[7] gives you a byte from the middle of a multi-byte character

// RIGHT: Use runes for character access
runes := []rune(s)
char := runes[7]  // '世'
```

## String Concatenation

### Using + Operator

```go
s1 := "Hello"
s2 := "World"
result := s1 + ", " + s2  // "Hello, World"

// Multi-line concatenation
message := "This is " +
           "a long " +
           "message"
```

### Using += Operator

```go
s := "Hello"
s += ", "
s += "World"  // "Hello, World"
```

### Using fmt.Sprintf

```go
import "fmt"

name := "Alice"
age := 30
message := fmt.Sprintf("%s is %d years old", name, age)
// "Alice is 30 years old"
```

### Using strings.Join

```go
import "strings"

parts := []string{"Hello", "World", "from", "Go"}
result := strings.Join(parts, " ")  // "Hello World from Go"
result = strings.Join(parts, ", ")  // "Hello, World, from, Go"
```

### Using strings.Builder (Most Efficient)

```go
import "strings"

var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(", ")
builder.WriteString("World")
result := builder.String()  // "Hello, World"
```

## String Comparison

### Equality

```go
s1 := "Hello"
s2 := "Hello"
s3 := "hello"

equal := s1 == s2     // true
equal = s1 == s3      // false (case-sensitive)
notEqual := s1 != s3  // true
```

### Lexicographic Comparison

```go
s1 := "apple"
s2 := "banana"

less := s1 < s2          // true
greater := s1 > s2       // false
lessEqual := s1 <= s2    // true
greaterEqual := s1 >= s2 // false
```

### Case-Insensitive Comparison

```go
import "strings"

s1 := "Hello"
s2 := "hello"

equal := strings.EqualFold(s1, s2)  // true

// Manual conversion
equal = strings.ToLower(s1) == strings.ToLower(s2)  // true
```

### Comparing with strings.Compare

```go
import "strings"

result := strings.Compare("a", "b")  // -1 (a < b)
result = strings.Compare("b", "a")   // 1 (b > a)
result = strings.Compare("a", "a")   // 0 (a == a)
```

## String Iteration

### Iterating Over Bytes

```go
s := "Hello"

// Method 1: For loop with index
for i := 0; i < len(s); i++ {
    fmt.Printf("Byte %d: %c\n", i, s[i])
}

// Method 2: Range (gives bytes for ASCII)
for i, b := range s {
    fmt.Printf("Byte %d: %c\n", i, b)
}
```

### Iterating Over Runes (Characters)

```go
s := "Hello, 世界"

// Range automatically decodes UTF-8 runes
for i, r := range s {
    fmt.Printf("Index %d: %c (Unicode: %U)\n", i, r, r)
}
// Output:
// Index 0: H (Unicode: U+0048)
// Index 1: e (Unicode: U+0065)
// ...
// Index 7: 世 (Unicode: U+4E16)
// Index 10: 界 (Unicode: U+754C)
```

### Converting to Rune Slice

```go
s := "Hello, 世界"
runes := []rune(s)

for i, r := range runes {
    fmt.Printf("Character %d: %c\n", i, r)
}
```

## Common String Operations

### Contains

```go
import "strings"

s := "Hello, World!"

contains := strings.Contains(s, "World")  // true
contains = strings.Contains(s, "world")   // false (case-sensitive)
contains = strings.Contains(s, "xyz")     // false
```

### HasPrefix and HasSuffix

```go
import "strings"

s := "filename.txt"

hasPrefix := strings.HasPrefix(s, "file")    // true
hasSuffix := strings.HasSuffix(s, ".txt")    // true
hasSuffix = strings.HasSuffix(s, ".pdf")     // false
```

### Index (Find Position)

```go
import "strings"

s := "Hello, World!"

index := strings.Index(s, "World")      // 7
index = strings.Index(s, "xyz")         // -1 (not found)
lastIndex := strings.LastIndex(s, "o")  // 8
```

### Count Occurrences

```go
import "strings"

s := "banana"
count := strings.Count(s, "a")   // 3
count = strings.Count(s, "an")   // 2
count = strings.Count(s, "xyz")  // 0
```

### Replace

```go
import "strings"

s := "Hello, World!"

// Replace all occurrences
result := strings.ReplaceAll(s, "o", "0")  // "Hell0, W0rld!"

// Replace n occurrences (-1 means all)
result = strings.Replace(s, "l", "L", 2)   // "HeLLo, World!"
result = strings.Replace(s, "l", "L", -1)  // "HeLLo, WorLd!" (all)
```

### Trim

```go
import "strings"

s := "  Hello, World!  "

// Trim whitespace
trimmed := strings.TrimSpace(s)  // "Hello, World!"

// Trim specific characters
s2 := "!!!Hello!!!"
trimmed = strings.Trim(s2, "!")       // "Hello"
trimmed = strings.TrimLeft(s2, "!")   // "Hello!!!"
trimmed = strings.TrimRight(s2, "!")  // "!!!Hello"

// Trim prefix/suffix
s3 := "filename.txt"
trimmed = strings.TrimPrefix(s3, "file")   // "name.txt"
trimmed = strings.TrimSuffix(s3, ".txt")   // "filename"
```

### Split

```go
import "strings"

s := "apple,banana,cherry"

// Split by delimiter
parts := strings.Split(s, ",")  // ["apple", "banana", "cherry"]

// Split with limit
parts = strings.SplitN(s, ",", 2)  // ["apple", "banana,cherry"]

// Split by whitespace
s2 := "one two  three"
parts = strings.Fields(s2)  // ["one", "two", "three"]
```

### Case Conversion

```go
import "strings"

s := "Hello, World!"

upper := strings.ToUpper(s)  // "HELLO, WORLD!"
lower := strings.ToLower(s)  // "hello, world!"
title := strings.Title(s)    // "Hello, World!" (deprecated)

// Modern title case
import "golang.org/x/text/cases"
import "golang.org/x/text/language"

caser := cases.Title(language.English)
title = caser.String(s)
```

### Repeat

```go
import "strings"

s := "Go"
repeated := strings.Repeat(s, 3)  // "GoGoGo"
line := strings.Repeat("-", 20)   // "--------------------"
```

## String Conversion

### String to/from Byte Slice

```go
// String to byte slice
s := "Hello"
bytes := []byte(s)  // [72 101 108 108 111]

// Byte slice to string
s2 := string(bytes)  // "Hello"
```

### String to/from Rune Slice

```go
// String to rune slice
s := "Hello, 世界"
runes := []rune(s)  // [72 101 108 108 111 44 32 19990 30028]

// Rune slice to string
s2 := string(runes)  // "Hello, 世界"
```

### Number to String

```go
import (
    "fmt"
    "strconv"
)

// Integer to string
num := 42
str := strconv.Itoa(num)              // "42"
str = strconv.FormatInt(int64(num), 10)  // "42" (base 10)
str = fmt.Sprintf("%d", num)          // "42"

// Float to string
f := 3.14159
str = strconv.FormatFloat(f, 'f', 2, 64)  // "3.14"
str = fmt.Sprintf("%.2f", f)              // "3.14"

// Boolean to string
b := true
str = strconv.FormatBool(b)  // "true"
str = fmt.Sprintf("%t", b)   // "true"
```

### String to Number

```go
import "strconv"

// String to integer
str := "42"
num, err := strconv.Atoi(str)  // 42, nil

// String to int64
num64, err := strconv.ParseInt(str, 10, 64)  // 42, nil

// String to float
str = "3.14"
f, err := strconv.ParseFloat(str, 64)  // 3.14, nil

// String to boolean
str = "true"
b, err := strconv.ParseBool(str)  // true, nil
```

## String Formatting

### fmt.Printf and fmt.Sprintf

```go
import "fmt"

name := "Alice"
age := 30
height := 5.6

// Printf (prints to stdout)
fmt.Printf("Name: %s, Age: %d, Height: %.1f\n", name, age, height)

// Sprintf (returns string)
message := fmt.Sprintf("Name: %s, Age: %d", name, age)
```

### Format Verbs

| Verb   | Description              | Example                                    |
| ------ | ------------------------ | ------------------------------------------ |
| `%v`   | Default format           | `fmt.Sprintf("%v", 42)` → "42"             |
| `%+v`  | Struct with field names  | `fmt.Sprintf("%+v", person)`               |
| `%#v`  | Go-syntax representation | `fmt.Sprintf("%#v", []int{1,2})`           |
| `%T`   | Type                     | `fmt.Sprintf("%T", 42)` → "int"            |
| `%s`   | String                   | `fmt.Sprintf("%s", "hello")`               |
| `%q`   | Quoted string            | `fmt.Sprintf("%q", "hello")` → "\"hello\"" |
| `%d`   | Decimal integer          | `fmt.Sprintf("%d", 42)`                    |
| `%b`   | Binary                   | `fmt.Sprintf("%b", 42)` → "101010"         |
| `%x`   | Hexadecimal              | `fmt.Sprintf("%x", 42)` → "2a"             |
| `%f`   | Float                    | `fmt.Sprintf("%f", 3.14)` → "3.140000"     |
| `%.2f` | Float with precision     | `fmt.Sprintf("%.2f", 3.14159)` → "3.14"    |
| `%t`   | Boolean                  | `fmt.Sprintf("%t", true)` → "true"         |
| `%p`   | Pointer                  | `fmt.Sprintf("%p", &x)`                    |

### Width and Alignment

```go
import "fmt"

// Width
fmt.Sprintf("%5d", 42)      // "   42" (right-aligned, width 5)
fmt.Sprintf("%-5d", 42)     // "42   " (left-aligned)
fmt.Sprintf("%05d", 42)     // "00042" (zero-padded)

// Strings
fmt.Sprintf("%10s", "Go")   // "        Go"
fmt.Sprintf("%-10s", "Go")  // "Go        "
```

## String Builder

For efficient string concatenation in loops, use `strings.Builder`.

```go
import "strings"

var builder strings.Builder

// Build string efficiently
for i := 0; i < 100; i++ {
    builder.WriteString("item ")
    builder.WriteString(strconv.Itoa(i))
    builder.WriteByte('\n')
}

result := builder.String()

// Check size and capacity
size := builder.Len()      // Current length
cap := builder.Cap()       // Current capacity

// Reset builder
builder.Reset()

// Pre-allocate capacity
var builder2 strings.Builder
builder2.Grow(1000)  // Pre-allocate for 1000 bytes
```

## Strings Package

The `strings` package provides many useful functions:

```go
import "strings"

// Case conversion
strings.ToUpper("hello")     // "HELLO"
strings.ToLower("HELLO")     // "hello"

// Search
strings.Contains("hello", "ell")      // true
strings.HasPrefix("hello", "he")      // true
strings.HasSuffix("hello", "lo")      // true
strings.Index("hello", "ll")          // 2
strings.Count("hello", "l")           // 2

// Modification
strings.Replace("hello", "l", "L", 1) // "heLlo"
strings.ReplaceAll("hello", "l", "L") // "heLLo"
strings.Trim("  hello  ", " ")        // "hello"
strings.TrimSpace("  hello  ")        // "hello"

// Split and join
strings.Split("a,b,c", ",")           // ["a", "b", "c"]
strings.Join([]string{"a","b"}, "-")  // "a-b"
strings.Fields("a b  c")              // ["a", "b", "c"]

// Repeat
strings.Repeat("Go", 3)               // "GoGoGo"
```

## Unicode and Runes

### Rune Type

A `rune` is an alias for `int32` and represents a Unicode code point.

```go
var r rune = 'A'     // Unicode code point U+0041
var r2 rune = '世'    // Unicode code point U+4E16
var r3 rune = '😀'    // Unicode code point U+1F600
```

### Working with Unicode

```go
import (
    "unicode"
    "unicode/utf8"
)

s := "Hello, 世界! 123"

// Count runes (characters)
count := utf8.RuneCountInString(s)  // 13

// Check if valid UTF-8
valid := utf8.ValidString(s)  // true

// Decode first rune
r, size := utf8.DecodeRuneInString(s)
// r = 'H', size = 1

// Iterate runes manually
for i, w := 0, 0; i < len(s); i += w {
    r, width := utf8.DecodeRuneInString(s[i:])
    fmt.Printf("%c ", r)
    w = width
}

// Character classification
unicode.IsLetter('A')    // true
unicode.IsDigit('5')     // true
unicode.IsSpace(' ')     // true
unicode.IsUpper('A')     // true
unicode.IsLower('a')     // true
unicode.ToUpper('a')     // 'A'
unicode.ToLower('A')     // 'a'
```

## String Immutability

Strings in Go are immutable - you cannot modify a string in place.

```go
s := "Hello"

// INVALID: Cannot assign to string index
// s[0] = 'h'  // Compile error

// Must create new string
s = "hello"

// Or use byte slice for modifications
bytes := []byte(s)
bytes[0] = 'h'
s = string(bytes)  // "hello"
```

### Why Immutability Matters

```go
// Strings can be safely shared
s1 := "Hello"
s2 := s1  // Both point to same underlying data (efficient)

// Modification creates new string
s2 = s2 + ", World"  // s1 is unchanged
```

## Best Practices

### 1. Use strings.Builder for Concatenation in Loops

```go
// Bad: Creates many intermediate strings
var result string
for i := 0; i < 1000; i++ {
    result += "item"
}

// Good: Efficient building
var builder strings.Builder
for i := 0; i < 1000; i++ {
    builder.WriteString("item")
}
result := builder.String()
```

### 2. Use Raw Strings for Paths and Regex

```go
// Bad: Escaping backslashes
path := "C:\\Users\\Alice\\Documents"

// Good: Raw string
path := `C:\Users\Alice\Documents`

// Regular expressions
regex := `\d{3}-\d{2}-\d{4}`
```

### 3. Check for Empty Strings with len()

```go
// Preferred (more efficient)
if len(s) == 0 {
    // ...
}

// Also acceptable
if s == "" {
    // ...
}
```

### 4. Use strings Package Functions

```go
// Don't reinvent the wheel
import "strings"

// Use built-in functions
if strings.Contains(s, "search") {
    // ...
}
```

### 5. Be Aware of UTF-8 Encoding

```go
s := "Hello, 世界"

// len() returns bytes, not characters
byteCount := len(s)  // 13

// Use utf8 package for character count
import "unicode/utf8"
charCount := utf8.RuneCountInString(s)  // 9

// Use range for proper character iteration
for i, r := range s {
    // r is a rune (character)
}
```

### 6. Use strconv for Conversions

```go
// Preferred
import "strconv"
s := strconv.Itoa(42)

// Avoid unnecessary Sprintf for simple conversions
// s := fmt.Sprintf("%d", 42)  // Slower
```

### 7. Pre-allocate strings.Builder Capacity

```go
var builder strings.Builder
builder.Grow(expectedSize)  // Pre-allocate memory
for _, item := range items {
    builder.WriteString(item)
}
```

### 8. Use Efficient String Comparison

```go
// For case-insensitive comparison
import "strings"

if strings.EqualFold(s1, s2) {
    // More efficient than ToLower() + ==
}
```

## Summary

- Strings are immutable UTF-8 encoded sequences of characters
- Zero value is empty string `""`
- Use double quotes for interpreted strings, backticks for raw strings
- `len()` returns byte count, not character count
- Use `[]rune` conversion or range for character-level operations
- The `strings` package provides comprehensive string manipulation
- Use `strings.Builder` for efficient concatenation
- Always check errors when converting strings to numbers
- Be mindful of UTF-8 multi-byte characters
- Strings are safe for concurrent read access

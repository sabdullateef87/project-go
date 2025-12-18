# Composite Types in Go

Composite types allow you to combine basic types into more complex data structures. Go provides powerful composite types for organizing and managing data efficiently.

## Table of Contents

1. [Arrays](arrays.md) - Fixed-size sequences of elements
2. [Slices](slices.md) - Dynamic, flexible sequences
3. [Maps](maps.md) - Key-value data structures
4. [Structs](structs.md) - Custom data types with fields
5. [JSON](json.md) - JSON encoding and decoding
6. [Text Templates](text-templates.md) - Text template processing
7. [HTML Templates](html-templates.md) - HTML template rendering

## Overview

### Core Composite Types

| Type       | Description         | Fixed Size | Ordered | Key-Value |
| ---------- | ------------------- | ---------- | ------- | --------- |
| **Array**  | Fixed-size sequence | ✅ Yes     | ✅ Yes  | ❌ No     |
| **Slice**  | Dynamic sequence    | ❌ No      | ✅ Yes  | ❌ No     |
| **Map**    | Key-value pairs     | ❌ No      | ❌ No   | ✅ Yes    |
| **Struct** | Named fields        | ✅ Yes     | ❌ No   | ❌ No     |

### Data Format Support

| Format             | Package         | Purpose                          |
| ------------------ | --------------- | -------------------------------- |
| **JSON**           | `encoding/json` | Serialize/deserialize JSON data  |
| **Text Templates** | `text/template` | Generate text from templates     |
| **HTML Templates** | `html/template` | Generate HTML with auto-escaping |

## Quick Comparison

### Arrays vs Slices

```go
// Array: Fixed size, value type
var arr [5]int = [5]int{1, 2, 3, 4, 5}

// Slice: Dynamic size, reference type
var slice []int = []int{1, 2, 3, 4, 5}
```

### Maps vs Structs

```go
// Map: Dynamic keys, homogeneous values
userMap := map[string]string{
    "name": "Alice",
    "email": "alice@example.com",
}

// Struct: Fixed fields, heterogeneous types
type User struct {
    Name  string
    Email string
    Age   int
}
```

## Zero Values

Each composite type has a specific zero value:

```go
var arr [5]int        // [0 0 0 0 0] - array of zero values
var slice []int       // nil - nil slice
var m map[string]int  // nil - nil map
var s struct{ X int } // {0} - struct with zero-valued fields
```

## Type Categories

### Collection Types

**Arrays** and **Slices** store ordered sequences of elements:

```go
// Arrays
numbers := [3]int{1, 2, 3}
length := len(numbers)  // 3

// Slices
scores := []int{95, 87, 92}
scores = append(scores, 88)  // Dynamic growth
```

**Maps** store key-value pairs:

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}
ages["Charlie"] = 35  // Add new entry
```

### Structured Types

**Structs** group related data with named fields:

```go
type Person struct {
    Name string
    Age  int
}

person := Person{Name: "Alice", Age: 30}
```

### Serialization Support

**JSON** encoding/decoding for data interchange:

```go
import "encoding/json"

// Struct to JSON
data, _ := json.Marshal(person)

// JSON to struct
json.Unmarshal(data, &person)
```

### Template Processing

**Text and HTML Templates** for generating formatted output:

```go
import "text/template"

tmpl := template.Must(template.New("test").Parse("Hello, {{.Name}}!"))
tmpl.Execute(os.Stdout, person)
```

## Common Operations

### Iteration

```go
// Array/Slice iteration
for i, v := range slice {
    fmt.Printf("Index %d: %v\n", i, v)
}

// Map iteration
for key, value := range myMap {
    fmt.Printf("Key %s: %v\n", key, value)
}

// Struct field iteration (using reflection)
import "reflect"
v := reflect.ValueOf(myStruct)
for i := 0; i < v.NumField(); i++ {
    fmt.Printf("Field %d: %v\n", i, v.Field(i))
}
```

### Length and Capacity

```go
arr := [5]int{1, 2, 3, 4, 5}
slice := []int{1, 2, 3, 4, 5}

// Arrays: len only
len(arr)  // 5

// Slices: len and cap
len(slice)  // 5
cap(slice)  // 5

// Maps: len only
len(myMap)  // Number of entries
```

## Memory Characteristics

### Value vs Reference Types

**Value Types** (copied on assignment):

- Arrays
- Structs

**Reference Types** (share underlying data):

- Slices
- Maps

```go
// Array: Value type
arr1 := [3]int{1, 2, 3}
arr2 := arr1      // Copy
arr2[0] = 99      // arr1 unchanged

// Slice: Reference type
slice1 := []int{1, 2, 3}
slice2 := slice1  // Shares data
slice2[0] = 99    // slice1[0] is now 99
```

## Best Practices

### 1. Choose the Right Type

```go
// Use arrays for fixed-size collections
var matrix [3][3]int

// Use slices for dynamic collections
var items []Item

// Use maps for lookups
var cache map[string]Value

// Use structs for grouped data
type Config struct {
    Host string
    Port int
}
```

### 2. Initialize Properly

```go
// Maps must be initialized before use
m := make(map[string]int)  // Good
// var m map[string]int    // Bad: nil map, cannot add entries

// Slices can be nil or initialized
var s []int                // OK: nil slice
s = make([]int, 0, 10)     // OK: with capacity
s = []int{}                // OK: empty slice
```

### 3. Check for nil

```go
// Check slice/map before use
if slice != nil {
    // Safe to use
}

if m != nil && len(m) > 0 {
    // Map has entries
}
```

### 4. Use Composite Literals

```go
// Good: Clear initialization
person := Person{
    Name: "Alice",
    Age:  30,
}

// Good: Inline slice initialization
numbers := []int{1, 2, 3, 4, 5}

// Good: Inline map initialization
config := map[string]string{
    "host": "localhost",
    "port": "8080",
}
```

## Learning Path

For comprehensive understanding of Go's composite types, follow this order:

1. **Start with [Arrays](arrays.md)** - Understand fixed-size collections and array fundamentals
2. **Move to [Slices](slices.md)** - Learn dynamic sequences and slice internals
3. **Learn [Maps](maps.md)** - Master key-value data structures
4. **Study [Structs](structs.md)** - Create custom types and understand composition
5. **Explore [JSON](json.md)** - Handle data serialization and APIs
6. **Try [Text Templates](text-templates.md)** - Generate formatted text output
7. **Master [HTML Templates](html-templates.md)** - Build web applications safely

## Common Patterns

### Builder Pattern

```go
type Builder struct {
    items []string
}

func (b *Builder) Add(item string) *Builder {
    b.items = append(b.items, item)
    return b
}

func (b *Builder) Build() []string {
    return b.items
}

// Usage
items := (&Builder{}).Add("a").Add("b").Add("c").Build()
```

### Options Pattern

```go
type Server struct {
    Host string
    Port int
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) { s.Host = host }
}

func WithPort(port int) Option {
    return func(s *Server) { s.Port = port }
}

func NewServer(opts ...Option) *Server {
    s := &Server{Host: "localhost", Port: 8080}
    for _, opt := range opts {
        opt(s)
    }
    return s
}

// Usage
server := NewServer(WithHost("0.0.0.0"), WithPort(9000))
```

### Data Transfer Object (DTO)

```go
type UserDTO struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

func (u *User) ToDTO() UserDTO {
    return UserDTO{
        ID:    u.ID,
        Name:  u.Name,
        Email: u.Email,
    }
}
```

## Performance Considerations

### Slice Growth

```go
// Bad: Repeated append without capacity
var slice []int
for i := 0; i < 1000; i++ {
    slice = append(slice, i)  // Multiple reallocations
}

// Good: Pre-allocate capacity
slice := make([]int, 0, 1000)
for i := 0; i < 1000; i++ {
    slice = append(slice, i)  // No reallocation
}
```

### Map Pre-allocation

```go
// Good: Pre-allocate for known size
m := make(map[string]int, 100)
```

### Struct Field Ordering

```go
// Less optimal: 24 bytes (with padding)
type Bad struct {
    a bool   // 1 byte + 7 padding
    b int64  // 8 bytes
    c bool   // 1 byte + 7 padding
}

// Optimized: 16 bytes
type Good struct {
    b int64  // 8 bytes
    a bool   // 1 byte
    c bool   // 1 byte + 6 padding
}
```

## Additional Resources

- [Go Language Specification - Types](https://go.dev/ref/spec#Types)
- [Effective Go - Data](https://go.dev/doc/effective_go#data)
- [Go by Example - Collections](https://gobyexample.com/arrays)
- [Go Blog - Slices: usage and internals](https://go.dev/blog/slices-intro)

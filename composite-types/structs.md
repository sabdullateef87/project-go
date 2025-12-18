# Structs in Go

Structs are composite data types that group together fields of different types under a single name. They're fundamental to organizing and modeling data in Go.

## Table of Contents

- [What are Structs?](#what-are-structs)
- [Declaration and Initialization](#declaration-and-initialization)
- [Accessing Fields](#accessing-fields)
- [Anonymous Structs](#anonymous-structs)
- [Struct Embedding](#struct-embedding)
- [Methods](#methods)
- [Pointer Receivers vs Value Receivers](#pointer-receivers-vs-value-receivers)
- [Struct Tags](#struct-tags)
- [Struct Comparison](#struct-comparison)
- [Struct Copying](#struct-copying)
- [Common Patterns](#common-patterns)
- [Memory Layout](#memory-layout)
- [Best Practices](#best-practices)

## What are Structs?

A struct is:

- **Composite type** - Groups multiple fields together
- **Value type** - Copied when assigned or passed to functions
- **Named fields** - Each field has a name and type
- **Custom types** - Define your own data structures
- **Can have methods** - Attach behavior to data

```go
type Person struct {
    Name string
    Age  int
}
```

## Declaration and Initialization

### Defining a Struct Type

```go
// Basic struct
type Person struct {
    Name string
    Age  int
}

// Struct with multiple field types
type Address struct {
    Street  string
    City    string
    State   string
    ZipCode string
}

// Struct with embedded types
type Employee struct {
    ID       int
    Person   Person
    Address  Address
    Salary   float64
}
```

### Zero Value

Uninitialized structs have zero values for all fields.

```go
var p Person
// p.Name = "" (zero value for string)
// p.Age = 0 (zero value for int)
```

### Struct Literal Initialization

```go
// Method 1: With field names (preferred)
p := Person{
    Name: "Alice",
    Age:  30,
}

// Method 2: Positional (order matters, discouraged)
p := Person{"Alice", 30}

// Method 3: Partial initialization (rest are zero values)
p := Person{Name: "Bob"}  // Age will be 0

// Method 4: Empty initialization
p := Person{}  // All fields are zero values
```

### Using new()

`new()` allocates memory and returns a pointer to a zero-valued struct.

```go
p := new(Person)  // Returns *Person
// Equivalent to:
p := &Person{}

// Access fields through pointer
p.Name = "Alice"
p.Age = 30
```

### Pointer to Struct

```go
// Create and get pointer
p := &Person{
    Name: "Alice",
    Age:  30,
}

// new() returns pointer
p := new(Person)
p.Name = "Alice"
```

## Accessing Fields

### Dot Notation

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}

// Read fields
name := p.Name  // "Alice"
age := p.Age    // 30

// Write fields
p.Name = "Bob"
p.Age = 25
```

### Accessing Through Pointers

Go automatically dereferences struct pointers.

```go
p := &Person{Name: "Alice", Age: 30}

// These are equivalent:
name := p.Name      // Automatic dereference
name := (*p).Name   // Explicit dereference

// Assignment also works
p.Name = "Bob"      // Automatic dereference
(*p).Name = "Bob"   // Explicit dereference
```

### Nested Struct Access

```go
type Address struct {
    Street string
    City   string
}

type Person struct {
    Name    string
    Address Address
}

p := Person{
    Name: "Alice",
    Address: Address{
        Street: "123 Main St",
        City:   "NYC",
    },
}

// Access nested fields
city := p.Address.City
p.Address.Street = "456 Oak Ave"
```

## Anonymous Structs

Structs without a defined type name.

### Inline Declaration

```go
// Declare and initialize
person := struct {
    Name string
    Age  int
}{
    Name: "Alice",
    Age:  30,
}

fmt.Println(person.Name)  // "Alice"
```

### Use Cases

```go
// 1. One-time data structures
func getConfig() struct {
    Host string
    Port int
} {
    return struct {
        Host string
        Port int
    }{
        Host: "localhost",
        Port: 8080,
    }
}

// 2. Table-driven tests
tests := []struct {
    input    int
    expected int
}{
    {input: 1, expected: 2},
    {input: 2, expected: 4},
    {input: 3, expected: 6},
}

// 3. JSON unmarshaling for one-off structures
var response struct {
    Status string `json:"status"`
    Data   struct {
        ID   int    `json:"id"`
        Name string `json:"name"`
    } `json:"data"`
}
json.Unmarshal(data, &response)
```

## Struct Embedding

Embedding allows a struct to include another struct's fields directly (composition).

### Basic Embedding

```go
type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person        // Embedded struct
    EmployeeID int
    Department string
}

// Create Employee
emp := Employee{
    Person: Person{
        Name: "Alice",
        Age:  30,
    },
    EmployeeID: 12345,
    Department: "Engineering",
}

// Access embedded fields directly
fmt.Println(emp.Name)       // "Alice" (promoted from Person)
fmt.Println(emp.Person.Name) // "Alice" (explicit)
fmt.Println(emp.Age)        // 30 (promoted)
fmt.Println(emp.EmployeeID) // 12345
```

### Method Promotion

Methods on embedded structs are promoted to the embedding struct.

```go
type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() string {
    return "Hello, I'm " + p.Name
}

type Employee struct {
    Person
    EmployeeID int
}

emp := Employee{
    Person:     Person{Name: "Alice", Age: 30},
    EmployeeID: 12345,
}

// Method promoted from Person
greeting := emp.Greet()  // "Hello, I'm Alice"
```

### Multiple Embedding

```go
type Person struct {
    Name string
}

type Address struct {
    City string
}

type Employee struct {
    Person
    Address
    EmployeeID int
}

emp := Employee{
    Person:     Person{Name: "Alice"},
    Address:    Address{City: "NYC"},
    EmployeeID: 12345,
}

// Access fields from both embedded structs
fmt.Println(emp.Name)  // "Alice"
fmt.Println(emp.City)  // "NYC"
```

### Name Conflicts

When embedded structs have same field names, use explicit path.

```go
type A struct {
    Value int
}

type B struct {
    Value int
}

type C struct {
    A
    B
}

c := C{
    A: A{Value: 1},
    B: B{Value: 2},
}

// Must be explicit (ambiguous otherwise)
fmt.Println(c.A.Value)  // 1
fmt.Println(c.B.Value)  // 2
// fmt.Println(c.Value) // Compile error: ambiguous
```

## Methods

Methods are functions associated with a struct type.

### Defining Methods

```go
type Rectangle struct {
    Width  float64
    Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// Usage
rect := Rectangle{Width: 10, Height: 5}
area := rect.Area()  // 50

rect.Scale(2)
fmt.Println(rect.Width)  // 20
```

### Method vs Function

```go
// Function
func Area(r Rectangle) float64 {
    return r.Width * r.Height
}

// Method
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Usage
rect := Rectangle{Width: 10, Height: 5}

// Function call
area := Area(rect)

// Method call
area := rect.Area()
```

### Constructor Pattern

Go doesn't have constructors, but uses factory functions.

```go
type Person struct {
    name string  // unexported
    age  int     // unexported
}

// Constructor function
func NewPerson(name string, age int) *Person {
    return &Person{
        name: name,
        age:  age,
    }
}

// Getter methods
func (p *Person) Name() string {
    return p.name
}

func (p *Person) Age() int {
    return p.age
}

// Setter methods
func (p *Person) SetAge(age int) {
    if age > 0 {
        p.age = age
    }
}

// Usage
person := NewPerson("Alice", 30)
fmt.Println(person.Name())  // "Alice"
person.SetAge(31)
```

## Pointer Receivers vs Value Receivers

### Value Receiver

Receives a copy of the struct.

```go
type Counter struct {
    count int
}

// Value receiver: operates on copy
func (c Counter) Increment() {
    c.count++  // Modifies copy, not original
}

func (c Counter) Value() int {
    return c.count
}

// Usage
counter := Counter{count: 0}
counter.Increment()
fmt.Println(counter.Value())  // 0 (unchanged!)
```

### Pointer Receiver

Receives a pointer to the struct.

```go
type Counter struct {
    count int
}

// Pointer receiver: modifies original
func (c *Counter) Increment() {
    c.count++  // Modifies original
}

func (c *Counter) Value() int {
    return c.count
}

// Usage
counter := Counter{count: 0}
counter.Increment()
fmt.Println(counter.Value())  // 1 (modified!)

// Works with pointer too
counter2 := &Counter{count: 0}
counter2.Increment()  // Go handles automatically
```

### When to Use Each

**Use pointer receivers when:**

- Method needs to modify the receiver
- Struct is large (avoid copying)
- Consistency (if any method uses pointer, all should)

**Use value receivers when:**

- Method doesn't modify the receiver
- Struct is small and copyable
- Receiver is a map, slice, or function (already reference types)

```go
type SmallStruct struct {
    x, y int
}

// Value receiver OK (small, immutable operation)
func (s SmallStruct) Sum() int {
    return s.x + s.y
}

type LargeStruct struct {
    data [1000]int
}

// Pointer receiver preferred (large struct)
func (l *LargeStruct) Process() {
    // Avoid copying 1000 integers
}
```

## Struct Tags

Struct tags are metadata attached to fields, used by packages like `json`, `xml`, `yaml`.

### Basic Syntax

```go
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email,omitempty"`
}
```

### JSON Tags

```go
type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    Email     string `json:"email,omitempty"`     // Omit if empty
    Password  string `json:"-"`                   // Never marshal
    CreatedAt string `json:"created_at,omitempty"`
}

user := User{
    ID:    1,
    Name:  "Alice",
    Email: "alice@example.com",
}

data, _ := json.Marshal(user)
// {"id":1,"name":"Alice","email":"alice@example.com"}
```

### Multiple Tags

```go
type Product struct {
    ID    int     `json:"id" xml:"id" db:"product_id"`
    Name  string  `json:"name" xml:"name" db:"product_name"`
    Price float64 `json:"price" xml:"price" db:"price"`
}
```

### Validation Tags

```go
type User struct {
    Email    string `json:"email" validate:"required,email"`
    Age      int    `json:"age" validate:"required,min=18,max=100"`
    Username string `json:"username" validate:"required,min=3,max=20"`
}
```

### Reading Tags with Reflection

```go
import "reflect"

type Person struct {
    Name string `json:"name" description:"Person's name"`
    Age  int    `json:"age" description:"Person's age"`
}

func printTags() {
    t := reflect.TypeOf(Person{})

    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        jsonTag := field.Tag.Get("json")
        desc := field.Tag.Get("description")
        fmt.Printf("%s: json=%s, desc=%s\n", field.Name, jsonTag, desc)
    }
}
```

## Struct Comparison

### Comparable Structs

Structs are comparable if all fields are comparable.

```go
type Point struct {
    X, Y int
}

p1 := Point{X: 1, Y: 2}
p2 := Point{X: 1, Y: 2}
p3 := Point{X: 3, Y: 4}

fmt.Println(p1 == p2)  // true
fmt.Println(p1 == p3)  // false
fmt.Println(p1 != p3)  // true
```

### Non-Comparable Structs

Structs with slices, maps, or functions are not comparable.

```go
type Person struct {
    Name    string
    Hobbies []string  // Slice: not comparable
}

p1 := Person{Name: "Alice", Hobbies: []string{"reading"}}
p2 := Person{Name: "Alice", Hobbies: []string{"reading"}}

// fmt.Println(p1 == p2)  // Compile error: invalid operation
```

### Custom Comparison

```go
func (p Person) Equals(other Person) bool {
    if p.Name != other.Name {
        return false
    }
    if len(p.Hobbies) != len(other.Hobbies) {
        return false
    }
    for i, hobby := range p.Hobbies {
        if hobby != other.Hobbies[i] {
            return false
        }
    }
    return true
}

// Or use reflect.DeepEqual
import "reflect"

equal := reflect.DeepEqual(p1, p2)
```

## Struct Copying

### Shallow Copy

Assignment copies structs (value type).

```go
type Person struct {
    Name string
    Age  int
}

p1 := Person{Name: "Alice", Age: 30}
p2 := p1  // Copy

p2.Name = "Bob"
fmt.Println(p1.Name)  // "Alice" (unchanged)
fmt.Println(p2.Name)  // "Bob"
```

### Shallow Copy with Pointers

```go
type Person struct {
    Name    string
    Friends []string  // Slice is a reference type
}

p1 := Person{
    Name:    "Alice",
    Friends: []string{"Bob", "Carol"},
}

p2 := p1  // Shallow copy

p2.Friends[0] = "Dave"
fmt.Println(p1.Friends[0])  // "Dave" (shared slice!)
```

### Deep Copy

```go
func (p Person) DeepCopy() Person {
    friends := make([]string, len(p.Friends))
    copy(friends, p.Friends)

    return Person{
        Name:    p.Name,
        Friends: friends,
    }
}

// Usage
p1 := Person{
    Name:    "Alice",
    Friends: []string{"Bob", "Carol"},
}

p2 := p1.DeepCopy()
p2.Friends[0] = "Dave"
fmt.Println(p1.Friends[0])  // "Bob" (independent)
```

## Common Patterns

### Builder Pattern

```go
type ServerBuilder struct {
    host string
    port int
    timeout int
}

func NewServerBuilder() *ServerBuilder {
    return &ServerBuilder{
        host: "localhost",
        port: 8080,
        timeout: 30,
    }
}

func (b *ServerBuilder) Host(host string) *ServerBuilder {
    b.host = host
    return b
}

func (b *ServerBuilder) Port(port int) *ServerBuilder {
    b.port = port
    return b
}

func (b *ServerBuilder) Timeout(timeout int) *ServerBuilder {
    b.timeout = timeout
    return b
}

func (b *ServerBuilder) Build() *Server {
    return &Server{
        host: b.host,
        port: b.port,
        timeout: b.timeout,
    }
}

// Usage
server := NewServerBuilder().
    Host("0.0.0.0").
    Port(9000).
    Timeout(60).
    Build()
```

### Options Pattern

```go
type Server struct {
    host string
    port int
}

type Option func(*Server)

func WithHost(host string) Option {
    return func(s *Server) {
        s.host = host
    }
}

func WithPort(port int) Option {
    return func(s *Server) {
        s.port = port
    }
}

func NewServer(opts ...Option) *Server {
    s := &Server{
        host: "localhost",
        port: 8080,
    }

    for _, opt := range opts {
        opt(s)
    }

    return s
}

// Usage
server := NewServer(
    WithHost("0.0.0.0"),
    WithPort(9000),
)
```

### Singleton Pattern

```go
import "sync"

type Database struct {
    connection string
}

var (
    instance *Database
    once     sync.Once
)

func GetDatabase() *Database {
    once.Do(func() {
        instance = &Database{
            connection: "connected",
        }
    })
    return instance
}
```

### State Pattern

```go
type State interface {
    Handle()
}

type Context struct {
    state State
}

func (c *Context) SetState(state State) {
    c.state = state
}

func (c *Context) Request() {
    c.state.Handle()
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
    fmt.Println("Handling in State A")
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
    fmt.Println("Handling in State B")
}
```

## Memory Layout

### Field Alignment

Go aligns struct fields for performance, which can add padding.

```go
// Not optimized: 24 bytes (with padding)
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

### Checking Size

```go
import "unsafe"

type Example struct {
    a bool
    b int64
    c bool
}

size := unsafe.Sizeof(Example{})
fmt.Println(size)  // 24 bytes (with padding)
```

### Empty Struct

Empty structs consume zero bytes.

```go
type Empty struct{}

size := unsafe.Sizeof(Empty{})  // 0 bytes

// Useful for signaling channels
done := make(chan struct{})
done <- struct{}{}

// Or sets
set := make(map[string]struct{})
set["item"] = struct{}{}
```

## Best Practices

### 1. Use Named Initialization

```go
// Good: Clear and maintainable
p := Person{
    Name: "Alice",
    Age:  30,
}

// Avoid: Order-dependent, breaks with changes
p := Person{"Alice", 30}
```

### 2. Pointer vs Value Receivers

```go
// Consistent: All methods use pointer receivers
type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
}

func (c *Counter) Value() int {
    return c.count
}
```

### 3. Use Constructor Functions

```go
// Good: Encapsulation and validation
func NewPerson(name string, age int) (*Person, error) {
    if age < 0 {
        return nil, errors.New("age cannot be negative")
    }
    return &Person{name: name, age: age}, nil
}
```

### 4. Export Selectively

```go
type person struct {  // unexported
    name string       // unexported
    age  int          // unexported
}

func NewPerson(name string, age int) *person {
    return &person{name: name, age: age}
}

func (p *person) Name() string {  // exported
    return p.name
}
```

### 5. Prefer Composition Over Inheritance

```go
// Good: Composition via embedding
type Employee struct {
    Person
    EmployeeID int
}

// Go doesn't have inheritance
```

### 6. Use Empty Struct for Signals

```go
// Memory-efficient
done := make(chan struct{})
set := make(map[string]struct{})
```

### 7. Document Exported Types

```go
// Person represents a person with a name and age.
type Person struct {
    // Name is the person's full name.
    Name string
    // Age is the person's age in years.
    Age int
}
```

### 8. Group Related Fields

```go
// Good: Logical grouping
type Server struct {
    // Network settings
    Host string
    Port int

    // Security settings
    TLS    bool
    CertFile string
    KeyFile  string

    // Timeouts
    ReadTimeout  time.Duration
    WriteTimeout time.Duration
}
```

### 9. Use Struct Tags Consistently

```go
// Good: Consistent tagging
type User struct {
    ID        int    `json:"id" db:"id"`
    Username  string `json:"username" db:"username"`
    Email     string `json:"email" db:"email"`
}
```

### 10. Consider Field Ordering for Padding

```go
// Optimized layout (largest to smallest)
type Optimized struct {
    ptr  *int    // 8 bytes
    i64  int64   // 8 bytes
    i32  int32   // 4 bytes
    i16  int16   // 2 bytes
    b    bool    // 1 byte + 1 padding
}
```

## Summary

- Structs are **value types** that group related data
- Use **struct literals** with field names for initialization
- **Embedding** allows composition and method promotion
- **Methods** can have value or pointer receivers
- Use **pointer receivers** for mutations or large structs
- **Struct tags** provide metadata for marshaling and validation
- Structs are **comparable** if all fields are comparable
- Use **constructor functions** for complex initialization
- **Field alignment** affects memory usage
- Prefer **composition** over inheritance through embedding

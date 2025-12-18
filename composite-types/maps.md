# Maps in Go

Maps are Go's built-in hash table data structure that store key-value pairs. They provide fast lookups and are essential for many programming tasks.

## Table of Contents

- [What are Maps?](#what-are-maps)
- [Declaration and Initialization](#declaration-and-initialization)
- [Map Operations](#map-operations)
- [Checking for Keys](#checking-for-keys)
- [Iterating Over Maps](#iterating-over-maps)
- [Deleting Elements](#deleting-elements)
- [Map as Function Parameters](#map-as-function-parameters)
- [Map of Maps](#map-of-maps)
- [Common Patterns](#common-patterns)
- [Concurrency and Maps](#concurrency-and-maps)
- [Performance Considerations](#performance-considerations)
- [Best Practices](#best-practices)

## What are Maps?

A map is:

- **Reference type** - Points to underlying hash table
- **Unordered** - Iteration order is random
- **Dynamic** - Can grow as needed
- **Key-value pairs** - Each key maps to exactly one value
- **Fast lookups** - O(1) average case

```go
// Map type: map[KeyType]ValueType
var ages map[string]int           // Map of string keys to int values
var scores map[int]float64         // Map of int keys to float64 values
```

## Declaration and Initialization

### Nil Map

The zero value of a map is `nil`. A nil map cannot be written to.

```go
var m map[string]int
fmt.Println(m == nil)  // true
fmt.Println(len(m))    // 0

// Reading from nil map is safe (returns zero value)
value := m["key"]      // 0

// Writing to nil map causes panic!
// m["key"] = 1        // Panic: assignment to entry in nil map
```

### Using make()

Use `make()` to create an initialized, empty map ready for use.

```go
// Create empty map
m := make(map[string]int)
fmt.Println(m == nil)  // false
fmt.Println(len(m))    // 0

// Create with initial capacity hint
m := make(map[string]int, 100)  // Capacity is a hint, not a limit
```

### Map Literal

```go
// Initialize with values
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

// Empty map (not nil)
m := map[string]int{}

// Map with struct values
type Person struct {
    Name string
    Age  int
}

people := map[string]Person{
    "emp1": {Name: "Alice", Age: 30},
    "emp2": {Name: "Bob", Age: 25},
}

// Nested map literal
config := map[string]map[string]string{
    "database": {
        "host": "localhost",
        "port": "5432",
    },
    "cache": {
        "host": "localhost",
        "port": "6379",
    },
}
```

### Allowed Key Types

Keys must be **comparable** (support `==` and `!=`).

```go
// Valid key types
map[string]int            // ✅ String keys
map[int]string            // ✅ Integer keys
map[bool]int              // ✅ Boolean keys
map[float64]string        // ✅ Float keys (use carefully!)
map[[3]int]string         // ✅ Array keys
map[struct{x, y int}]int  // ✅ Struct keys (if all fields comparable)

// Invalid key types
// map[[]int]string       // ❌ Slice keys (not comparable)
// map[map[string]int]int // ❌ Map keys (not comparable)
// map[func()]int         // ❌ Function keys (not comparable)
```

## Map Operations

### Adding and Updating

```go
m := make(map[string]int)

// Add new entry
m["Alice"] = 30

// Update existing entry (same syntax)
m["Alice"] = 31

// Add multiple entries
m["Bob"] = 25
m["Carol"] = 35
```

### Reading Values

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
}

// Get value
age := ages["Alice"]  // 30

// Key doesn't exist: returns zero value
age = ages["Dave"]    // 0 (zero value for int)
```

### The Comma-Ok Idiom

Check if a key exists with the two-value assignment.

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   0,  // Note: zero value is a valid value
}

// Two-value assignment
age, ok := ages["Alice"]
if ok {
    fmt.Printf("Alice is %d years old\n", age)
} else {
    fmt.Println("Alice not found")
}

// Distinguish between missing key and zero value
age, ok = ages["Bob"]
if ok {
    fmt.Printf("Bob is %d years old\n", age)  // "Bob is 0 years old"
} else {
    fmt.Println("Bob not found")
}

age, ok = ages["Carol"]
if ok {
    fmt.Printf("Carol is %d years old\n", age)
} else {
    fmt.Println("Carol not found")  // This executes
}
```

### Length

```go
m := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

size := len(m)  // 3
```

## Checking for Keys

### Check if Key Exists

```go
ages := map[string]int{"Alice": 30, "Bob": 25}

// Method 1: Comma-ok idiom (preferred)
if age, ok := ages["Alice"]; ok {
    fmt.Printf("Alice is %d\n", age)
}

// Method 2: Check existence only
if _, ok := ages["Carol"]; ok {
    fmt.Println("Carol exists")
} else {
    fmt.Println("Carol not found")
}

// Method 3: Using value (only if zero value means "not present")
age := ages["Alice"]
if age != 0 {  // Unsafe if 0 is a valid value!
    fmt.Println("Found")
}
```

### Get with Default Value

```go
func getOrDefault(m map[string]int, key string, defaultValue int) int {
    if value, ok := m[key]; ok {
        return value
    }
    return defaultValue
}

// Usage
ages := map[string]int{"Alice": 30}
age := getOrDefault(ages, "Bob", 18)  // Returns 18
```

## Iterating Over Maps

### Range Loop

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

// Iterate over key-value pairs
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}

// Iterate over keys only
for name := range ages {
    fmt.Println(name)
}

// Iterate over values only
for _, age := range ages {
    fmt.Println(age)
}
```

### Iteration Order

**Maps are unordered!** Iteration order is randomized intentionally.

```go
m := map[string]int{"a": 1, "b": 2, "c": 3}

// Different order each time!
for k, v := range m {
    fmt.Printf("%s: %d\n", k, v)
}
```

### Sorted Iteration

To iterate in order, sort the keys first.

```go
import "sort"

ages := map[string]int{
    "Carol": 35,
    "Alice": 30,
    "Bob":   25,
}

// Get keys
keys := make([]string, 0, len(ages))
for k := range ages {
    keys = append(keys, k)
}

// Sort keys
sort.Strings(keys)

// Iterate in sorted order
for _, name := range keys {
    fmt.Printf("%s: %d\n", name, ages[name])
}
// Output:
// Alice: 30
// Bob: 25
// Carol: 35
```

## Deleting Elements

### Using delete()

```go
ages := map[string]int{
    "Alice": 30,
    "Bob":   25,
    "Carol": 35,
}

// Delete entry
delete(ages, "Bob")
fmt.Println(len(ages))  // 2

// Delete non-existent key (safe, no-op)
delete(ages, "Dave")    // No error
```

### Clear All Elements

```go
// Method 1: Assign new empty map
ages = make(map[string]int)

// Method 2: Delete all keys
for k := range ages {
    delete(ages, k)
}

// Method 3: In Go 1.21+, use clear()
clear(ages)
```

## Map as Function Parameters

### Maps are Reference Types

Maps are passed by reference, so modifications affect the original.

```go
func addPerson(people map[string]int, name string, age int) {
    people[name] = age  // Modifies original map
}

func main() {
    ages := map[string]int{"Alice": 30}
    addPerson(ages, "Bob", 25)
    fmt.Println(ages)  // map[Alice:30 Bob:25] - modified
}
```

### Returning Maps

```go
func createAgeMap() map[string]int {
    return map[string]int{
        "Alice": 30,
        "Bob":   25,
    }
}

func filterByAge(people map[string]int, minAge int) map[string]int {
    result := make(map[string]int)
    for name, age := range people {
        if age >= minAge {
            result[name] = age
        }
    }
    return result
}
```

## Map of Maps

### Nested Maps

```go
// Map of maps
userPrefs := make(map[string]map[string]string)

// Initialize inner map before use!
userPrefs["alice"] = make(map[string]string)
userPrefs["alice"]["theme"] = "dark"
userPrefs["alice"]["language"] = "en"

// Or use literal
userPrefs := map[string]map[string]string{
    "alice": {
        "theme":    "dark",
        "language": "en",
    },
    "bob": {
        "theme":    "light",
        "language": "es",
    },
}

// Access nested values
theme := userPrefs["alice"]["theme"]  // "dark"

// Check nested key safely
if user, ok := userPrefs["carol"]; ok {
    if theme, ok := user["theme"]; ok {
        fmt.Println(theme)
    }
}
```

### Helper for Nested Maps

```go
func setNested(m map[string]map[string]string, key1, key2, value string) {
    if m[key1] == nil {
        m[key1] = make(map[string]string)
    }
    m[key1][key2] = value
}

// Usage
userPrefs := make(map[string]map[string]string)
setNested(userPrefs, "alice", "theme", "dark")
```

## Common Patterns

### Set Using Map

```go
// Set implementation using map[T]bool
type Set map[string]bool

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(item string) {
    s[item] = true
}

func (s Set) Remove(item string) {
    delete(s, item)
}

func (s Set) Contains(item string) bool {
    return s[item]
}

func (s Set) Size() int {
    return len(s)
}

// Usage
fruits := NewSet()
fruits.Add("apple")
fruits.Add("banana")
fruits.Add("apple")  // Duplicate, no effect
fmt.Println(fruits.Size())  // 2
fmt.Println(fruits.Contains("apple"))  // true
```

### Using map[T]struct{} for Sets

More memory-efficient than `map[T]bool`.

```go
type Set map[string]struct{}

func (s Set) Add(item string) {
    s[item] = struct{}{}
}

func (s Set) Contains(item string) bool {
    _, ok := s[item]
    return ok
}
```

### Counting Occurrences

```go
func countWords(text string) map[string]int {
    words := strings.Fields(text)
    counts := make(map[string]int)

    for _, word := range words {
        counts[word]++  // Automatically initializes to 0 if not exists
    }

    return counts
}

// Usage
text := "hello world hello go world"
counts := countWords(text)
fmt.Println(counts)  // map[go:1 hello:2 world:2]
```

### Grouping Data

```go
type Person struct {
    Name string
    City string
    Age  int
}

func groupByCity(people []Person) map[string][]Person {
    groups := make(map[string][]Person)

    for _, person := range people {
        groups[person.City] = append(groups[person.City], person)
    }

    return groups
}

// Usage
people := []Person{
    {"Alice", "NYC", 30},
    {"Bob", "LA", 25},
    {"Carol", "NYC", 35},
}

byCity := groupByCity(people)
fmt.Println(byCity["NYC"])  // [{Alice NYC 30} {Carol NYC 35}]
```

### Inverse Map

```go
func invertMap(m map[string]int) map[int]string {
    inverse := make(map[int]string)
    for k, v := range m {
        inverse[v] = k  // Last key wins if values duplicate
    }
    return inverse
}

// Handle duplicates
func invertMapMulti(m map[string]int) map[int][]string {
    inverse := make(map[int][]string)
    for k, v := range m {
        inverse[v] = append(inverse[v], k)
    }
    return inverse
}
```

### Memoization/Caching

```go
type Fibonacci struct {
    cache map[int]int
}

func NewFibonacci() *Fibonacci {
    return &Fibonacci{
        cache: make(map[int]int),
    }
}

func (f *Fibonacci) Compute(n int) int {
    // Check cache first
    if result, ok := f.cache[n]; ok {
        return result
    }

    // Compute
    var result int
    if n <= 1 {
        result = n
    } else {
        result = f.Compute(n-1) + f.Compute(n-2)
    }

    // Store in cache
    f.cache[n] = result
    return result
}

// Usage
fib := NewFibonacci()
fmt.Println(fib.Compute(50))  // Fast due to caching
```

### Default Values with Map

```go
type Config struct {
    defaults map[string]string
    values   map[string]string
}

func NewConfig() *Config {
    return &Config{
        defaults: map[string]string{
            "host": "localhost",
            "port": "8080",
        },
        values: make(map[string]string),
    }
}

func (c *Config) Get(key string) string {
    if value, ok := c.values[key]; ok {
        return value
    }
    return c.defaults[key]
}

func (c *Config) Set(key, value string) {
    c.values[key] = value
}
```

### Merging Maps

```go
func mergeMaps(maps ...map[string]int) map[string]int {
    result := make(map[string]int)
    for _, m := range maps {
        for k, v := range m {
            result[k] = v  // Last value wins
        }
    }
    return result
}

// Usage
m1 := map[string]int{"a": 1, "b": 2}
m2 := map[string]int{"b": 3, "c": 4}
merged := mergeMaps(m1, m2)  // map[a:1 b:3 c:4]
```

## Concurrency and Maps

### Maps are Not Thread-Safe

```go
// UNSAFE: Concurrent access without synchronization
m := make(map[string]int)

go func() {
    m["key"] = 1  // Write
}()

go func() {
    _ = m["key"]  // Read
}()
// This can cause race conditions and panics!
```

### Using sync.Mutex

```go
import "sync"

type SafeMap struct {
    mu sync.Mutex
    m  map[string]int
}

func NewSafeMap() *SafeMap {
    return &SafeMap{
        m: make(map[string]int),
    }
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.m[key] = value
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    val, ok := sm.m[key]
    return val, ok
}

func (sm *SafeMap) Delete(key string) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    delete(sm.m, key)
}
```

### Using sync.RWMutex

```go
import "sync"

type SafeMap struct {
    mu sync.RWMutex
    m  map[string]int
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()  // Read lock
    defer sm.mu.RUnlock()
    val, ok := sm.m[key]
    return val, ok
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()  // Write lock
    defer sm.mu.Unlock()
    sm.m[key] = value
}
```

### Using sync.Map

Go provides `sync.Map` for concurrent use without explicit locking.

```go
import "sync"

var m sync.Map

// Store
m.Store("key", "value")

// Load
value, ok := m.Load("key")
if ok {
    fmt.Println(value)
}

// LoadOrStore
actual, loaded := m.LoadOrStore("key", "newvalue")
// loaded is true if key existed

// Delete
m.Delete("key")

// Range
m.Range(func(key, value interface{}) bool {
    fmt.Printf("%v: %v\n", key, value)
    return true  // Continue iteration
})
```

**Note:** `sync.Map` is optimized for specific use cases:

- Keys are written once but read many times
- Multiple goroutines read/write/overwrite disjoint sets of keys

## Performance Considerations

### Pre-allocate Capacity

```go
// Bad: Multiple reallocations
m := make(map[string]int)
for i := 0; i < 10000; i++ {
    m[fmt.Sprintf("key%d", i)] = i
}

// Good: Pre-allocate
m := make(map[string]int, 10000)
for i := 0; i < 10000; i++ {
    m[fmt.Sprintf("key%d", i)] = i
}
```

### Avoid Float Keys

Floating-point keys can cause issues due to precision.

```go
// Problematic
m := make(map[float64]string)
m[0.1+0.2] = "a"
m[0.3] = "b"
// These might not be equal due to float precision!
```

### Delete vs New Map

For clearing large maps, creating a new map can be faster than deleting all keys.

```go
// For small maps
for k := range m {
    delete(m, k)
}

// For large maps (helps GC)
m = make(map[string]int)
```

### Key Type Performance

```go
// Fast: Integer keys
map[int]string

// Fast: String keys (for reasonable-length strings)
map[string]int

// Slower: Large struct keys
map[LargeStruct]int

// Consider: Use pointer or hash
type Key struct {
    // Large fields
}
map[*Key]int  // Compare pointers, not values
```

## Best Practices

### 1. Initialize Before Use

```go
// Bad: Nil map
var m map[string]int
// m["key"] = 1  // Panic!

// Good: Initialize with make or literal
m := make(map[string]int)
m["key"] = 1  // OK
```

### 2. Use Comma-Ok for Existence Check

```go
// Good: Clear existence check
if value, ok := m["key"]; ok {
    // Use value
}

// Bad: Ambiguous with zero values
value := m["key"]
if value != 0 {  // What if 0 is valid?
    // Use value
}
```

### 3. Delete in Loop is Safe

```go
// Safe: Delete while iterating
for k, v := range m {
    if someCondition(v) {
        delete(m, k)  // OK
    }
}
```

### 4. Don't Rely on Iteration Order

```go
// Bad: Assuming order
for k, v := range m {
    // Expecting consistent order - WRONG!
}

// Good: Sort keys if order matters
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    // Process in order
}
```

### 5. Use Struct Keys Carefully

```go
// Good: Simple comparable struct
type Point struct {
    X, Y int
}
m := make(map[Point]string)

// Bad: Struct with uncomparable fields
type Bad struct {
    Data []int  // Slice is not comparable
}
// m := make(map[Bad]string)  // Won't compile
```

### 6. Pre-allocate for Known Size

```go
// Good: Known approximate size
m := make(map[string]int, 1000)

// Avoid: Unknown size, start with 0
m := make(map[string]int)
```

### 7. Check Nil Before Operations

```go
func processMap(m map[string]int) {
    if m == nil {
        return  // or handle appropriately
    }
    // Safe to use m
}
```

### 8. Use sync.Map for Concurrent Access

```go
// Bad: Regular map with concurrent access
m := make(map[string]int)
// Multiple goroutines reading/writing - RACE!

// Good: Use sync.Map or add mutex
var m sync.Map
// Or
type SafeMap struct {
    mu sync.RWMutex
    m  map[string]int
}
```

### 9. Consider Memory for Large Maps

```go
// For millions of entries, consider:
// - Pre-allocation
// - Smaller key/value types
// - Periodic cleanup
// - Alternative data structures (B-tree, trie, etc.)
```

### 10. Document Key Guarantees

```go
// Good: Clear documentation
// UserCache maps user ID to user data.
// Keys are unique user identifiers.
// Safe for concurrent read-only access.
type UserCache map[int]*User
```

## Summary

- Maps are **reference types** storing key-value pairs
- Maps must be **initialized** before use (nil maps panic on write)
- Use **make()** or **map literals** to initialize
- **Comma-ok idiom** to check if key exists
- Maps are **unordered** - iteration order is random
- Use **delete()** to remove entries
- Maps are **not thread-safe** - use mutex or sync.Map
- **Pre-allocate capacity** when size is known
- Keys must be **comparable** types
- Maps are passed by **reference** to functions
- Use `map[T]struct{}` for **memory-efficient sets**

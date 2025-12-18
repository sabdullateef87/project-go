# JSON in Go

Go provides the `encoding/json` package for encoding and decoding JSON data. It's essential for web APIs, configuration files, and data serialization.

## Table of Contents

- [Overview](#overview)
- [Marshaling (Encoding)](#marshaling-encoding)
- [Unmarshaling (Decoding)](#unmarshaling-decoding)
- [Struct Tags](#struct-tags)
- [Working with Different JSON Structures](#working-with-different-json-structures)
- [Custom Marshaling](#custom-marshaling)
- [Custom Unmarshaling](#custom-unmarshaling)
- [Streaming JSON](#streaming-json)
- [Common Patterns](#common-patterns)
- [Error Handling](#error-handling)
- [Performance Considerations](#performance-considerations)
- [Best Practices](#best-practices)

## Overview

The `encoding/json` package provides:

- **Marshal** - Convert Go values to JSON
- **Unmarshal** - Parse JSON into Go values
- **Encoder/Decoder** - Streaming JSON processing
- **Struct tags** - Control JSON field names and behavior
- **Custom marshaling** - Implement custom JSON representation

```go
import "encoding/json"
```

## Marshaling (Encoding)

Converting Go values to JSON.

### Basic Types

```go
// Numbers
num := 42
data, _ := json.Marshal(num)
fmt.Println(string(data))  // "42"

// Strings
str := "Hello, World"
data, _ = json.Marshal(str)
fmt.Println(string(data))  // "\"Hello, World\""

// Booleans
flag := true
data, _ = json.Marshal(flag)
fmt.Println(string(data))  // "true"

// nil
data, _ = json.Marshal(nil)
fmt.Println(string(data))  // "null"
```

### Arrays and Slices

```go
// Array
arr := [3]int{1, 2, 3}
data, _ := json.Marshal(arr)
fmt.Println(string(data))  // "[1,2,3]"

// Slice
slice := []string{"apple", "banana", "cherry"}
data, _ = json.Marshal(slice)
fmt.Println(string(data))  // "[\"apple\",\"banana\",\"cherry\"]"

// Nil slice
var nilSlice []int
data, _ = json.Marshal(nilSlice)
fmt.Println(string(data))  // "null"

// Empty slice
emptySlice := []int{}
data, _ = json.Marshal(emptySlice)
fmt.Println(string(data))  // "[]"
```

### Maps

```go
m := map[string]int{
    "age":   30,
    "score": 95,
}

data, _ := json.Marshal(m)
fmt.Println(string(data))  // {"age":30,"score":95}
```

### Structs

```go
type Person struct {
    Name string
    Age  int
    City string
}

person := Person{
    Name: "Alice",
    Age:  30,
    City: "NYC",
}

data, err := json.Marshal(person)
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(data))
// {"Name":"Alice","Age":30,"City":"NYC"}
```

### Pretty Printing

```go
person := Person{Name: "Alice", Age: 30, City: "NYC"}

// MarshalIndent for formatted output
data, err := json.MarshalIndent(person, "", "  ")
if err != nil {
    log.Fatal(err)
}

fmt.Println(string(data))
// {
//   "Name": "Alice",
//   "Age": 30,
//   "City": "NYC"
// }
```

### Marshal Return Values

```go
data, err := json.Marshal(value)

// data is []byte
// err is error (nil if successful)
```

## Unmarshaling (Decoding)

Parsing JSON into Go values.

### Basic Types

```go
// Numbers
var num int
json.Unmarshal([]byte("42"), &num)
fmt.Println(num)  // 42

// Strings
var str string
json.Unmarshal([]byte("\"Hello\""), &str)
fmt.Println(str)  // "Hello"

// Booleans
var flag bool
json.Unmarshal([]byte("true"), &flag)
fmt.Println(flag)  // true
```

### Arrays and Slices

```go
// Into slice
var slice []int
json.Unmarshal([]byte("[1,2,3]"), &slice)
fmt.Println(slice)  // [1 2 3]

// Into array
var arr [3]int
json.Unmarshal([]byte("[1,2,3]"), &arr)
fmt.Println(arr)  // [1 2 3]
```

### Maps

```go
var m map[string]int
json.Unmarshal([]byte(`{"age":30,"score":95}`), &m)
fmt.Println(m)  // map[age:30 score:95]
```

### Structs

```go
type Person struct {
    Name string
    Age  int
    City string
}

jsonData := []byte(`{"Name":"Alice","Age":30,"City":"NYC"}`)

var person Person
err := json.Unmarshal(jsonData, &person)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("%+v\n", person)
// {Name:Alice Age:30 City:NYC}
```

### Unknown Structure (interface{})

```go
jsonData := []byte(`{"name":"Alice","age":30,"active":true}`)

var data interface{}
json.Unmarshal(jsonData, &data)

// Type assertion required
m := data.(map[string]interface{})
name := m["name"].(string)
age := m["age"].(float64)  // Numbers are float64
active := m["active"].(bool)

fmt.Println(name, age, active)  // Alice 30 true
```

## Struct Tags

Control how struct fields are encoded/decoded.

### Basic Tags

```go
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age"`
    Email string `json:"email"`
}

person := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
data, _ := json.Marshal(person)
fmt.Println(string(data))
// {"name":"Alice","age":30,"email":"alice@example.com"}
```

### Omit Empty Fields

```go
type Person struct {
    Name  string `json:"name"`
    Age   int    `json:"age,omitempty"`
    Email string `json:"email,omitempty"`
}

person := Person{Name: "Alice"}
data, _ := json.Marshal(person)
fmt.Println(string(data))
// {"name":"Alice"}
// Age and Email omitted because they're zero values
```

### Ignore Fields

```go
type User struct {
    Username string `json:"username"`
    Password string `json:"-"`  // Never marshal/unmarshal
    Email    string `json:"email"`
}

user := User{
    Username: "alice",
    Password: "secret123",
    Email:    "alice@example.com",
}

data, _ := json.Marshal(user)
fmt.Println(string(data))
// {"username":"alice","email":"alice@example.com"}
// Password not included
```

### String Option

Forces encoding as string (useful for numbers).

```go
type Product struct {
    ID    int     `json:"id,string"`
    Price float64 `json:"price,string"`
}

product := Product{ID: 123, Price: 19.99}
data, _ := json.Marshal(product)
fmt.Println(string(data))
// {"id":"123","price":"19.99"}
```

### Embedded Structs

```go
type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

type Person struct {
    Name    string  `json:"name"`
    Address Address `json:"address"`
}

person := Person{
    Name: "Alice",
    Address: Address{
        Street: "123 Main St",
        City:   "NYC",
    },
}

data, _ := json.Marshal(person)
fmt.Println(string(data))
// {"name":"Alice","address":{"street":"123 Main St","city":"NYC"}}
```

### Inline/Flatten Embedded Structs

```go
type Address struct {
    Street string `json:"street"`
    City   string `json:"city"`
}

type Person struct {
    Name    string `json:"name"`
    Address `json:"-"`  // Ignore the field itself
    Street  string `json:"street"`  // Promote manually
    City    string `json:"city"`    // Promote manually
}

// Or use anonymous embedding without tag:
type Person2 struct {
    Name string `json:"name"`
    Address     // Fields promoted to same level
}

person := Person2{
    Name: "Alice",
    Address: Address{
        Street: "123 Main St",
        City:   "NYC",
    },
}

data, _ := json.Marshal(person)
// {"name":"Alice","street":"123 Main St","city":"NYC"}
```

## Working with Different JSON Structures

### Nested Objects

```go
type Company struct {
    Name    string `json:"name"`
    Address struct {
        Street  string `json:"street"`
        City    string `json:"city"`
        Country string `json:"country"`
    } `json:"address"`
}

jsonData := []byte(`{
    "name": "TechCorp",
    "address": {
        "street": "123 Tech Lane",
        "city": "San Francisco",
        "country": "USA"
    }
}`)

var company Company
json.Unmarshal(jsonData, &company)
fmt.Printf("%+v\n", company)
```

### Arrays of Objects

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

jsonData := []byte(`[
    {"name":"Alice","age":30},
    {"name":"Bob","age":25}
]`)

var people []Person
json.Unmarshal(jsonData, &people)
fmt.Printf("%+v\n", people)
// [{Name:Alice Age:30} {Name:Bob Age:25}]
```

### Optional Fields

```go
type Config struct {
    Host    string  `json:"host"`
    Port    int     `json:"port,omitempty"`
    Timeout *int    `json:"timeout,omitempty"`  // Pointer distinguishes nil from 0
    Debug   bool    `json:"debug,omitempty"`
}

// With timeout
timeout := 30
config1 := Config{Host: "localhost", Timeout: &timeout}

// Without timeout (will be omitted)
config2 := Config{Host: "localhost"}
```

### Dynamic Keys

```go
jsonData := []byte(`{
    "user123": {"name": "Alice", "age": 30},
    "user456": {"name": "Bob", "age": 25}
}`)

type User struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

var users map[string]User
json.Unmarshal(jsonData, &users)

for id, user := range users {
    fmt.Printf("%s: %+v\n", id, user)
}
// user123: {Name:Alice Age:30}
// user456: {Name:Bob Age:25}
```

### Union Types (Multiple Possible Types)

```go
type Response struct {
    Status string          `json:"status"`
    Data   json.RawMessage `json:"data"`  // Defer parsing
}

jsonData := []byte(`{"status":"success","data":{"name":"Alice","age":30}}`)

var resp Response
json.Unmarshal(jsonData, &resp)

// Parse data based on status
if resp.Status == "success" {
    var person Person
    json.Unmarshal(resp.Data, &person)
    fmt.Printf("%+v\n", person)
}
```

## Custom Marshaling

Implement `json.Marshaler` interface.

### Basic Custom Marshaler

```go
type Person struct {
    Name string
    Age  int
}

// Implement MarshalJSON
func (p Person) MarshalJSON() ([]byte, error) {
    // Custom JSON representation
    return json.Marshal(struct {
        FullName string `json:"full_name"`
        YearsOld int    `json:"years_old"`
    }{
        FullName: p.Name,
        YearsOld: p.Age,
    })
}

person := Person{Name: "Alice", Age: 30}
data, _ := json.Marshal(person)
fmt.Println(string(data))
// {"full_name":"Alice","years_old":30}
```

### Custom Time Format

```go
type CustomTime time.Time

func (ct CustomTime) MarshalJSON() ([]byte, error) {
    t := time.Time(ct)
    formatted := t.Format("2006-01-02")
    return json.Marshal(formatted)
}

type Event struct {
    Name string     `json:"name"`
    Date CustomTime `json:"date"`
}

event := Event{
    Name: "Conference",
    Date: CustomTime(time.Now()),
}

data, _ := json.Marshal(event)
fmt.Println(string(data))
// {"name":"Conference","date":"2025-12-18"}
```

### Pointer Receiver

```go
type Person struct {
    Name string
    Age  int
}

// Use pointer receiver to modify during marshal
func (p *Person) MarshalJSON() ([]byte, error) {
    type Alias Person  // Prevent recursion
    return json.Marshal(&struct {
        *Alias
        IsAdult bool `json:"is_adult"`
    }{
        Alias:   (*Alias)(p),
        IsAdult: p.Age >= 18,
    })
}

person := Person{Name: "Alice", Age: 30}
data, _ := json.Marshal(&person)
fmt.Println(string(data))
// {"Name":"Alice","Age":30,"is_adult":true}
```

## Custom Unmarshaling

Implement `json.Unmarshaler` interface.

### Basic Custom Unmarshaler

```go
type Person struct {
    Name string
    Age  int
}

func (p *Person) UnmarshalJSON(data []byte) error {
    // Custom parsing
    var aux struct {
        FullName string `json:"full_name"`
        YearsOld int    `json:"years_old"`
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    p.Name = aux.FullName
    p.Age = aux.YearsOld
    return nil
}

jsonData := []byte(`{"full_name":"Alice","years_old":30}`)

var person Person
json.Unmarshal(jsonData, &person)
fmt.Printf("%+v\n", person)
// {Name:Alice Age:30}
```

### Flexible Input Formats

```go
type Duration time.Duration

func (d *Duration) UnmarshalJSON(data []byte) error {
    var v interface{}
    if err := json.Unmarshal(data, &v); err != nil {
        return err
    }

    switch value := v.(type) {
    case float64:
        *d = Duration(time.Duration(value) * time.Second)
    case string:
        parsed, err := time.ParseDuration(value)
        if err != nil {
            return err
        }
        *d = Duration(parsed)
    default:
        return errors.New("invalid duration")
    }

    return nil
}

// Can accept "5s" or 5
jsonData1 := []byte(`{"timeout":"5s"}`)
jsonData2 := []byte(`{"timeout":5}`)
```

### Validation During Unmarshal

```go
type Email string

func (e *Email) UnmarshalJSON(data []byte) error {
    var s string
    if err := json.Unmarshal(data, &s); err != nil {
        return err
    }

    // Validate email format
    if !strings.Contains(s, "@") {
        return errors.New("invalid email format")
    }

    *e = Email(s)
    return nil
}

type User struct {
    Name  string `json:"name"`
    Email Email  `json:"email"`
}
```

## Streaming JSON

Use `Encoder` and `Decoder` for streaming.

### Encoding to Stream

```go
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

// Encode to io.Writer
var buf bytes.Buffer
encoder := json.NewEncoder(&buf)

people := []Person{
    {Name: "Alice", Age: 30},
    {Name: "Bob", Age: 25},
}

for _, person := range people {
    if err := encoder.Encode(person); err != nil {
        log.Fatal(err)
    }
}

fmt.Println(buf.String())
// {"name":"Alice","age":30}
// {"name":"Bob","age":25}
```

### Decoding from Stream

```go
jsonData := `
{"name":"Alice","age":30}
{"name":"Bob","age":25}
{"name":"Carol","age":35}
`

decoder := json.NewDecoder(strings.NewReader(jsonData))

for {
    var person Person
    if err := decoder.Decode(&person); err == io.EOF {
        break
    } else if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%+v\n", person)
}
```

### HTTP Response Streaming

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    encoder := json.NewEncoder(w)

    result := map[string]interface{}{
        "status": "success",
        "data":   []string{"item1", "item2"},
    }

    if err := encoder.Encode(result); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
```

### Reading HTTP Response

```go
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

var users []User
decoder := json.NewDecoder(resp.Body)
if err := decoder.Decode(&users); err != nil {
    log.Fatal(err)
}
```

## Common Patterns

### API Response Wrapper

```go
type APIResponse struct {
    Success bool        `json:"success"`
    Data    interface{} `json:"data,omitempty"`
    Error   string      `json:"error,omitempty"`
}

func SuccessResponse(data interface{}) APIResponse {
    return APIResponse{
        Success: true,
        Data:    data,
    }
}

func ErrorResponse(err string) APIResponse {
    return APIResponse{
        Success: false,
        Error:   err,
    }
}

// Usage
data, _ := json.Marshal(SuccessResponse(user))
```

### Pagination

```go
type PaginatedResponse struct {
    Data       interface{} `json:"data"`
    Page       int         `json:"page"`
    PerPage    int         `json:"per_page"`
    Total      int         `json:"total"`
    TotalPages int         `json:"total_pages"`
}

func Paginate(items []User, page, perPage int) PaginatedResponse {
    start := (page - 1) * perPage
    end := start + perPage

    if end > len(items) {
        end = len(items)
    }

    return PaginatedResponse{
        Data:       items[start:end],
        Page:       page,
        PerPage:    perPage,
        Total:      len(items),
        TotalPages: (len(items) + perPage - 1) / perPage,
    }
}
```

### Enum/Constant Values

```go
type Status int

const (
    StatusPending Status = iota
    StatusActive
    StatusInactive
)

func (s Status) MarshalJSON() ([]byte, error) {
    names := []string{"pending", "active", "inactive"}
    if s < 0 || int(s) >= len(names) {
        return nil, errors.New("invalid status")
    }
    return json.Marshal(names[s])
}

func (s *Status) UnmarshalJSON(data []byte) error {
    var name string
    if err := json.Unmarshal(data, &name); err != nil {
        return err
    }

    switch name {
    case "pending":
        *s = StatusPending
    case "active":
        *s = StatusActive
    case "inactive":
        *s = StatusInactive
    default:
        return errors.New("invalid status")
    }

    return nil
}
```

### Partial Updates

```go
type User struct {
    ID    int     `json:"id"`
    Name  *string `json:"name,omitempty"`
    Email *string `json:"email,omitempty"`
    Age   *int    `json:"age,omitempty"`
}

// Only updates provided fields
func UpdateUser(id int, updates User) error {
    if updates.Name != nil {
        // Update name
    }
    if updates.Email != nil {
        // Update email
    }
    if updates.Age != nil {
        // Update age
    }
    return nil
}
```

### Converting Between Types

```go
// Convert struct to map
func StructToMap(obj interface{}) (map[string]interface{}, error) {
    data, err := json.Marshal(obj)
    if err != nil {
        return nil, err
    }

    var m map[string]interface{}
    err = json.Unmarshal(data, &m)
    return m, err
}

// Convert map to struct
func MapToStruct(m map[string]interface{}, result interface{}) error {
    data, err := json.Marshal(m)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, result)
}
```

## Error Handling

### Common Errors

```go
// SyntaxError: Invalid JSON
jsonData := []byte(`{"name": "Alice"`) // Missing }
var person Person
err := json.Unmarshal(jsonData, &person)
if syntaxErr, ok := err.(*json.SyntaxError); ok {
    fmt.Printf("Syntax error at byte %d\n", syntaxErr.Offset)
}

// UnmarshalTypeError: Type mismatch
jsonData = []byte(`{"name":"Alice","age":"thirty"}`)
err = json.Unmarshal(jsonData, &person)
if typeErr, ok := err.(*json.UnmarshalTypeError); ok {
    fmt.Printf("Type error: expected %s, got %s for field %s\n",
        typeErr.Type, typeErr.Value, typeErr.Field)
}
```

### Validation After Unmarshal

```go
type User struct {
    Username string `json:"username"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
}

func (u *User) Validate() error {
    if u.Username == "" {
        return errors.New("username is required")
    }
    if !strings.Contains(u.Email, "@") {
        return errors.New("invalid email")
    }
    if u.Age < 0 || u.Age > 150 {
        return errors.New("invalid age")
    }
    return nil
}

// Usage
var user User
if err := json.Unmarshal(jsonData, &user); err != nil {
    return err
}
if err := user.Validate(); err != nil {
    return err
}
```

### Safe Unmarshaling

```go
func SafeUnmarshal(data []byte, v interface{}) error {
    decoder := json.NewDecoder(bytes.NewReader(data))
    decoder.DisallowUnknownFields()  // Strict mode

    if err := decoder.Decode(v); err != nil {
        return fmt.Errorf("json decode error: %w", err)
    }

    return nil
}
```

## Performance Considerations

### Reuse Buffers

```go
// Pool for byte buffers
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func MarshalWithPool(v interface{}) ([]byte, error) {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)

    encoder := json.NewEncoder(buf)
    if err := encoder.Encode(v); err != nil {
        return nil, err
    }

    // Copy because buffer will be reused
    result := make([]byte, buf.Len())
    copy(result, buf.Bytes())
    return result, nil
}
```

### Use json.RawMessage

```go
type Response struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`  // Delay parsing
}

// Parse only when needed
if response.Type == "user" {
    var user User
    json.Unmarshal(response.Data, &user)
}
```

### Streaming for Large Data

```go
// Don't load entire array into memory
func ProcessLargeJSON(r io.Reader) error {
    decoder := json.NewDecoder(r)

    // Read opening bracket
    _, err := decoder.Token()
    if err != nil {
        return err
    }

    // Process items one by one
    for decoder.More() {
        var item Item
        if err := decoder.Decode(&item); err != nil {
            return err
        }
        processItem(item)
    }

    return nil
}
```

## Best Practices

### 1. Always Check Errors

```go
// Good
data, err := json.Marshal(value)
if err != nil {
    return fmt.Errorf("marshal error: %w", err)
}

// Bad
data, _ := json.Marshal(value)
```

### 2. Use Struct Tags

```go
// Good: Clear field mappings
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
}

// Bad: Relies on struct field names
type User struct {
    ID       int
    Username string
}
```

### 3. Omit Empty for Optional Fields

```go
type Config struct {
    Required string  `json:"required"`
    Optional string  `json:"optional,omitempty"`
    Debug    bool    `json:"debug,omitempty"`
}
```

### 4. Use Pointers for Nullable Fields

```go
// Distinguishes between zero value and null
type User struct {
    Name  string `json:"name"`
    Age   *int   `json:"age,omitempty"`  // nil = not provided
    Score *int   `json:"score,omitempty"` // 0 is valid
}
```

### 5. Validate After Unmarshal

```go
var user User
if err := json.Unmarshal(data, &user); err != nil {
    return err
}
if err := user.Validate(); err != nil {
    return err
}
```

### 6. Use json.RawMessage for Deferred Parsing

```go
type Envelope struct {
    Type string          `json:"type"`
    Data json.RawMessage `json:"data"`
}

// Parse data based on type
```

### 7. Export Only What's Needed

```go
type User struct {
    id       int    // unexported, not marshaled
    password string // unexported, not marshaled
    Username string `json:"username"` // exported, marshaled
}
```

### 8. Document JSON Format

```go
// User represents a user account.
// JSON format:
//   {
//     "id": 123,
//     "username": "alice",
//     "email": "alice@example.com"
//   }
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
}
```

### 9. Use Encoder/Decoder for Streams

```go
// Good: Streaming
encoder := json.NewEncoder(w)
encoder.Encode(data)

// Less efficient: Marshal to memory first
data, _ := json.Marshal(data)
w.Write(data)
```

### 10. Handle Time Properly

```go
type Event struct {
    Name      string    `json:"name"`
    Timestamp time.Time `json:"timestamp"`
}

// time.Time marshals to RFC3339 by default
// "2025-12-18T10:30:00Z"
```

## Summary

- Use **json.Marshal** to encode Go values to JSON
- Use **json.Unmarshal** to decode JSON into Go values
- **Struct tags** control field names and behavior (`json:"name,omitempty"`)
- Implement **MarshalJSON/UnmarshalJSON** for custom encoding
- Use **Encoder/Decoder** for streaming large data
- Use **json.RawMessage** to defer parsing
- **Pointer fields** distinguish between zero values and null
- Always **validate** after unmarshaling
- **Handle errors** properly in production code
- Use **omitempty** for optional fields

# Interfaces in Go

Interfaces are one of Go's most powerful features, enabling polymorphism and flexible, decoupled design. An interface defines a contract that types can satisfy by implementing its methods. This guide covers everything from basic interface concepts to advanced patterns.

## Table of Contents

1. [Basic Interface Concepts](#basic-interface-concepts)
2. [Interface Declaration and Implementation](#interface-declaration-and-implementation)
3. [Empty Interface](#empty-interface)
4. [Type Assertions and Type Switches](#type-assertions-and-type-switches)
5. [Interface Embedding](#interface-embedding)
6. [Common Standard Library Interfaces](#common-standard-library-interfaces)
7. [Interface Best Practices](#interface-best-practices)
8. [Advanced Interface Patterns](#advanced-interface-patterns)
9. [Interface Examples and Use Cases](#interface-examples-and-use-cases)

## Basic Interface Concepts

### What is an Interface?

An interface is a type that specifies a set of method signatures. Any type that implements all the methods of an interface automatically satisfies that interface - this is called **implicit implementation**.

```go
// Define an interface
type Writer interface {
    Write([]byte) (int, error)
}

// Any type with a Write method satisfies the Writer interface
type FileWriter struct {
    filename string
}

func (fw FileWriter) Write(data []byte) (int, error) {
    // Implementation here
    return len(data), nil
}
```

### Key Characteristics

| Feature | Description |
| ------- | ----------- |
| **Implicit Implementation** | No explicit declaration needed |
| **Duck Typing** | "If it walks like a duck and quacks like a duck, it's a duck" |
| **Zero Value** | `nil` |
| **Composition** | Interfaces can embed other interfaces |
| **Polymorphism** | Different types can satisfy the same interface |

## Interface Declaration and Implementation

### Basic Interface Declaration

```go
// Single method interface
type Reader interface {
    Read([]byte) (int, error)
}

// Multiple method interface
type ReadWriter interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
}

// Interface with different method signatures
type Shape interface {
    Area() float64
    Perimeter() float64
    String() string
}
```

### Implementing Interfaces

```go
// Rectangle type implementing Shape interface
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
    return 2 * (r.Width + r.Height)
}

func (r Rectangle) String() string {
    return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

// Circle type implementing Shape interface
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
    return 2 * math.Pi * c.Radius
}

func (c Circle) String() string {
    return fmt.Sprintf("Circle(radius: %.2f)", c.Radius)
}

// Function that works with any Shape
func PrintShapeInfo(s Shape) {
    fmt.Printf("Shape: %s\n", s.String())
    fmt.Printf("Area: %.2f\n", s.Area())
    fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}
```

### Interface Satisfaction Rules

```go
// A type satisfies an interface if it implements ALL methods
type Printer interface {
    Print() string
    Format(string) string
}

type Document struct {
    content string
}

func (d Document) Print() string {
    return d.content
}

func (d Document) Format(style string) string {
    return fmt.Sprintf("[%s] %s", style, d.content)
}

// Document now satisfies Printer interface
```

## Empty Interface

The empty interface `interface{}` can hold values of any type since every type implements zero methods.

### Declaration and Usage

```go
// Empty interface
var anything interface{}

// Can hold any value
anything = 42
anything = "hello"
anything = []int{1, 2, 3}
anything = map[string]int{"a": 1}

// Function accepting any type
func PrintAnything(value interface{}) {
    fmt.Printf("Value: %v, Type: %T\n", value, value)
}

// Slice of any type
func ProcessItems(items []interface{}) {
    for i, item := range items {
        fmt.Printf("Item %d: %v (%T)\n", i, item, item)
    }
}
```

### Working with Empty Interface

```go
// Type checking with empty interface
func DescribeValue(v interface{}) {
    switch v := v.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s (length: %d)\n", v, len(v))
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    case []int:
        fmt.Printf("Integer slice: %v\n", v)
    case map[string]int:
        fmt.Printf("String-to-int map: %v\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}
```

## Type Assertions and Type Switches

### Type Assertions

Extract the concrete value from an interface.

```go
// Basic type assertion
var i interface{} = "hello"

// Safe type assertion (with ok check)
if s, ok := i.(string); ok {
    fmt.Printf("String value: %s\n", s)
} else {
    fmt.Println("Not a string")
}

// Unsafe type assertion (panics if wrong type)
s := i.(string) // Only use when you're certain of the type

// Multiple type assertions
func GetStringValue(v interface{}) (string, error) {
    switch val := v.(type) {
    case string:
        return val, nil
    case int:
        return strconv.Itoa(val), nil
    case float64:
        return strconv.FormatFloat(val, 'f', -1, 64), nil
    case bool:
        return strconv.FormatBool(val), nil
    default:
        return "", fmt.Errorf("cannot convert %T to string", v)
    }
}
```

### Type Switches

Pattern matching on interface types.

```go
// Basic type switch
func ProcessValue(v interface{}) {
    switch v := v.(type) {
    case nil:
        fmt.Println("Value is nil")
    case int:
        fmt.Printf("Integer: %d\n", v*2)
    case string:
        fmt.Printf("String: %s (uppercase: %s)\n", v, strings.ToUpper(v))
    case bool:
        fmt.Printf("Boolean: %t (inverted: %t)\n", v, !v)
    case []int:
        sum := 0
        for _, n := range v {
            sum += n
        }
        fmt.Printf("Integer slice sum: %d\n", sum)
    case map[string]interface{}:
        fmt.Printf("Map with %d keys\n", len(v))
        for key, val := range v {
            fmt.Printf("  %s: %v\n", key, val)
        }
    default:
        fmt.Printf("Unhandled type: %T\n", v)
    }
}

// Type switch with interface checking
func AnalyzeInterface(v interface{}) {
    switch val := v.(type) {
    case fmt.Stringer: // Check if implements Stringer interface
        fmt.Printf("Stringer: %s\n", val.String())
    case error: // Check if implements error interface
        fmt.Printf("Error: %s\n", val.Error())
    case io.Reader: // Check if implements Reader interface
        fmt.Println("This value can be read from")
    case io.Writer: // Check if implements Writer interface
        fmt.Println("This value can be written to")
    default:
        fmt.Printf("Value: %v (%T)\n", val, val)
    }
}
```

## Interface Embedding

Compose larger interfaces from smaller ones.

### Basic Embedding

```go
// Basic interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Embedded interfaces
type ReadWriter interface {
    Reader  // Embeds Reader interface
    Writer  // Embeds Writer interface
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// Equivalent to:
type ReadWriteCloserExplicit interface {
    Read([]byte) (int, error)
    Write([]byte) (int, error)
    Close() error
}
```

### Complex Interface Composition

```go
// Database interfaces
type Connector interface {
    Connect() error
    Disconnect() error
}

type Querier interface {
    Query(string) ([]map[string]interface{}, error)
    QueryRow(string) (map[string]interface{}, error)
}

type Executer interface {
    Execute(string) error
    ExecuteMany([]string) error
}

type Transactioner interface {
    BeginTransaction() (Transaction, error)
}

// Composed database interface
type Database interface {
    Connector
    Querier
    Executer
    Transactioner
}

// Implementation
type PostgresDB struct {
    host, user, password, database string
    connected bool
}

func (db *PostgresDB) Connect() error {
    fmt.Printf("Connecting to PostgreSQL: %s@%s/%s\n", db.user, db.host, db.database)
    db.connected = true
    return nil
}

func (db *PostgresDB) Disconnect() error {
    fmt.Println("Disconnecting from PostgreSQL")
    db.connected = false
    return nil
}

func (db *PostgresDB) Query(sql string) ([]map[string]interface{}, error) {
    if !db.connected {
        return nil, fmt.Errorf("database not connected")
    }
    fmt.Printf("Executing query: %s\n", sql)
    return []map[string]interface{}{}, nil
}

func (db *PostgresDB) QueryRow(sql string) (map[string]interface{}, error) {
    if !db.connected {
        return nil, fmt.Errorf("database not connected")
    }
    fmt.Printf("Executing single row query: %s\n", sql)
    return map[string]interface{}{}, nil
}

func (db *PostgresDB) Execute(sql string) error {
    if !db.connected {
        return fmt.Errorf("database not connected")
    }
    fmt.Printf("Executing statement: %s\n", sql)
    return nil
}

func (db *PostgresDB) ExecuteMany(sqls []string) error {
    if !db.connected {
        return fmt.Errorf("database not connected")
    }
    fmt.Printf("Executing %d statements\n", len(sqls))
    return nil
}

func (db *PostgresDB) BeginTransaction() (Transaction, error) {
    if !db.connected {
        return nil, fmt.Errorf("database not connected")
    }
    return &PostgresTransaction{}, nil
}

type Transaction interface {
    Commit() error
    Rollback() error
}

type PostgresTransaction struct{}

func (tx *PostgresTransaction) Commit() error {
    fmt.Println("Transaction committed")
    return nil
}

func (tx *PostgresTransaction) Rollback() error {
    fmt.Println("Transaction rolled back")
    return nil
}
```

## Common Standard Library Interfaces

### fmt Package Interfaces

```go
// Stringer interface for custom string representation
type Stringer interface {
    String() string
}

// Example implementation
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%s (age %d)", p.Name, p.Age)
}

// GoStringer interface for Go syntax representation
type GoStringer interface {
    GoString() string
}

func (p Person) GoString() string {
    return fmt.Sprintf("Person{Name: %q, Age: %d}", p.Name, p.Age)
}

// Using the interfaces
func main() {
    p := Person{Name: "Alice", Age: 30}
    fmt.Println(p)         // Uses String() method
    fmt.Printf("%#v\n", p) // Uses GoString() method if available
}
```

### io Package Interfaces

```go
// Core io interfaces
type Reader interface {
    Read([]byte) (int, error)
}

type Writer interface {
    Write([]byte) (int, error)
}

type Closer interface {
    Close() error
}

// Compound interfaces
type ReadCloser interface {
    Reader
    Closer
}

type WriteCloser interface {
    Writer
    Closer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}

// Example implementation - in-memory buffer
type MemoryBuffer struct {
    data   []byte
    offset int
    closed bool
}

func NewMemoryBuffer(data []byte) *MemoryBuffer {
    return &MemoryBuffer{data: data, offset: 0}
}

func (mb *MemoryBuffer) Read(p []byte) (int, error) {
    if mb.closed {
        return 0, fmt.Errorf("buffer is closed")
    }
    if mb.offset >= len(mb.data) {
        return 0, io.EOF
    }
    
    n := copy(p, mb.data[mb.offset:])
    mb.offset += n
    return n, nil
}

func (mb *MemoryBuffer) Write(p []byte) (int, error) {
    if mb.closed {
        return 0, fmt.Errorf("buffer is closed")
    }
    
    mb.data = append(mb.data, p...)
    return len(p), nil
}

func (mb *MemoryBuffer) Close() error {
    mb.closed = true
    return nil
}

// Function that works with any ReadWriteCloser
func ProcessData(rwc io.ReadWriteCloser) error {
    defer rwc.Close()
    
    // Write some data
    _, err := rwc.Write([]byte("Hello, World!"))
    if err != nil {
        return err
    }
    
    // Read it back
    buffer := make([]byte, 1024)
    n, err := rwc.Read(buffer)
    if err != nil && err != io.EOF {
        return err
    }
    
    fmt.Printf("Read %d bytes: %s\n", n, string(buffer[:n]))
    return nil
}
```

### sort Package Interfaces

```go
// Sort interface for custom sorting
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}

// Example: sorting a slice of people by age
type People []Person

func (p People) Len() int {
    return len(p)
}

func (p People) Less(i, j int) bool {
    return p[i].Age < p[j].Age
}

func (p People) Swap(i, j int) {
    p[i], p[j] = p[j], p[i]
}

// Usage
func SortExample() {
    people := People{
        {Name: "Alice", Age: 30},
        {Name: "Bob", Age: 25},
        {Name: "Charlie", Age: 35},
    }
    
    sort.Sort(people)
    fmt.Println("Sorted by age:", people)
}
```

### error Interface

```go
// Built-in error interface
type error interface {
    Error() string
}

// Custom error types
type ValidationError struct {
    Field   string
    Message string
}

func (ve ValidationError) Error() string {
    return fmt.Sprintf("validation failed for field '%s': %s", ve.Field, ve.Message)
}

// Error with additional methods
type HTTPError struct {
    StatusCode int
    Message    string
    Cause      error
}

func (he HTTPError) Error() string {
    if he.Cause != nil {
        return fmt.Sprintf("HTTP %d: %s (caused by: %v)", he.StatusCode, he.Message, he.Cause)
    }
    return fmt.Sprintf("HTTP %d: %s", he.StatusCode, he.Message)
}

func (he HTTPError) Unwrap() error {
    return he.Cause
}

// Usage
func ValidateUser(name string, age int) error {
    if name == "" {
        return ValidationError{Field: "name", Message: "cannot be empty"}
    }
    if age < 0 {
        return ValidationError{Field: "age", Message: "cannot be negative"}
    }
    return nil
}
```

## Interface Best Practices

### 1. Keep Interfaces Small

```go
// Good: Small, focused interface
type Logger interface {
    Log(message string)
}

// Good: Single responsibility
type ConfigReader interface {
    ReadConfig() (map[string]string, error)
}

// Avoid: Large interfaces with many methods
type BadService interface {
    Connect() error
    Disconnect() error
    Read() ([]byte, error)
    Write([]byte) error
    Process() error
    Validate() error
    Transform() error
    Save() error
    // ... many more methods
}
```

### 2. Accept Interfaces, Return Concrete Types

```go
// Good: Accept interface parameter
func ProcessFile(r io.Reader) ([]byte, error) {
    return io.ReadAll(r)
}

// Good: Return concrete type
func NewFileReader(filename string) (*os.File, error) {
    return os.Open(filename)
}

// Avoid: Returning interface when not necessary
func BadNewReader() io.Reader {
    return strings.NewReader("data") // Limits flexibility
}
```

### 3. Define Interfaces at the Point of Use

```go
// Package A defines a service
package service

type EmailService struct{}

func (e EmailService) SendEmail(to, subject, body string) error {
    // Implementation
    return nil
}

func (e EmailService) SendBulkEmail(recipients []string, subject, body string) error {
    // Implementation
    return nil
}

// Package B defines interface for what it needs
package notification

// Define interface where you use it
type EmailSender interface {
    SendEmail(to, subject, body string) error
}

type NotificationManager struct {
    emailSender EmailSender
}

func NewNotificationManager(sender EmailSender) *NotificationManager {
    return &NotificationManager{emailSender: sender}
}

func (nm *NotificationManager) SendWelcomeEmail(userEmail, userName string) error {
    subject := "Welcome!"
    body := fmt.Sprintf("Welcome %s!", userName)
    return nm.emailSender.SendEmail(userEmail, subject, body)
}
```

### 4. Use Interface Segregation

```go
// Good: Segregated interfaces
type Reader interface {
    Read() ([]byte, error)
}

type Writer interface {
    Write([]byte) error
}

type Validator interface {
    Validate() error
}

// Compose when needed
type ReadWriteValidator interface {
    Reader
    Writer
    Validator
}

// Client uses only what it needs
func ReadData(r Reader) ([]byte, error) {
    return r.Read()
}

func WriteData(w Writer, data []byte) error {
    return w.Write(data)
}
```

## Advanced Interface Patterns

### 1. Interface Assertions for Optional Methods

```go
type BasicProcessor interface {
    Process(data []byte) ([]byte, error)
}

type OptionalValidator interface {
    Validate(data []byte) error
}

type OptionalInitializer interface {
    Initialize() error
}

func ProcessWithOptions(processor BasicProcessor, data []byte) ([]byte, error) {
    // Check for optional initialization
    if initializer, ok := processor.(OptionalInitializer); ok {
        if err := initializer.Initialize(); err != nil {
            return nil, fmt.Errorf("initialization failed: %w", err)
        }
    }
    
    // Check for optional validation
    if validator, ok := processor.(OptionalValidator); ok {
        if err := validator.Validate(data); err != nil {
            return nil, fmt.Errorf("validation failed: %w", err)
        }
    }
    
    // Process the data
    return processor.Process(data)
}

// Implementation with optional methods
type AdvancedProcessor struct {
    initialized bool
}

func (ap *AdvancedProcessor) Process(data []byte) ([]byte, error) {
    if !ap.initialized {
        return nil, fmt.Errorf("processor not initialized")
    }
    return append([]byte("processed: "), data...), nil
}

func (ap *AdvancedProcessor) Validate(data []byte) error {
    if len(data) == 0 {
        return fmt.Errorf("data cannot be empty")
    }
    return nil
}

func (ap *AdvancedProcessor) Initialize() error {
    ap.initialized = true
    return nil
}
```

### 2. Interface Wrapping Pattern

```go
// Base interface
type DataStore interface {
    Save(key string, value []byte) error
    Load(key string) ([]byte, error)
    Delete(key string) error
}

// Wrapper with additional functionality
type LoggingDataStore struct {
    store  DataStore
    logger Logger
}

func NewLoggingDataStore(store DataStore, logger Logger) *LoggingDataStore {
    return &LoggingDataStore{store: store, logger: logger}
}

func (lds *LoggingDataStore) Save(key string, value []byte) error {
    lds.logger.Log(fmt.Sprintf("Saving key: %s, size: %d bytes", key, len(value)))
    err := lds.store.Save(key, value)
    if err != nil {
        lds.logger.Log(fmt.Sprintf("Save failed for key %s: %v", key, err))
    }
    return err
}

func (lds *LoggingDataStore) Load(key string) ([]byte, error) {
    lds.logger.Log(fmt.Sprintf("Loading key: %s", key))
    data, err := lds.store.Load(key)
    if err != nil {
        lds.logger.Log(fmt.Sprintf("Load failed for key %s: %v", key, err))
    } else {
        lds.logger.Log(fmt.Sprintf("Loaded key: %s, size: %d bytes", key, len(data)))
    }
    return data, err
}

func (lds *LoggingDataStore) Delete(key string) error {
    lds.logger.Log(fmt.Sprintf("Deleting key: %s", key))
    err := lds.store.Delete(key)
    if err != nil {
        lds.logger.Log(fmt.Sprintf("Delete failed for key %s: %v", key, err))
    }
    return err
}

// Chain multiple wrappers
type CachingDataStore struct {
    store DataStore
    cache map[string][]byte
}

func NewCachingDataStore(store DataStore) *CachingDataStore {
    return &CachingDataStore{
        store: store,
        cache: make(map[string][]byte),
    }
}

func (cds *CachingDataStore) Save(key string, value []byte) error {
    err := cds.store.Save(key, value)
    if err == nil {
        cds.cache[key] = value
    }
    return err
}

func (cds *CachingDataStore) Load(key string) ([]byte, error) {
    // Check cache first
    if data, found := cds.cache[key]; found {
        return data, nil
    }
    
    // Load from store and cache result
    data, err := cds.store.Load(key)
    if err == nil {
        cds.cache[key] = data
    }
    return data, err
}

func (cds *CachingDataStore) Delete(key string) error {
    delete(cds.cache, key)
    return cds.store.Delete(key)
}

// Usage: chain wrappers
func CreateEnhancedDataStore(base DataStore, logger Logger) DataStore {
    cached := NewCachingDataStore(base)
    logged := NewLoggingDataStore(cached, logger)
    return logged
}
```

### 3. Strategy Pattern with Interfaces

```go
// Strategy interface
type SortStrategy interface {
    Sort(data []int)
}

// Concrete strategies
type BubbleSort struct{}

func (bs BubbleSort) Sort(data []int) {
    n := len(data)
    for i := 0; i < n-1; i++ {
        for j := 0; j < n-i-1; j++ {
            if data[j] > data[j+1] {
                data[j], data[j+1] = data[j+1], data[j]
            }
        }
    }
}

type QuickSort struct{}

func (qs QuickSort) Sort(data []int) {
    if len(data) < 2 {
        return
    }
    qs.quickSort(data, 0, len(data)-1)
}

func (qs QuickSort) quickSort(data []int, low, high int) {
    if low < high {
        pi := qs.partition(data, low, high)
        qs.quickSort(data, low, pi-1)
        qs.quickSort(data, pi+1, high)
    }
}

func (qs QuickSort) partition(data []int, low, high int) int {
    pivot := data[high]
    i := low - 1
    
    for j := low; j < high; j++ {
        if data[j] < pivot {
            i++
            data[i], data[j] = data[j], data[i]
        }
    }
    data[i+1], data[high] = data[high], data[i+1]
    return i + 1
}

// Context that uses strategy
type Sorter struct {
    strategy SortStrategy
}

func NewSorter(strategy SortStrategy) *Sorter {
    return &Sorter{strategy: strategy}
}

func (s *Sorter) SetStrategy(strategy SortStrategy) {
    s.strategy = strategy
}

func (s *Sorter) Sort(data []int) {
    s.strategy.Sort(data)
}

// Usage
func DemonstrateStrategyPattern() {
    data1 := []int{64, 34, 25, 12, 22, 11, 90}
    data2 := []int{64, 34, 25, 12, 22, 11, 90}
    
    // Use bubble sort
    sorter := NewSorter(BubbleSort{})
    sorter.Sort(data1)
    fmt.Println("Bubble sorted:", data1)
    
    // Switch to quick sort
    sorter.SetStrategy(QuickSort{})
    sorter.Sort(data2)
    fmt.Println("Quick sorted:", data2)
}
```

## Interface Examples and Use Cases

### 1. Plugin System

```go
// Plugin interface
type Plugin interface {
    Name() string
    Version() string
    Execute(args map[string]interface{}) (interface{}, error)
}

// Plugin manager
type PluginManager struct {
    plugins map[string]Plugin
}

func NewPluginManager() *PluginManager {
    return &PluginManager{
        plugins: make(map[string]Plugin),
    }
}

func (pm *PluginManager) Register(plugin Plugin) {
    pm.plugins[plugin.Name()] = plugin
}

func (pm *PluginManager) Execute(pluginName string, args map[string]interface{}) (interface{}, error) {
    plugin, exists := pm.plugins[pluginName]
    if !exists {
        return nil, fmt.Errorf("plugin %s not found", pluginName)
    }
    return plugin.Execute(args)
}

func (pm *PluginManager) ListPlugins() []string {
    var names []string
    for name := range pm.plugins {
        names = append(names, name)
    }
    return names
}

// Example plugins
type CalculatorPlugin struct{}

func (cp CalculatorPlugin) Name() string { return "calculator" }
func (cp CalculatorPlugin) Version() string { return "1.0.0" }

func (cp CalculatorPlugin) Execute(args map[string]interface{}) (interface{}, error) {
    operation, ok := args["operation"].(string)
    if !ok {
        return nil, fmt.Errorf("operation not specified")
    }
    
    a, aOk := args["a"].(float64)
    b, bOk := args["b"].(float64)
    if !aOk || !bOk {
        return nil, fmt.Errorf("invalid operands")
    }
    
    switch operation {
    case "add":
        return a + b, nil
    case "subtract":
        return a - b, nil
    case "multiply":
        return a * b, nil
    case "divide":
        if b == 0 {
            return nil, fmt.Errorf("division by zero")
        }
        return a / b, nil
    default:
        return nil, fmt.Errorf("unknown operation: %s", operation)
    }
}

type TextProcessorPlugin struct{}

func (tp TextProcessorPlugin) Name() string { return "text-processor" }
func (tp TextProcessorPlugin) Version() string { return "1.0.0" }

func (tp TextProcessorPlugin) Execute(args map[string]interface{}) (interface{}, error) {
    text, ok := args["text"].(string)
    if !ok {
        return nil, fmt.Errorf("text not specified")
    }
    
    operation, ok := args["operation"].(string)
    if !ok {
        operation = "uppercase"
    }
    
    switch operation {
    case "uppercase":
        return strings.ToUpper(text), nil
    case "lowercase":
        return strings.ToLower(text), nil
    case "reverse":
        runes := []rune(text)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
            runes[i], runes[j] = runes[j], runes[i]
        }
        return string(runes), nil
    case "wordcount":
        return len(strings.Fields(text)), nil
    default:
        return nil, fmt.Errorf("unknown operation: %s", operation)
    }
}
```

### 2. Notification System

```go
// Notification interface
type Notifier interface {
    Send(message Message) error
    SupportsChannel(channel string) bool
}

// Message structure
type Message struct {
    Title     string
    Body      string
    Priority  string
    Metadata  map[string]string
    Recipients []string
}

// Different notification implementations
type EmailNotifier struct {
    smtpHost string
    port     int
}

func (en EmailNotifier) Send(message Message) error {
    fmt.Printf("Sending email: %s to %v\n", message.Title, message.Recipients)
    return nil
}

func (en EmailNotifier) SupportsChannel(channel string) bool {
    return channel == "email"
}

type SMSNotifier struct {
    apiKey string
}

func (sms SMSNotifier) Send(message Message) error {
    fmt.Printf("Sending SMS: %s to %v\n", message.Body, message.Recipients)
    return nil
}

func (sms SMSNotifier) SupportsChannel(channel string) bool {
    return channel == "sms"
}

type SlackNotifier struct {
    webhookURL string
}

func (slack SlackNotifier) Send(message Message) error {
    fmt.Printf("Sending Slack message: %s\n", message.Title)
    return nil
}

func (slack SlackNotifier) SupportsChannel(channel string) bool {
    return channel == "slack"
}

// Notification service
type NotificationService struct {
    notifiers []Notifier
}

func NewNotificationService() *NotificationService {
    return &NotificationService{
        notifiers: make([]Notifier, 0),
    }
}

func (ns *NotificationService) AddNotifier(notifier Notifier) {
    ns.notifiers = append(ns.notifiers, notifier)
}

func (ns *NotificationService) SendToChannel(channel string, message Message) error {
    for _, notifier := range ns.notifiers {
        if notifier.SupportsChannel(channel) {
            return notifier.Send(message)
        }
    }
    return fmt.Errorf("no notifier found for channel: %s", channel)
}

func (ns *NotificationService) Broadcast(message Message, channels []string) []error {
    var errors []error
    for _, channel := range channels {
        if err := ns.SendToChannel(channel, message); err != nil {
            errors = append(errors, err)
        }
    }
    return errors
}

// Usage
func DemonstrateNotificationSystem() {
    service := NewNotificationService()
    service.AddNotifier(EmailNotifier{smtpHost: "smtp.gmail.com", port: 587})
    service.AddNotifier(SMSNotifier{apiKey: "twilio-key"})
    service.AddNotifier(SlackNotifier{webhookURL: "https://hooks.slack.com/..."})
    
    message := Message{
        Title:      "System Alert",
        Body:       "Database connection restored",
        Priority:   "high",
        Recipients: []string{"admin@company.com", "+1234567890", "#general"},
    }
    
    errors := service.Broadcast(message, []string{"email", "sms", "slack"})
    if len(errors) > 0 {
        fmt.Printf("Broadcast completed with %d errors\n", len(errors))
    }
}
```

### 3. Middleware Chain Pattern

```go
// HTTP middleware using interfaces
type Handler interface {
    Handle(ctx Context) error
}

type Middleware interface {
    Process(ctx Context, next Handler) error
}

type Context interface {
    GetRequestID() string
    GetUserID() string
    SetValue(key string, value interface{})
    GetValue(key string) interface{}
    Write(data []byte) error
}

// Simple context implementation
type SimpleContext struct {
    requestID string
    userID    string
    values    map[string]interface{}
    response  []byte
}

func NewSimpleContext(requestID, userID string) *SimpleContext {
    return &SimpleContext{
        requestID: requestID,
        userID:    userID,
        values:    make(map[string]interface{}),
    }
}

func (sc *SimpleContext) GetRequestID() string { return sc.requestID }
func (sc *SimpleContext) GetUserID() string    { return sc.userID }

func (sc *SimpleContext) SetValue(key string, value interface{}) {
    sc.values[key] = value
}

func (sc *SimpleContext) GetValue(key string) interface{} {
    return sc.values[key]
}

func (sc *SimpleContext) Write(data []byte) error {
    sc.response = append(sc.response, data...)
    return nil
}

// Middleware implementations
type LoggingMiddleware struct {
    logger Logger
}

func (lm LoggingMiddleware) Process(ctx Context, next Handler) error {
    lm.logger.Log(fmt.Sprintf("Request started: %s", ctx.GetRequestID()))
    start := time.Now()
    
    err := next.Handle(ctx)
    
    duration := time.Since(start)
    lm.logger.Log(fmt.Sprintf("Request completed: %s (took %v)", ctx.GetRequestID(), duration))
    
    return err
}

type AuthMiddleware struct {
    requiredRole string
}

func (am AuthMiddleware) Process(ctx Context, next Handler) error {
    userID := ctx.GetUserID()
    if userID == "" {
        return fmt.Errorf("authentication required")
    }
    
    // Simulate role checking
    userRole := ctx.GetValue("role")
    if userRole != am.requiredRole && am.requiredRole != "" {
        return fmt.Errorf("insufficient permissions")
    }
    
    return next.Handle(ctx)
}

type RateLimitingMiddleware struct {
    requests map[string]int
    limit    int
}

func NewRateLimitingMiddleware(limit int) *RateLimitingMiddleware {
    return &RateLimitingMiddleware{
        requests: make(map[string]int),
        limit:    limit,
    }
}

func (rlm *RateLimitingMiddleware) Process(ctx Context, next Handler) error {
    userID := ctx.GetUserID()
    if userID == "" {
        return fmt.Errorf("user identification required for rate limiting")
    }
    
    rlm.requests[userID]++
    if rlm.requests[userID] > rlm.limit {
        return fmt.Errorf("rate limit exceeded for user %s", userID)
    }
    
    return next.Handle(ctx)
}

// Handler implementations
type UserHandler struct{}

func (uh UserHandler) Handle(ctx Context) error {
    userID := ctx.GetUserID()
    response := fmt.Sprintf("User profile for: %s", userID)
    return ctx.Write([]byte(response))
}

type AdminHandler struct{}

func (ah AdminHandler) Handle(ctx Context) error {
    response := "Admin dashboard data"
    return ctx.Write([]byte(response))
}

// Middleware chain
type HandlerChain struct {
    middlewares []Middleware
    handler     Handler
}

func NewHandlerChain(handler Handler) *HandlerChain {
    return &HandlerChain{
        middlewares: make([]Middleware, 0),
        handler:     handler,
    }
}

func (hc *HandlerChain) Use(middleware Middleware) *HandlerChain {
    hc.middlewares = append(hc.middlewares, middleware)
    return hc
}

func (hc *HandlerChain) Handle(ctx Context) error {
    return hc.processMiddleware(ctx, 0)
}

func (hc *HandlerChain) processMiddleware(ctx Context, index int) error {
    if index >= len(hc.middlewares) {
        return hc.handler.Handle(ctx)
    }
    
    middleware := hc.middlewares[index]
    next := HandlerFunc(func(ctx Context) error {
        return hc.processMiddleware(ctx, index+1)
    })
    
    return middleware.Process(ctx, next)
}

// Handler function adapter
type HandlerFunc func(ctx Context) error

func (hf HandlerFunc) Handle(ctx Context) error {
    return hf(ctx)
}

// Usage
func DemonstrateMiddlewareChain() {
    logger := &SimpleLogger{}
    
    // User endpoint with logging and rate limiting
    userChain := NewHandlerChain(UserHandler{}).
        Use(LoggingMiddleware{logger: logger}).
        Use(NewRateLimitingMiddleware(10))
    
    // Admin endpoint with logging, auth, and rate limiting
    adminChain := NewHandlerChain(AdminHandler{}).
        Use(LoggingMiddleware{logger: logger}).
        Use(AuthMiddleware{requiredRole: "admin"}).
        Use(NewRateLimitingMiddleware(5))
    
    // Simulate requests
    userCtx := NewSimpleContext("req-1", "user123")
    userCtx.SetValue("role", "user")
    
    adminCtx := NewSimpleContext("req-2", "admin456")
    adminCtx.SetValue("role", "admin")
    
    fmt.Println("Processing user request...")
    if err := userChain.Handle(userCtx); err != nil {
        fmt.Printf("User request failed: %v\n", err)
    }
    
    fmt.Println("Processing admin request...")
    if err := adminChain.Handle(adminCtx); err != nil {
        fmt.Printf("Admin request failed: %v\n", err)
    }
}

// Simple logger implementation
type SimpleLogger struct{}

func (sl *SimpleLogger) Log(message string) {
    fmt.Printf("[LOG] %s: %s\n", time.Now().Format(time.RFC3339), message)
}
```

## Summary

Go interfaces are a powerful tool for creating flexible, testable, and maintainable code. Key takeaways:

### Core Principles
- **Implicit implementation** - types automatically satisfy interfaces
- **Small interfaces** - prefer single-method interfaces when possible
- **Accept interfaces, return concrete types** - for maximum flexibility
- **Define interfaces where they're used** - not where they're implemented

### Common Patterns
- **Composition** - build complex interfaces from simple ones
- **Wrapping** - add behavior without changing existing code
- **Strategy** - swap implementations at runtime
- **Plugin systems** - load and use external components

### Best Practices
- Keep interfaces focused and cohesive
- Use type assertions carefully with proper error checking
- Leverage standard library interfaces when possible
- Design for testability and mockability

Interfaces are the key to writing idiomatic Go code that follows the principle of "Don't Repeat Yourself" while maintaining loose coupling between components. They enable polymorphism without inheritance and make Go code both powerful and elegant.

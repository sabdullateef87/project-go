# Text Templates in Go

The `text/template` package provides data-driven templates for generating text output. It's ideal for code generation, configuration files, and formatted text output.

## Table of Contents

- [Overview](#overview)
- [Basic Template Syntax](#basic-template-syntax)
- [Actions](#actions)
- [Variables](#variables)
- [Pipelines](#pipelines)
- [Functions](#functions)
- [Control Structures](#control-structures)
- [Template Composition](#template-composition)
- [Custom Functions](#custom-functions)
- [Error Handling](#error-handling)
- [Common Patterns](#common-patterns)
- [Performance](#performance)
- [Best Practices](#best-practices)

## Overview

The `text/template` package:

- **Data-driven** - Templates are populated with data
- **Logic in templates** - Conditionals, loops, variables
- **Composable** - Templates can include other templates
- **Extensible** - Custom functions can be added
- **Type-safe** - Compile-time template parsing

```go
import "text/template"
```

## Basic Template Syntax

### Simple Template

```go
package main

import (
    "os"
    "text/template"
)

func main() {
    tmpl := template.Must(template.New("test").Parse("Hello, {{.}}!"))
    tmpl.Execute(os.Stdout, "World")
    // Output: Hello, World!
}
```

### Template with Struct

```go
type Person struct {
    Name string
    Age  int
}

func main() {
    tmpl := template.Must(template.New("person").Parse(`
Name: {{.Name}}
Age: {{.Age}}
`))

    person := Person{Name: "Alice", Age: 30}
    tmpl.Execute(os.Stdout, person)
    // Output:
    // Name: Alice
    // Age: 30
}
```

### Creating Templates

```go
// Method 1: New + Parse
tmpl, err := template.New("name").Parse("template string")

// Method 2: Must (panics on error, useful for static templates)
tmpl := template.Must(template.New("name").Parse("template string"))

// Method 3: ParseFiles
tmpl, err := template.ParseFiles("template.txt")

// Method 4: ParseGlob
tmpl, err := template.ParseGlob("templates/*.txt")
```

## Actions

Actions are delimited by `{{` and `}}`.

### Basic Actions

```go
// Comment (ignored in output)
{{/* This is a comment */}}

// Output value
{{.}}

// Output field
{{.FieldName}}

// Output method result
{{.MethodName}}

// Pipeline
{{.Field | printf "%s"}}
```

### Whitespace Control

```go
// Remove whitespace before action
{{- .Value}}

// Remove whitespace after action
{{.Value -}}

// Remove both
{{- .Value -}}

// Example
tmpl := `
    {{- "Hello" -}}
    {{- " World" -}}
`
// Output: "HelloWorld" (no spaces or newlines)
```

## Variables

Define and use variables in templates.

### Basic Variables

```go
tmpl := `
{{$name := .Name}}
{{$age := .Age}}
Hello, {{$name}}! You are {{$age}} years old.
`

person := Person{Name: "Alice", Age: 30}
template.Must(template.New("vars").Parse(tmpl)).Execute(os.Stdout, person)
// Output: Hello, Alice! You are 30 years old.
```

### Variables in Range

```go
tmpl := `
{{range $index, $value := .Items}}
{{$index}}: {{$value}}
{{end}}
`

data := struct {
    Items []string
}{
    Items: []string{"apple", "banana", "cherry"},
}

template.Must(template.New("range").Parse(tmpl)).Execute(os.Stdout, data)
// Output:
// 0: apple
// 1: banana
// 2: cherry
```

### Reassignment

```go
tmpl := `
{{$count := 0}}
{{range .Items}}
{{$count = add $count 1}}
{{end}}
Total: {{$count}}
`
```

## Pipelines

Chain operations from left to right.

### Basic Pipeline

```go
// Value | function
{{.Name | printf "Hello, %s!"}}

// Multiple functions
{{.Name | printf "Name: %s" | println}}

// With arguments
{{.Price | printf "%.2f" | println}}
```

### Built-in Functions in Pipelines

```go
// print, printf, println
{{.Value | print}}
{{.Value | printf "Value: %d"}}
{{.Value | println}}

// and
{{and .Condition1 .Condition2}}

// or
{{or .Condition1 .Condition2}}

// not
{{not .Condition}}

// eq, ne, lt, le, gt, ge
{{if eq .Status "active"}}Active{{end}}
{{if ne .Count 0}}Count: {{.Count}}{{end}}
{{if lt .Age 18}}Minor{{else}}Adult{{end}}
```

## Functions

### Built-in Functions

```go
// Comparison
{{eq .X .Y}}    // X == Y
{{ne .X .Y}}    // X != Y
{{lt .X .Y}}    // X < Y
{{le .X .Y}}    // X <= Y
{{gt .X .Y}}    // X > Y
{{ge .X .Y}}    // X >= Y

// Logic
{{and .X .Y}}   // X && Y
{{or .X .Y}}    // X || Y
{{not .X}}      // !X

// Formatting
{{print .X}}
{{printf "%s: %d" .Name .Age}}
{{println .X}}

// Length
{{len .Slice}}
{{len .Map}}
{{len .String}}

// Index
{{index .Slice 0}}
{{index .Map "key"}}

// Slice
{{slice .Array 1 3}}  // .Array[1:3]
```

### Template Function Examples

```go
type Data struct {
    Name  string
    Items []string
    Count int
}

tmpl := `
Name: {{.Name}}
Items ({{len .Items}}):
{{range .Items}}
  - {{.}}
{{end}}
{{if gt .Count 0}}
Count is positive: {{.Count}}
{{end}}
`

data := Data{
    Name:  "Shopping List",
    Items: []string{"apple", "banana", "cherry"},
    Count: 5,
}
```

## Control Structures

### If-Else

```go
// Simple if
{{if .Condition}}
    Content when true
{{end}}

// If-else
{{if .Condition}}
    True content
{{else}}
    False content
{{end}}

// If-else if-else
{{if eq .Status "active"}}
    Active
{{else if eq .Status "pending"}}
    Pending
{{else}}
    Other
{{end}}

// With pipeline
{{if .Items}}
    Items exist
{{else}}
    No items
{{end}}
```

### Range

```go
// Range over slice
{{range .Items}}
    {{.}}
{{end}}

// Range with index and value
{{range $index, $value := .Items}}
    {{$index}}: {{$value}}
{{end}}

// Range over map
{{range $key, $value := .Map}}
    {{$key}} = {{$value}}
{{end}}

// Range with else (empty collection)
{{range .Items}}
    Item: {{.}}
{{else}}
    No items
{{end}}
```

### With

Change the context (dot).

```go
// With sets dot to pipeline value
{{with .Person}}
    Name: {{.Name}}
    Age: {{.Age}}
{{end}}

// With-else
{{with .OptionalValue}}
    Value exists: {{.}}
{{else}}
    No value
{{end}}

// Nested with
{{with .User}}
    User: {{.Name}}
    {{with .Address}}
        City: {{.City}}
    {{end}}
{{end}}
```

### Break and Continue

Not directly supported, but can use conditions.

```go
// Workaround for break
{{range .Items}}
    {{if condition}}
        {{/* continue logic */}}
    {{else}}
        {{/* break logic */}}
        {{break}}  // Not valid, use external logic
    {{end}}
{{end}}
```

## Template Composition

### Define and Template

```go
// Define named template
{{define "header"}}
    <header>{{.Title}}</header>
{{end}}

// Use template
{{template "header" .}}

// With data
{{template "header" .HeaderData}}
```

### Example: Layout with Partials

```go
const layout = `
{{define "layout"}}
<!DOCTYPE html>
<html>
<head>
    {{template "head" .}}
</head>
<body>
    {{template "body" .}}
</body>
</html>
{{end}}

{{define "head"}}
    <title>{{.Title}}</title>
{{end}}
`

const content = `
{{define "body"}}
    <h1>{{.Heading}}</h1>
    <p>{{.Content}}</p>
{{end}}
`

// Parse all templates
tmpl := template.Must(template.New("layout").Parse(layout))
tmpl = template.Must(tmpl.Parse(content))

// Execute the layout
data := struct {
    Title   string
    Heading string
    Content string
}{
    Title:   "My Page",
    Heading: "Welcome",
    Content: "Hello, World!",
}

tmpl.ExecuteTemplate(os.Stdout, "layout", data)
```

### Nested Templates

```go
// Base template
const base = `
{{define "base"}}
Header
{{template "content" .}}
Footer
{{end}}
`

// Content templates
const page1 = `
{{define "content"}}
This is page 1: {{.Name}}
{{end}}
`

const page2 = `
{{define "content"}}
This is page 2: {{.Name}}
{{end}}
`

// Parse all
tmpl := template.New("templates")
template.Must(tmpl.Parse(base))
template.Must(tmpl.Parse(page1))

// Execute
tmpl.ExecuteTemplate(os.Stdout, "base", map[string]string{"Name": "Alice"})
```

### Block

Define a template with default content that can be overridden.

```go
const base = `
{{define "base"}}
<html>
<head>
    {{block "title" .}}Default Title{{end}}
</head>
<body>
    {{block "content" .}}Default Content{{end}}
</body>
</html>
{{end}}
`

const page = `
{{define "title"}}Custom Title{{end}}
{{define "content"}}Custom Content{{end}}
`

// Parse both
tmpl := template.Must(template.New("base").Parse(base))
tmpl = template.Must(tmpl.Parse(page))

tmpl.ExecuteTemplate(os.Stdout, "base", nil)
```

## Custom Functions

Add custom functions to templates.

### Basic Custom Function

```go
func upper(s string) string {
    return strings.ToUpper(s)
}

funcMap := template.FuncMap{
    "upper": upper,
}

tmpl := template.Must(template.New("test").Funcs(funcMap).Parse(`
{{.Name | upper}}
`))

tmpl.Execute(os.Stdout, map[string]string{"Name": "alice"})
// Output: ALICE
```

### Multiple Functions

```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
    "lower": strings.ToLower,
    "title": strings.Title,
    "add": func(a, b int) int {
        return a + b
    },
    "multiply": func(a, b int) int {
        return a * b
    },
    "formatDate": func(t time.Time) string {
        return t.Format("2006-01-02")
    },
}

tmpl := template.New("test").Funcs(funcMap)
```

### Function with Multiple Return Values

```go
// Functions can return (value, error)
funcMap := template.FuncMap{
    "divide": func(a, b float64) (float64, error) {
        if b == 0 {
            return 0, errors.New("division by zero")
        }
        return a / b, nil
    },
}

tmpl := `Result: {{divide .A .B}}`
```

### Variadic Functions

```go
funcMap := template.FuncMap{
    "sum": func(nums ...int) int {
        total := 0
        for _, n := range nums {
            total += n
        }
        return total
    },
    "join": func(sep string, strs ...string) string {
        return strings.Join(strs, sep)
    },
}

tmpl := `
Total: {{sum 1 2 3 4 5}}
Joined: {{join ", " "a" "b" "c"}}
`
```

### Useful Custom Functions

```go
funcMap := template.FuncMap{
    // String operations
    "contains": strings.Contains,
    "hasPrefix": strings.HasPrefix,
    "hasSuffix": strings.HasSuffix,
    "replace": strings.ReplaceAll,
    "split": strings.Split,
    "trim": strings.TrimSpace,

    // Math operations
    "add": func(a, b int) int { return a + b },
    "sub": func(a, b int) int { return a - b },
    "mul": func(a, b int) int { return a * b },
    "div": func(a, b int) int { return a / b },

    // Formatting
    "formatInt": func(i int) string {
        return fmt.Sprintf("%d", i)
    },
    "formatFloat": func(f float64, precision int) string {
        return fmt.Sprintf("%.*f", precision, f)
    },

    // Date/Time
    "now": time.Now,
    "formatTime": func(t time.Time, layout string) string {
        return t.Format(layout)
    },

    // Collections
    "first": func(items []interface{}) interface{} {
        if len(items) > 0 {
            return items[0]
        }
        return nil
    },
    "last": func(items []interface{}) interface{} {
        if len(items) > 0 {
            return items[len(items)-1]
        }
        return nil
    },

    // Logic
    "default": func(defaultVal, val interface{}) interface{} {
        if val == nil || val == "" {
            return defaultVal
        }
        return val
    },
}
```

## Error Handling

### Template Parsing Errors

```go
tmpl, err := template.New("test").Parse("{{.InvalidSyntax")
if err != nil {
    log.Fatal("Parse error:", err)
}
```

### Execution Errors

```go
type Data struct {
    Name string
}

tmpl := template.Must(template.New("test").Parse("{{.NonExistentField}}"))

err := tmpl.Execute(os.Stdout, Data{Name: "Alice"})
if err != nil {
    log.Fatal("Execution error:", err)
}
```

### Custom Error Handling

```go
var buf bytes.Buffer
err := tmpl.Execute(&buf, data)
if err != nil {
    // Handle error
    log.Printf("Template execution failed: %v", err)
    // Return default content or error page
    return
}

// Use the buffer
output := buf.String()
```

### Option: MissingKey

Control behavior when accessing missing keys.

```go
// Default: "missingkey=default" - print "<no value>"
// Options: "missingkey=invalid" - return error
//          "missingkey=zero" - return zero value

tmpl := template.New("test").Option("missingkey=error")
tmpl.Parse("{{.MissingField}}")

err := tmpl.Execute(os.Stdout, struct{}{})
// err will be non-nil
```

## Common Patterns

### Configuration File Generation

```go
type Config struct {
    AppName     string
    Port        int
    Database    string
    LogLevel    string
    EnableDebug bool
}

const configTemplate = `
# {{.AppName}} Configuration

app_name = "{{.AppName}}"
port = {{.Port}}
database = "{{.Database}}"
log_level = "{{.LogLevel}}"
enable_debug = {{.EnableDebug}}
`

func generateConfig() {
    config := Config{
        AppName:     "MyApp",
        Port:        8080,
        Database:    "postgres://localhost/mydb",
        LogLevel:    "info",
        EnableDebug: false,
    }

    tmpl := template.Must(template.New("config").Parse(configTemplate))

    file, _ := os.Create("config.toml")
    defer file.Close()

    tmpl.Execute(file, config)
}
```

### Code Generation

```go
type Model struct {
    Name   string
    Fields []Field
}

type Field struct {
    Name string
    Type string
}

const modelTemplate = `
type {{.Name}} struct {
{{range .Fields}}
    {{.Name}} {{.Type}}
{{end}}
}

func New{{.Name}}({{range $i, $f := .Fields}}{{if $i}}, {{end}}{{$f.Name}} {{$f.Type}}{{end}}) *{{.Name}} {
    return &{{.Name}}{
{{range .Fields}}
        {{.Name}}: {{.Name}},
{{end}}
    }
}
`

func generateModel() {
    model := Model{
        Name: "User",
        Fields: []Field{
            {Name: "ID", Type: "int"},
            {Name: "Name", Type: "string"},
            {Name: "Email", Type: "string"},
        },
    }

    tmpl := template.Must(template.New("model").Parse(modelTemplate))
    tmpl.Execute(os.Stdout, model)
}
```

### Email Templates

```go
type EmailData struct {
    RecipientName string
    Subject       string
    Message       string
    ActionURL     string
    ActionText    string
}

const emailTemplate = `
Hello {{.RecipientName}},

{{.Message}}

{{if .ActionURL}}
Click here to continue: {{.ActionURL}}
{{.ActionText}}
{{end}}

Best regards,
The Team
`

func sendEmail(to string, data EmailData) {
    tmpl := template.Must(template.New("email").Parse(emailTemplate))

    var body bytes.Buffer
    tmpl.Execute(&body, data)

    // Send email with body.String()
}
```

### Report Generation

```go
type Report struct {
    Title       string
    GeneratedAt time.Time
    Data        []ReportItem
    TotalItems  int
    Summary     string
}

type ReportItem struct {
    Name  string
    Value float64
}

const reportTemplate = `
{{.Title}}
Generated: {{.GeneratedAt.Format "2006-01-02 15:04:05"}}
========================================

Items ({{.TotalItems}}):
{{range .Data}}
{{.Name}}: {{printf "%.2f" .Value}}
{{end}}

Summary:
{{.Summary}}
`

funcMap := template.FuncMap{
    "add": func(a, b int) int { return a + b },
}

func generateReport(report Report) string {
    tmpl := template.Must(
        template.New("report").
            Funcs(funcMap).
            Parse(reportTemplate),
    )

    var buf bytes.Buffer
    tmpl.Execute(&buf, report)
    return buf.String()
}
```

### SQL Query Generation

```go
type Query struct {
    Table   string
    Columns []string
    Where   []Condition
    OrderBy string
    Limit   int
}

type Condition struct {
    Column   string
    Operator string
    Value    interface{}
}

const queryTemplate = `
SELECT {{range $i, $col := .Columns}}{{if $i}}, {{end}}{{$col}}{{end}}
FROM {{.Table}}
{{if .Where}}
WHERE {{range $i, $cond := .Where}}{{if $i}} AND {{end}}{{$cond.Column}} {{$cond.Operator}} {{$cond.Value}}{{end}}
{{end}}
{{if .OrderBy}}
ORDER BY {{.OrderBy}}
{{end}}
{{if .Limit}}
LIMIT {{.Limit}}
{{end}}
`
```

## Performance

### Parse Once, Execute Many

```go
// Bad: Parse template repeatedly
for _, user := range users {
    tmpl := template.Must(template.New("user").Parse(userTemplate))
    tmpl.Execute(writer, user)
}

// Good: Parse once, execute many times
tmpl := template.Must(template.New("user").Parse(userTemplate))
for _, user := range users {
    tmpl.Execute(writer, user)
}
```

### Buffer Output

```go
// Good: Write to buffer first
var buf bytes.Buffer
tmpl.Execute(&buf, data)

// Then write to final destination
writer.Write(buf.Bytes())

// Or use sync.Pool for buffers
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func executeTemplate(tmpl *template.Template, data interface{}) []byte {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)

    tmpl.Execute(buf, data)

    result := make([]byte, buf.Len())
    copy(result, buf.Bytes())
    return result
}
```

### Cache Parsed Templates

```go
var templateCache = make(map[string]*template.Template)
var cacheMutex sync.RWMutex

func getTemplate(name, content string) (*template.Template, error) {
    cacheMutex.RLock()
    tmpl, ok := templateCache[name]
    cacheMutex.RUnlock()

    if ok {
        return tmpl, nil
    }

    cacheMutex.Lock()
    defer cacheMutex.Unlock()

    // Double-check after acquiring write lock
    if tmpl, ok := templateCache[name]; ok {
        return tmpl, nil
    }

    tmpl, err := template.New(name).Parse(content)
    if err != nil {
        return nil, err
    }

    templateCache[name] = tmpl
    return tmpl, nil
}
```

## Best Practices

### 1. Use Must for Static Templates

```go
// Good: Panic at startup if template is invalid
var tmpl = template.Must(template.New("name").Parse(templateString))

// For dynamic templates, handle error properly
tmpl, err := template.New("name").Parse(dynamicString)
if err != nil {
    return fmt.Errorf("template parse error: %w", err)
}
```

### 2. Validate Data Before Execution

```go
func executeTemplate(tmpl *template.Template, data interface{}) error {
    // Validate data
    if err := validateData(data); err != nil {
        return err
    }

    return tmpl.Execute(os.Stdout, data)
}
```

### 3. Use Named Templates

```go
// Good: Named templates are reusable
{{define "header"}}...{{end}}
{{template "header" .}}

// Avoid: Inline everything
```

### 4. Add Custom Functions for Common Operations

```go
funcMap := template.FuncMap{
    "formatDate": formatDate,
    "currency":   formatCurrency,
    "truncate":   truncateString,
}

tmpl := template.New("test").Funcs(funcMap)
```

### 5. Handle Whitespace Carefully

```go
// Use {{- and -}} to control whitespace
{{- .Value -}}

// Or format template for readability
{{range .Items}}
  - {{.Name}}
{{end}}
```

### 6. Separate Template Files

```go
// Load from files
tmpl, err := template.ParseFiles(
    "templates/layout.tmpl",
    "templates/header.tmpl",
    "templates/footer.tmpl",
)
```

### 7. Document Template Data Structures

```go
// UserTemplate expects this data structure:
//   {
//     Name: string
//     Email: string
//     Age: int
//     IsActive: bool
//   }
const UserTemplate = `...`
```

### 8. Test Templates

```go
func TestUserTemplate(t *testing.T) {
    tmpl := template.Must(template.New("user").Parse(UserTemplate))

    var buf bytes.Buffer
    data := User{Name: "Alice", Age: 30}

    err := tmpl.Execute(&buf, data)
    if err != nil {
        t.Fatal(err)
    }

    if !strings.Contains(buf.String(), "Alice") {
        t.Error("Expected name in output")
    }
}
```

### 9. Use Option for Strict Checking

```go
tmpl := template.New("test").Option("missingkey=error")
```

### 10. Don't Put Complex Logic in Templates

```go
// Bad: Complex logic in template
{{if and (gt .Age 18) (eq .Country "US") (not .Banned)}}

// Good: Prepare data in Go code
data := struct {
    CanAccess bool
}{
    CanAccess: user.Age > 18 && user.Country == "US" && !user.Banned,
}

{{if .CanAccess}}
```

## Summary

- Use `text/template` for **generating text output**
- Templates use `{{` and `}}` for **actions**
- **Dot (.)** represents the current data context
- **Pipelines** chain operations with `|`
- **Control structures**: `if`, `range`, `with`
- **Define templates** with `{{define}}` and use with `{{template}}`
- Add **custom functions** with `FuncMap`
- **Parse once**, execute many times for performance
- Use **template.Must** for static templates
- Handle **errors** during both parsing and execution

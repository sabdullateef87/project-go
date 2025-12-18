# HTML Templates in Go

The `html/template` package provides data-driven templates for generating HTML output with automatic contextual escaping to prevent XSS attacks. It has the same interface as `text/template` but is designed specifically for HTML.

## Table of Contents

- [Overview](#overview)
- [Basic Usage](#basic-usage)
- [Automatic Escaping](#automatic-escaping)
- [Context-Aware Escaping](#context-aware-escaping)
- [Template Syntax](#template-syntax)
- [HTML Specific Features](#html-specific-features)
- [Template Composition](#template-composition)
- [Custom Functions](#custom-functions)
- [Working with Forms](#working-with-forms)
- [Common Patterns](#common-patterns)
- [Security Considerations](#security-considerations)
- [Performance](#performance)
- [Best Practices](#best-practices)

## Overview

The `html/template` package:

- **Automatic escaping** - Prevents XSS attacks
- **Context-aware** - Different escaping for HTML, JS, CSS, URLs
- **Same syntax as text/template** - Easy to learn
- **Type-safe** - Compile-time template parsing
- **Composable** - Template inheritance and partials

```go
import "html/template"
```

### Key Differences from text/template

1. **Automatic HTML escaping** - Special characters are escaped
2. **Context-aware** - Escaping changes based on context (HTML, JS, CSS, URL)
3. **Security focus** - Designed to prevent injection attacks

## Basic Usage

### Simple Template

```go
package main

import (
    "html/template"
    "os"
)

func main() {
    tmpl := template.Must(template.New("test").Parse(`
        <h1>Hello, {{.}}!</h1>
    `))

    tmpl.Execute(os.Stdout, "World")
    // Output: <h1>Hello, World!</h1>
}
```

### Template with Struct

```go
type Page struct {
    Title   string
    Heading string
    Content string
}

func main() {
    tmpl := template.Must(template.New("page").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Heading}}</h1>
    <p>{{.Content}}</p>
</body>
</html>
    `))

    page := Page{
        Title:   "My Page",
        Heading: "Welcome",
        Content: "Hello, World!",
    }

    tmpl.Execute(os.Stdout, page)
}
```

### HTTP Handler Example

```go
func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("template.html"))

    data := struct {
        Title string
        Items []string
    }{
        Title: "My Page",
        Items: []string{"Item 1", "Item 2", "Item 3"},
    }

    tmpl.Execute(w, data)
}
```

## Automatic Escaping

HTML templates automatically escape dangerous content.

### HTML Escaping

```go
type Data struct {
    UserInput string
}

tmpl := template.Must(template.New("test").Parse(`
    <p>{{.UserInput}}</p>
`))

data := Data{
    UserInput: "<script>alert('XSS')</script>",
}

tmpl.Execute(os.Stdout, data)
// Output: <p>&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;</p>
// Safe: Script tags are escaped
```

### Escaping Special Characters

```go
input := `<div>Hello & "Welcome" to 'Go'</div>`
tmpl := template.Must(template.New("test").Parse(`<p>{{.}}</p>`))
tmpl.Execute(os.Stdout, input)
// Output: <p>&lt;div&gt;Hello &amp; &#34;Welcome&#34; to &#39;Go&#39;&lt;/div&gt;</p>
```

### Safe HTML with template.HTML

When you have trusted HTML that should not be escaped:

```go
type Data struct {
    SafeHTML template.HTML
}

tmpl := template.Must(template.New("test").Parse(`
    <div>{{.SafeHTML}}</div>
`))

data := Data{
    SafeHTML: template.HTML("<strong>Bold text</strong>"),
}

tmpl.Execute(os.Stdout, data)
// Output: <div><strong>Bold text</strong></div>
// HTML is not escaped
```

**Warning**: Only use `template.HTML` with trusted content!

## Context-Aware Escaping

Go templates escape differently based on where the value appears.

### HTML Context

```go
tmpl := template.Must(template.New("test").Parse(`
    <div>{{.Content}}</div>
`))

data := map[string]string{
    "Content": "<script>alert('XSS')</script>",
}

tmpl.Execute(os.Stdout, data)
// Output: <div>&lt;script&gt;alert(&#39;XSS&#39;)&lt;/script&gt;</div>
```

### JavaScript Context

```go
tmpl := template.Must(template.New("test").Parse(`
    <script>
        var message = "{{.Message}}";
    </script>
`))

data := map[string]string{
    "Message": `"; alert("XSS"); var x="`,
}

tmpl.Execute(os.Stdout, data)
// Output: <script>
//     var message = "\"; alert(\"XSS\"); var x=\"";
// </script>
// Properly escaped for JavaScript context
```

### URL Context

```go
tmpl := template.Must(template.New("test").Parse(`
    <a href="/search?q={{.Query}}">Search</a>
`))

data := map[string]string{
    "Query": "hello world & more",
}

tmpl.Execute(os.Stdout, data)
// Output: <a href="/search?q=hello%20world%20%26%20more">Search</a>
```

### CSS Context

```go
tmpl := template.Must(template.New("test").Parse(`
    <div style="color: {{.Color}}">Text</div>
`))

data := map[string]string{
    "Color": "red; background: url(javascript:alert('XSS'))",
}

tmpl.Execute(os.Stdout, data)
// Dangerous CSS is sanitized
```

### Safe Types for Different Contexts

```go
type Data struct {
    HTML template.HTML        // Safe HTML content
    CSS  template.CSS         // Safe CSS content
    JS   template.JS          // Safe JavaScript content
    URL  template.URL         // Safe URL content
    Attr template.HTMLAttr    // Safe HTML attribute
    JSStr template.JSStr      // Safe JS string
}
```

## Template Syntax

Same as `text/template` - see [text-templates.md](text-templates.md) for full syntax.

### Quick Reference

```go
{{/* Comment */}}
{{.}}                    // Current value
{{.FieldName}}           // Field access
{{.Method}}              // Method call

{{if .Condition}}...{{end}}
{{range .Items}}...{{end}}
{{with .Value}}...{{end}}

{{$variable := .Value}}  // Variable

{{template "name" .}}    // Include template
{{define "name"}}...{{end}}  // Define template
```

## HTML Specific Features

### HTML Attributes

```go
type Link struct {
    URL   string
    Title string
    Class string
}

tmpl := template.Must(template.New("link").Parse(`
    <a href="{{.URL}}" title="{{.Title}}" class="{{.Class}}">Link</a>
`))

link := Link{
    URL:   "/page?id=1&ref=home",
    Title: "Go to \"Page 1\"",
    Class: "btn btn-primary",
}

tmpl.Execute(os.Stdout, link)
// Output: <a href="/page?id=1&amp;ref=home" title="Go to &#34;Page 1&#34;" class="btn btn-primary">Link</a>
```

### Conditional Attributes

```go
type Button struct {
    Text     string
    Disabled bool
    Class    string
}

tmpl := template.Must(template.New("button").Parse(`
    <button class="{{.Class}}"{{if .Disabled}} disabled{{end}}>
        {{.Text}}
    </button>
`))

button := Button{
    Text:     "Submit",
    Disabled: true,
    Class:    "btn",
}

tmpl.Execute(os.Stdout, button)
// Output: <button class="btn" disabled>Submit</button>
```

### Dynamic Classes

```go
type Item struct {
    Name   string
    Active bool
}

tmpl := template.Must(template.New("item").Parse(`
    <div class="item{{if .Active}} active{{end}}">
        {{.Name}}
    </div>
`))
```

### Data Attributes

```go
type Product struct {
    ID    int
    Name  string
    Price float64
}

tmpl := template.Must(template.New("product").Parse(`
    <div class="product" data-id="{{.ID}}" data-price="{{.Price}}">
        {{.Name}}
    </div>
`))
```

## Template Composition

### Layout Pattern

```go
// layout.html
{{define "layout"}}
<!DOCTYPE html>
<html>
<head>
    <title>{{template "title" .}}</title>
    {{template "styles" .}}
</head>
<body>
    {{template "header" .}}
    <main>
        {{template "content" .}}
    </main>
    {{template "footer" .}}
    {{template "scripts" .}}
</body>
</html>
{{end}}

{{define "styles"}}
<link rel="stylesheet" href="/css/main.css">
{{end}}

{{define "scripts"}}
<script src="/js/main.js"></script>
{{end}}
```

```go
// page.html
{{define "title"}}Home Page{{end}}

{{define "content"}}
<h1>Welcome</h1>
<p>This is the home page.</p>
{{end}}
```

```go
// Go code
func pageHandler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles(
        "templates/layout.html",
        "templates/page.html",
        "templates/header.html",
        "templates/footer.html",
    ))

    data := struct {
        PageTitle string
    }{
        PageTitle: "Home",
    }

    tmpl.ExecuteTemplate(w, "layout", data)
}
```

### Partials/Components

```go
// navbar.html
{{define "navbar"}}
<nav>
    <ul>
        {{range .NavItems}}
        <li><a href="{{.URL}}">{{.Text}}</a></li>
        {{end}}
    </ul>
</nav>
{{end}}
```

```go
// Use in multiple pages
{{template "navbar" .}}
```

### Block with Default Content

```go
{{define "layout"}}
<!DOCTYPE html>
<html>
<head>
    {{block "head" .}}
    <title>Default Title</title>
    {{end}}
</head>
<body>
    {{block "content" .}}
    <p>Default content</p>
    {{end}}
</body>
</html>
{{end}}
```

```go
// Override in child template
{{define "head"}}
<title>Custom Title</title>
<link rel="stylesheet" href="/custom.css">
{{end}}

{{define "content"}}
<h1>Custom Content</h1>
{{end}}
```

## Custom Functions

### Registering Functions

```go
funcMap := template.FuncMap{
    "upper": strings.ToUpper,
    "lower": strings.ToLower,
    "add": func(a, b int) int {
        return a + b
    },
}

tmpl := template.New("test").Funcs(funcMap)
```

### Useful HTML Functions

```go
funcMap := template.FuncMap{
    // Date formatting
    "formatDate": func(t time.Time) string {
        return t.Format("January 2, 2006")
    },

    // Currency formatting
    "currency": func(amount float64) string {
        return fmt.Sprintf("$%.2f", amount)
    },

    // Truncate text
    "truncate": func(s string, length int) string {
        if len(s) <= length {
            return s
        }
        return s[:length] + "..."
    },

    // Markdown to HTML (use with template.HTML)
    "markdown": func(s string) template.HTML {
        // Use markdown library
        html := markdownToHTML(s)
        return template.HTML(html)
    },

    // Pluralize
    "pluralize": func(count int, singular, plural string) string {
        if count == 1 {
            return singular
        }
        return plural
    },

    // Active class for navigation
    "activeClass": func(current, page string) string {
        if current == page {
            return "active"
        }
        return ""
    },
}
```

### Usage in Templates

```go
tmpl := `
<p>Price: {{.Price | currency}}</p>
<p>Posted: {{.Date | formatDate}}</p>
<p>{{.Description | truncate 100}}</p>
<p>{{.Count}} {{pluralize .Count "item" "items"}}</p>
<a class="{{activeClass .CurrentPage "home"}}">Home</a>
`
```

## Working with Forms

### Rendering Forms

```go
type UserForm struct {
    Name     string
    Email    string
    Age      int
    Country  string
    Subscribe bool
}

tmpl := template.Must(template.New("form").Parse(`
<form method="POST" action="/submit">
    <label>Name:
        <input type="text" name="name" value="{{.Name}}">
    </label>

    <label>Email:
        <input type="email" name="email" value="{{.Email}}">
    </label>

    <label>Age:
        <input type="number" name="age" value="{{.Age}}">
    </label>

    <label>Country:
        <select name="country">
            <option value="">Select...</option>
            <option value="US"{{if eq .Country "US"}} selected{{end}}>USA</option>
            <option value="UK"{{if eq .Country "UK"}} selected{{end}}>UK</option>
            <option value="CA"{{if eq .Country "CA"}} selected{{end}}>Canada</option>
        </select>
    </label>

    <label>
        <input type="checkbox" name="subscribe"{{if .Subscribe}} checked{{end}}>
        Subscribe to newsletter
    </label>

    <button type="submit">Submit</button>
</form>
`))
```

### Form with Validation Errors

```go
type FormData struct {
    Name   string
    Email  string
    Errors map[string]string
}

tmpl := template.Must(template.New("form").Parse(`
<form method="POST">
    <div>
        <label>Name:</label>
        <input type="text" name="name" value="{{.Name}}">
        {{if .Errors.Name}}
        <span class="error">{{.Errors.Name}}</span>
        {{end}}
    </div>

    <div>
        <label>Email:</label>
        <input type="email" name="email" value="{{.Email}}">
        {{if .Errors.Email}}
        <span class="error">{{.Errors.Email}}</span>
        {{end}}
    </div>

    <button type="submit">Submit</button>
</form>
`))
```

### CSRF Token

```go
type PageData struct {
    CSRFToken string
}

tmpl := template.Must(template.New("form").Parse(`
<form method="POST">
    <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
    <!-- form fields -->
    <button type="submit">Submit</button>
</form>
`))
```

## Common Patterns

### List with Alternating Rows

```go
type Item struct {
    Name string
}

tmpl := template.Must(template.New("list").Parse(`
<table>
{{range $index, $item := .Items}}
    <tr class="{{if even $index}}even{{else}}odd{{end}}">
        <td>{{$item.Name}}</td>
    </tr>
{{end}}
</table>
`))

// With custom function
funcMap := template.FuncMap{
    "even": func(n int) bool {
        return n%2 == 0
    },
}
```

### Breadcrumbs

```go
type Breadcrumb struct {
    Text string
    URL  string
}

tmpl := template.Must(template.New("breadcrumbs").Parse(`
<nav>
    {{range $index, $crumb := .Breadcrumbs}}
        {{if $index}} &gt; {{end}}
        {{if $crumb.URL}}
            <a href="{{$crumb.URL}}">{{$crumb.Text}}</a>
        {{else}}
            <span>{{$crumb.Text}}</span>
        {{end}}
    {{end}}
</nav>
`))
```

### Pagination

```go
type Pagination struct {
    CurrentPage int
    TotalPages  int
    BaseURL     string
}

tmpl := template.Must(template.New("pagination").Parse(`
<div class="pagination">
    {{if gt .CurrentPage 1}}
        <a href="{{.BaseURL}}?page={{add .CurrentPage -1}}">Previous</a>
    {{end}}

    {{range $page := sequence 1 .TotalPages}}
        {{if eq $page $.CurrentPage}}
            <span class="current">{{$page}}</span>
        {{else}}
            <a href="{{$.BaseURL}}?page={{$page}}">{{$page}}</a>
        {{end}}
    {{end}}

    {{if lt .CurrentPage .TotalPages}}
        <a href="{{.BaseURL}}?page={{add .CurrentPage 1}}">Next</a>
    {{end}}
</div>
`))

// Custom functions
funcMap := template.FuncMap{
    "add": func(a, b int) int { return a + b },
    "sequence": func(start, end int) []int {
        seq := make([]int, end-start+1)
        for i := range seq {
            seq[i] = start + i
        }
        return seq
    },
}
```

### Cards/Grid Layout

```go
type Card struct {
    Title       string
    Description string
    ImageURL    string
    Link        string
}

tmpl := template.Must(template.New("cards").Parse(`
<div class="card-grid">
{{range .Cards}}
    <div class="card">
        {{if .ImageURL}}
        <img src="{{.ImageURL}}" alt="{{.Title}}">
        {{end}}
        <h3>{{.Title}}</h3>
        <p>{{.Description}}</p>
        {{if .Link}}
        <a href="{{.Link}}">Read more</a>
        {{end}}
    </div>
{{end}}
</div>
`))
```

### Alert/Flash Messages

```go
type Message struct {
    Type    string // "success", "error", "warning", "info"
    Content string
}

tmpl := template.Must(template.New("messages").Parse(`
{{range .Messages}}
<div class="alert alert-{{.Type}}">
    {{.Content}}
</div>
{{end}}
`))
```

## Security Considerations

### Never Trust User Input

```go
// Bad: Using template.HTML with user input
type Data struct {
    UserHTML template.HTML // DANGEROUS!
}

data := Data{
    UserHTML: template.HTML(userInput), // XSS vulnerability
}

// Good: Let template escape automatically
type Data struct {
    UserText string
}

data := Data{
    UserText: userInput, // Safely escaped
}
```

### Sanitize Before Using Safe Types

```go
import "github.com/microcosm-cc/bluemonday"

func sanitizeHTML(input string) template.HTML {
    policy := bluemonday.UGCPolicy()
    safe := policy.Sanitize(input)
    return template.HTML(safe)
}
```

### Validate URLs

```go
func safeURL(input string) template.URL {
    u, err := url.Parse(input)
    if err != nil {
        return template.URL("")
    }

    // Only allow http and https
    if u.Scheme != "http" && u.Scheme != "https" {
        return template.URL("")
    }

    return template.URL(u.String())
}
```

### CSRF Protection

```go
// Generate token
token := generateCSRFToken()

// Store in session
session.Values["csrf_token"] = token

// Pass to template
data := struct {
    CSRFToken string
}{
    CSRFToken: token,
}

// In template
<form method="POST">
    <input type="hidden" name="csrf_token" value="{{.CSRFToken}}">
</form>

// Verify on submission
if r.FormValue("csrf_token") != session.Values["csrf_token"] {
    http.Error(w, "Invalid CSRF token", http.StatusForbidden)
    return
}
```

### Content Security Policy

```go
func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Security-Policy",
        "default-src 'self'; script-src 'self'; style-src 'self'")

    tmpl.Execute(w, data)
}
```

## Performance

### Parse Templates Once

```go
// Bad: Parse on every request
func handler(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("template.html"))
    tmpl.Execute(w, data)
}

// Good: Parse once at startup
var templates = template.Must(template.ParseGlob("templates/*.html"))

func handler(w http.ResponseWriter, r *http.Request) {
    templates.ExecuteTemplate(w, "template.html", data)
}
```

### Use Buffer for Complex Templates

```go
var bufferPool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

func handler(w http.ResponseWriter, r *http.Request) {
    buf := bufferPool.Get().(*bytes.Buffer)
    buf.Reset()
    defer bufferPool.Put(buf)

    if err := tmpl.Execute(buf, data); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    buf.WriteTo(w)
}
```

### Cache Rendered Output

```go
var cache = make(map[string][]byte)
var cacheMutex sync.RWMutex

func getCachedPage(key string, generate func() []byte) []byte {
    cacheMutex.RLock()
    if cached, ok := cache[key]; ok {
        cacheMutex.RUnlock()
        return cached
    }
    cacheMutex.RUnlock()

    cacheMutex.Lock()
    defer cacheMutex.Unlock()

    // Double-check
    if cached, ok := cache[key]; ok {
        return cached
    }

    result := generate()
    cache[key] = result
    return result
}
```

## Best Practices

### 1. Always Escape User Input

```go
// Good: Automatic escaping
<p>{{.UserInput}}</p>

// Bad: Using template.HTML with user input
<p>{{.UserInput | template.HTML}}</p>
```

### 2. Use Semantic HTML

```go
// Good
<article>
    <header>
        <h1>{{.Title}}</h1>
    </header>
    <main>{{.Content}}</main>
</article>

// Less semantic
<div>
    <div><span>{{.Title}}</span></div>
    <div>{{.Content}}</div>
</div>
```

### 3. Separate Layout from Content

```go
// layout.html - structure
// page.html - content
// partials/ - reusable components
```

### 4. Use Named Templates

```go
{{define "button"}}
<button class="btn">{{.}}</button>
{{end}}

{{template "button" "Click Me"}}
```

### 5. Validate Template Data

```go
func handler(w http.ResponseWriter, r *http.Request) {
    data := prepareData()

    if err := validateData(data); err != nil {
        http.Error(w, "Invalid data", http.StatusInternalServerError)
        return
    }

    tmpl.Execute(w, data)
}
```

### 6. Handle Template Errors

```go
func handler(w http.ResponseWriter, r *http.Request) {
    var buf bytes.Buffer

    if err := tmpl.Execute(&buf, data); err != nil {
        log.Printf("Template error: %v", err)
        http.Error(w, "Internal error", http.StatusInternalServerError)
        return
    }

    buf.WriteTo(w)
}
```

### 7. Use Descriptive Template Names

```go
// Good
{{template "user-profile" .}}
{{template "navigation-menu" .}}

// Less clear
{{template "temp1" .}}
{{template "part2" .}}
```

### 8. Document Expected Data Structure

```go
// This template expects:
// type PageData struct {
//     Title string
//     User  User
//     Posts []Post
// }
{{define "page"}}
```

### 9. Test Templates

```go
func TestTemplate(t *testing.T) {
    tmpl := template.Must(template.ParseFiles("template.html"))

    var buf bytes.Buffer
    data := PageData{Title: "Test"}

    err := tmpl.Execute(&buf, data)
    if err != nil {
        t.Fatal(err)
    }

    html := buf.String()
    if !strings.Contains(html, "Test") {
        t.Error("Expected title in output")
    }
}
```

### 10. Use Template Inheritance

```go
// base.html
{{define "base"}}
<!DOCTYPE html>
<html>
<head>{{block "head" .}}{{end}}</head>
<body>{{block "body" .}}{{end}}</body>
</html>
{{end}}

// page.html
{{template "base" .}}
{{define "head"}}<title>{{.Title}}</title>{{end}}
{{define "body"}}<h1>{{.Heading}}</h1>{{end}}
```

## Summary

- Use `html/template` for **generating HTML** (not text/template)
- **Automatic escaping** prevents XSS attacks
- Escaping is **context-aware** (HTML, JS, CSS, URL)
- Use `template.HTML` **only with trusted content**
- Same **syntax as text/template**
- Use **layout patterns** for template composition
- Add **custom functions** for formatting
- Always **validate and sanitize** user input
- **Parse templates once** at startup for performance
- Use **CSRF tokens** for form protection

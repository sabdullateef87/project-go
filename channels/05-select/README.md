# Select

`select` multiplexes communication on multiple channels. It is central to building responsive, cancellable concurrency.

Example:

```go
select {
case res := <-execCh:
    fmt.Println("exec", res.ID)
case <-quit:
    fmt.Println("shutting down")
}
```

Use `default` for non-blocking operations and combine `select` with `context.Done()` for cancellations.

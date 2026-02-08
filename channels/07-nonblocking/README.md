# Non-Blocking Channel Operations

Use `select` with `default` to attempt send/receive without blocking.

```go
select {
case ch <- v:
    // sent
default:
    // couldn't send; handle backpressure/drop
}
```

Useful for metrics, logging, or best-effort publishes where you don't want producers to block.

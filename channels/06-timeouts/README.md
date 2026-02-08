# Timeouts

Use `time.After` or `time.NewTimer` in `select` to implement timeouts.

```go
select {
case res := <-longOpCh:
    fmt.Println("result", res)
case <-time.After(2 * time.Second):
    fmt.Println("timeout")
}
```

Finance POV: fallback to cached price or mark stale when external pricing doesn't respond within SLA.

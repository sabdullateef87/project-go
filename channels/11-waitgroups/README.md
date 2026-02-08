# WaitGroups

`sync.WaitGroup` coordinates goroutine lifetimes. Use `Add`, `Done`, `Wait` to synchronize.

```go
var wg sync.WaitGroup
wg.Add(1)
go func(){ defer wg.Done(); /* work */ }()
wg.Wait()
```

Combine `WaitGroup` with closing channels to signal result completion.

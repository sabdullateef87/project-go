# Atomic Counters

Use `sync/atomic` for lock-free counters.

```go
import "sync/atomic"
var processed uint64
atomic.AddUint64(&processed, 1)
_ = atomic.LoadUint64(&processed)
```

Useful for metrics like orders processed where low-overhead counters matter.

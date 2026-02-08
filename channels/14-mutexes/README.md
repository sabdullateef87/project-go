# Mutexes

`sync.Mutex` and `sync.RWMutex` protect shared mutable state.

Example: ledger protecting balances

```go
type Ledger struct{
    mu sync.RWMutex
    balances map[string]float64
}

func (l *Ledger) Get(id string) float64 {
    l.mu.RLock(); defer l.mu.RUnlock(); return l.balances[id]
}

func (l *Ledger) Add(id string, amt float64) {
    l.mu.Lock(); defer l.mu.Unlock(); l.balances[id] += amt
}
```

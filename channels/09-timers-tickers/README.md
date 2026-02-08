# Timers & Tickers

`time.NewTimer` fires once; `time.NewTicker` fires periodically. Always `Stop()` when no longer needed.

Ticker example for periodic snapshots:

```go
ticker := time.NewTicker(time.Second)
defer ticker.Stop()
for range ticker.C {
    // snapshot P&L or risk metrics
}
```

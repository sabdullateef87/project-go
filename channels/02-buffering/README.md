# Channel Buffering

Buffered channels allow sends without immediate receiver until the buffer is full. Useful to absorb bursts.

```go
ch := make(chan int, 3)
ch <- 1
ch <- 2
ch <- 3
```

Finance POV: model an ingress queue for trade messages from multiple market feeds. Tune buffer size to expected burstiness.

Example (ingest):

```go
package main

import (
    "fmt"
)

type Trade struct{ ID string; Amount float64 }

func ingest(feed chan<- Trade) {
    for i := 0; i < 100; i++ {
        feed <- Trade{ID: fmt.Sprintf("T%d", i), Amount: float64(i) * 100}
    }
    close(feed)
}

func main() {
    feed := make(chan Trade, 50)
    go ingest(feed)
    for t := range feed { _ = t }
}
```

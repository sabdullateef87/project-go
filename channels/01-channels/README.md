# Channels

Channels are typed conduits through which goroutines communicate. They can be unbuffered (synchronous) or buffered (asynchronous).

Example (unbuffered):

```go
package main

import "fmt"

func main() {
    ch := make(chan string)
    go func() { ch <- "hello from goroutine" }()
    fmt.Println(<-ch)
}
```

Finance POV: channels model event streams such as market data feeds or trade events.

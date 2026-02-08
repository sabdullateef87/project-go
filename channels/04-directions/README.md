# Channel Directions

Use directional channel types in function signatures to clarify intent.

```go
func producer(out chan<- int) { out <- 42 }
func consumer(in <-chan int)  { fmt.Println(<-in) }
```

This prevents accidental misuse like receiving on a send-only channel.

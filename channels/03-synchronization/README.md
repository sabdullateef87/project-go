# Channel Synchronization

Unbuffered channels are synchronization points — send and receive rendezvous.

```go
done := make(chan struct{})
go func(){ /* work */ close(done) }()
<-done // wait for goroutine
```

Use channels to signal completion, readiness, or to sequence stages in pipelines.

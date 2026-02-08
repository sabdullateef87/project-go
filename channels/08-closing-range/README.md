# Closing Channels & Range

Senders should close channels to indicate no more values. Receivers detect closure and `range` ends.

```go
ch := make(chan int)
go func(){ defer close(ch); for i:=0;i<5;i++{ ch<-i }}()
for v := range ch { println(v) }
```

Do not close a channel from multiple senders; only the owner should close.

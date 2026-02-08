# Channels (Go)

This document provides a deep dive into Go channels with extensive examples. All runnable samples now use a **payments** domain (auth, settlement, risk) so you can copy/paste and run them. Topics covered:

- Channels
- Channel Buffering
- Channel Synchronization
- Channel Directions
- Select
- Timeouts
- Non-Blocking Channel Operations
- Closing Channels
- Range over Channels
- Timers
- Tickers
- Worker Pools
- WaitGroups
- Rate Limiting
- Atomic Counters
- Mutexes
- Stateful Goroutines

Each section includes explanation and elaborate, practical examples. Examples use idiomatic Go and are small enough to copy into a `main.go` for experimentation.

## Index of runnable examples (payments domain)

- [01-channels](01-channels) — basic producer/consumer for payment events
- [02-buffering](02-buffering) — buffered ingress queue for bursty payment submissions
- [03-synchronization](03-synchronization) — stage orchestration for startup steps
- [04-directions](04-directions) — producer/consumer with directional channels
- [05-select](05-select) — merge two price/fee feeds with timeout
- [06-timeouts](06-timeouts) — gateway call with timeout fallback
- [07-nonblocking](07-nonblocking) — best-effort logging without blocking
- [08-closing-range](08-closing-range) — producer signals completion
- [09-timers-tickers](09-timers-tickers) — periodic settlement snapshots
- [10-worker-pools](10-worker-pools) — payment validation workers
- [11-waitgroups](11-waitgroups) — coordinate worker lifetimes
- [12-rate-limiting](12-rate-limiting) — token bucket for gateway calls
- [13-atomic-counters](13-atomic-counters) — fast metrics increments
- [14-mutexes](14-mutexes) — shared ledger with locks
- [15-stateful-goroutines](15-stateful-goroutines) — account actor (serializes balance changes)

How to run an example (from repo root):

```sh
go run ./channels/10-worker-pools
```

---

**Channels**: Channels are typed conduits through which goroutines communicate. They can be unbuffered (synchronous) or buffered (asynchronous).

Example: basic unbuffered channel

```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello from goroutine"
	}()

	msg := <-ch
	fmt.Println(msg)
}
```

---

**Channel Buffering**: A buffered channel allows sending without immediate receiver until the buffer is full.

```go
ch := make(chan int, 3)
ch <- 1 // does not block until buffer full
ch <- 2
ch <- 3
// next send would block until a receive.
```

Finance POV: buffering can model an ingress queue of trade messages where producers (market feeds) are bursty and workers drain the buffer.

```go
// simplified trade ingestion
type Trade struct{ ID string; Amount float64 }

func ingest(feed chan<- Trade) {
	for i := 0; i < 100; i++ {
		feed <- Trade{ID: fmt.Sprintf("T%d", i), Amount: float64(i) * 100}
	}
	close(feed)
}

func main() {
	feed := make(chan Trade, 50) // buffer to absorb bursts
	go ingest(feed)
	for t := range feed {
		// process trade
		_ = t
	}
}
```

---

**Channel Synchronization**: Channels can enforce ordering and synchronization. An unbuffered channel is a synchronization point.

```go
done := make(chan struct{})
go func() {
	// do work
	close(done) // signal completion
}()
<-done // wait
```

---

**Channel Directions**: Specify direction in function signatures for clarity and safety.

```go
func producer(out chan<- int) { out <- 42 }
func consumer(in <-chan int)  { fmt.Println(<-in) }
```

---

**Select**: Multiplex multiple channel operations. Useful for handling multiple inputs, cancellations, and timeouts.

Example: select over trade channels and a quit signal

```go
type ExecResult struct{ ID string; Err error }

func main() {
	execCh := make(chan ExecResult)
	quit := make(chan struct{})

	go func() {
		// simulate result
		execCh <- ExecResult{ID: "order-1"}
	}()

	select {
	case res := <-execCh:
		fmt.Println("exec", res.ID)
	case <-quit:
		fmt.Println("shutting down")
	}
}
```

**Select with default (non-blocking)**

```go
select {
case msg := <-ch:
	fmt.Println("received", msg)
default:
	fmt.Println("no message available, continue")
}
```

---

**Timeouts**: Use `time.After` (or `time.NewTimer`) inside `select` to implement timeouts.

```go
select {
case res := <-longOpCh:
	fmt.Println("result", res)
case <-time.After(2 * time.Second):
	fmt.Println("timeout")
}
```

Finance POV: if an external pricing service doesn't respond within SLA, failover to cached price or mark as stale.

---

**Non-Blocking Channel Operations**: Use `select` with `default` to try sends/receives without blocking.

```go
// try send without blocking
select {
case ch <- v:
	// sent
default:
	// couldn't send; drop or handle backpressure
}
```

---

**Closing Channels**: `close(ch)` indicates no more values will be sent. Receivers can detect closure.

Important: only the sender should close a channel. Closing a channel lets `range` finish.

```go
ch := make(chan int)
go func() {
	defer close(ch)
	for i := 0; i < 5; i++ { ch <- i }
}()
for v := range ch { fmt.Println(v) }
```

---

**Range over Channels**: `for v := range ch` iterates until `ch` is closed.

---

**Timers**: `time.NewTimer` returns a timer; useful when you need to stop or reset.

```go
t := time.NewTimer(1 * time.Second)
select {
case <-t.C:
	fmt.Println("timer fired")
}
```

**Tickers**: `time.NewTicker` for periodic events (remember to `Stop()` when done).

```go
t := time.NewTicker(time.Second)
defer t.Stop()
for range t.C {
	// periodic work
}
```

Finance POV: Tickers can drive periodic risk calculations or P&L snapshots.

---

**Worker Pools**: Common concurrency pattern — spawn N workers reading from a jobs channel and sending results to a results channel.

Finance example: process trade validations concurrently and aggregate results.

```go
package main

import (
	"fmt"
	"sync"
)

type Trade struct{ ID string; Amount float64 }

func worker(id int, jobs <-chan Trade, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for t := range jobs {
		// simulate validation
		results <- fmt.Sprintf("worker %d validated %s", id, t.ID)
	}
}

func main() {
	jobs := make(chan Trade)
	results := make(chan string)

	var wg sync.WaitGroup
	numWorkers := 4

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// producer
	go func() {
		for i := 0; i < 10; i++ {
			jobs <- Trade{ID: fmt.Sprintf("T%d", i), Amount: float64(i) * 100}
		}
		close(jobs)
	}()

	// collector
	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results { fmt.Println(r) }
}
```

Notes: Using `WaitGroup` in combination with closing the results channel is a common pattern.

---

**WaitGroups**: Coordinate goroutine lifetimes. `wg.Add`, `wg.Done`, `wg.Wait`.

```go
var wg sync.WaitGroup
wg.Add(1)
go func(){ defer wg.Done(); /* work */ }()
wg.Wait()
```

---

**Rate Limiting**: Implement with a token bucket using a buffered channel or `time.Ticker`.

Example token-bucket limiter (per-second) using ticker:

```go
tokens := make(chan struct{}, 5) // capacity = burst
// fill bucket
for i := 0; i < cap(tokens); i++ { tokens <- struct{}{} }

go func() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for range ticker.C {
		select {
		case tokens <- struct{}{}:
		default:
		}
	}
}()

// use token to perform action
select {
case <-tokens:
	// allowed
default:
	// rate limit hit
}
```

Finance POV: Ensure order entry to an external venue doesn't exceed API rate limits.

---

**Atomic Counters**: Use `sync/atomic` for simple counters without mutexes.

```go
import "sync/atomic"

var processed uint64
atomic.AddUint64(&processed, 1)
v := atomic.LoadUint64(&processed)
```

Use atomics for metrics like `ordersProcessed` or `failedAttempts`.

---

**Mutexes**: `sync.Mutex` or `sync.RWMutex` protect shared mutable state.

Example: protecting a map of account balances

```go
type Ledger struct{
	mu sync.RWMutex
	balances map[string]float64
}

func (l *Ledger) Get(id string) float64 {
	l.mu.RLock()
	defer l.mu.RUnlock()
	return l.balances[id]
}

func (l *Ledger) Add(id string, amt float64) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.balances[id] += amt
}
```

---

**Stateful Goroutines (Actor pattern)**: Run a goroutine that serializes access to some state by handling requests over channels. This avoids explicit locking.

Example: account actor that serially processes deposits and balance queries

```go
package main

import (
	"fmt"
)

type acctCmd struct{
	op string
	amt float64
	resp chan float64
}

func accountActor(start float64) chan acctCmd {
	cmds := make(chan acctCmd)
	go func() {
		bal := start
		for c := range cmds {
			switch c.op {
			case "deposit":
				bal += c.amt
				c.resp <- bal
			case "balance":
				c.resp <- bal
			}
		}
	}()
	return cmds
}

func main() {
	a := accountActor(1000)
	resp := make(chan float64)
	a <- acctCmd{op: "deposit", amt: 250, resp: resp}
	fmt.Println("after deposit", <-resp)
	a <- acctCmd{op: "balance", resp: resp}
	fmt.Println("balance", <-resp)
	close(a)
}
```

This pattern is powerful in finance where operations on an account must be strictly serialized to avoid race conditions and inconsistent balances.

---

Notes and Best Practices

- Prefer channels and goroutines for communication, and use mutexes for shared memory when necessary.
- Close channels only from the sending side.
- For high-throughput systems (e.g., order matching), measure and tune buffer sizes and worker counts.
- Use contexts (`context.Context`) to propagate cancellation and deadlines across goroutines—combine `context.Done()` and channels in `select`.

Further steps I can take if you want:

- Extract examples into runnable `.go` files under `channels/examples/`.
- Add tests or small bench programs demonstrating throughput and latencies.
- Add diagrams or ASCII sequence diagrams showing message flow.

---

If you'd like, I will extract these examples into separate runnable files and add a short README with how to run them. Which of the follow-up tasks would you like next?


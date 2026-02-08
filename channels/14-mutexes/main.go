package main

import (
	"fmt"
	"sync"
)

// Ledger protects shared balances with an RWMutex.
// In a payments system this could back a simple in-memory wallet.
type Ledger struct {
	mu       sync.RWMutex
	balances map[string]float64
}

func NewLedger() *Ledger {
	return &Ledger{balances: make(map[string]float64)}
}

func main() {
	l := NewLedger()
	var wg sync.WaitGroup

	// simulate concurrent balance updates
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			acct := fmt.Sprintf("A%d", i%2) // two accounts updated concurrently
			l.Add(acct, 100.0)
		}(i)
	}

	wg.Wait()
	fmt.Println("A0 balance", l.Get("A0"))
	fmt.Println("A1 balance", l.Get("A1"))
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

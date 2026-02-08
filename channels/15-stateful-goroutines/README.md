# Stateful Goroutines (Actor Pattern)

Run a goroutine that serializes state changes by handling commands over a channel. This avoids explicit locks.

Example: account actor (see `main.go`).

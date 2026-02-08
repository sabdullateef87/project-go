package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	stop := time.After(2 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Println("settlement snapshot at", time.Now().Format(time.StampMilli))
		case <-stop:
			fmt.Println("stopping snapshots")
			return
		}
	}
}

package main

import "fmt"

// paymentCmd is an actor message for a payment ledger.
// ops: "auth" (reserve funds), "capture" (settle), "refund", "balance".
type paymentCmd struct {
	op       string
	amount   float64
	respChan chan float64
}

func main() {
	ledger := ledgerActor(1000)
	resp := make(chan float64)

	ledger <- paymentCmd{op: "auth", amount: 200, respChan: resp}
	fmt.Println("reserved after auth:", <-resp)

	ledger <- paymentCmd{op: "capture", amount: 150, respChan: resp}
	fmt.Println("balance after capture:", <-resp)

	ledger <- paymentCmd{op: "refund", amount: 50, respChan: resp}
	fmt.Println("balance after refund:", <-resp)

	ledger <- paymentCmd{op: "balance", respChan: resp}
	fmt.Println("final balance:", <-resp)

	close(ledger)
}

func ledgerActor(start float64) chan paymentCmd {
	cmds := make(chan paymentCmd)
	go func() {
		balance := start
		reserved := 0.0
		for c := range cmds {
			switch c.op {
			case "auth":
				reserved += c.amount
				c.respChan <- reserved
			case "capture":
				// move from reserved to settled balance
				if c.amount <= reserved {
					reserved -= c.amount
					balance += c.amount
					c.respChan <- balance
				} else {
					c.respChan <- balance // insufficient reserve
				}
			case "refund":
				if c.amount <= balance {
					balance -= c.amount
					c.respChan <- balance
				} else {
					c.respChan <- balance // decline refund beyond balance
				}
			case "balance":
				c.respChan <- balance
			}
		}
	}()
	return cmds
}

package main

import (
	"fmt"

	"github.com/panchoh/history-tree/histree"
)

func main() {
	ht := histree.NewHisTree()

	for i := 0; i < 5; i++ {
		event := fmt.Sprintf("Event #%d", i)
		fmt.Println("\nAdding event", i)
		commitment := ht.Add(
			&histree.Event{
				Value: []byte(event),
			},
		)
		fmt.Printf(
			"Added event '%s', and received commitment with Version '%d' and Digest '%v'\n",
			event,
			commitment.Version,
			commitment.Digest,
		)
	}
}

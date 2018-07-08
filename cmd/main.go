package main

import (
	"fmt"

	"github.com/panchoh/history-tree/histree"
)

func main() {
	ht := histree.NewHisTree()

	for i := 0; i < 10; i++ {
		event := fmt.Sprintf("Event #%d", i)
		commitment := ht.Add(
			&histree.Event{
				Value: []byte(event),
			},
		)
		fmt.Printf("Added event '%s', and received commitment '%v'\n", event, commitment)
	}
}

package main

import "fmt"

var broadcast = make(chan message)

func forwardBroadcast() {
	for {
		msg := <-broadcast
		for _, cli := range clients {
			if err := cli.WriteJSON(msg); err != nil {
				fmt.Printf("Error: %v\n", err)
				_ = cli.Close()
			}
		}
	}
}

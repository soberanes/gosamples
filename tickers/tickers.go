package main

import (
	"fmt"
	"time"
)

// Timer ar for when you want to do something once in the
// future - Tickers are for when you want to do something
// repeatedly at regular intervals.
func main() {

	// Tickers use a similar mechanism to timers: a channel that is sent
	// values. Here we'll use the range builtin on the channel to
	// iterate over the values as they arrive every 500ms
	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	// Tickers can be stopped like timers. Once a ticker is stopped it
	// won't receive any more values on its channel. We'll stop ours
	// after 1600ms
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}

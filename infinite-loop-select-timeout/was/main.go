package main

import (
	"flag"
	"log"
	"time"

	ping "github.com/sparrc/go-ping"
)

var (
	durations = flag.Int("duration", 10, "The duration time")
)

func main() {
	ticker := time.NewTicker(time.Duration(*durations) * time.Second)
	defer ticker.Stop()

	log.Printf("Start ping, it will stop after %v seconds", *durations)
	for {
		select {
		case <-ticker.C:
			log.Print("Goroutine timeout")
			return
		default:
			// Ping a 8.8.8.8
			pinger, err := ping.NewPinger("8.8.8.8")
			if err != nil {
				panic(err)
			}
			pinger.Count = 3
			pinger.Run()                 // blocks until finished
			stats := pinger.Statistics() // get send/receive/rtt stats
			log.Print(stats)

			// Sleep 3 seconds
			time.Sleep(3 * time.Second)

			// Ping a 8.8.4.4
			pinger, err = ping.NewPinger("8.8.4.4")
			if err != nil {
				panic(err)
			}
			pinger.Count = 3
			pinger.Run()                // blocks until finished
			stats = pinger.Statistics() // get send/receive/rtt stats
			log.Print(stats)
		}
	}
}

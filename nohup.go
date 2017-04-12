package main

import "os"

// catch nohup goodness
func nohup(sigs chan os.Signal) {
	for {
		select {
		// catch the hangup signal for as long as kj is running
		case <-sigs:
			continue
		}
	}
}

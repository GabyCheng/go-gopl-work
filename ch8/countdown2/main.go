package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// ...create abort channel...
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		// Do nothing.
		fmt.Println("等时间")
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}

}

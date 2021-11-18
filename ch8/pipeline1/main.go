package main

import "fmt"

func main() {
	// Squarer
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break // channel was closed and drained
			}
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	for {
		fmt.Println(<-squares)
	}
}

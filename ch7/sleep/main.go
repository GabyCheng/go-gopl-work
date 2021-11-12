package main

import (
	"flag"
	"fmt"
	"os"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")

func main() {

	var x interface{} = []int{1, 2, 3}

	fmt.Printf("%T\n", x) // "<nil>"\
	os.Exit(1)
	flag.Parse()
	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}

package main

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestJoin(t *testing.T) {
	start := time.Now()
	for i := 0; i < 1e7; i++ {
		input := []string{"Welcome", "To", "China"}
		result := strings.Join(input, " ")
		if result != "Welcome To China" {
			t.Error("Unexcepted result:" + result)
		}
	}
	fmt.Printf("echo2: %fs\n", time.Since(start).Seconds())
}

func TestPlus(t *testing.T) {
	start := time.Now()
	for i := 0; i < 1e7; i++ {
		input := []string{"Welcome", "To", "China"}
		var s, sep string
		for j := 0; j < len(input); j++ {
			s += sep + input[j]
			sep = " "
		}
		if s != "Welcome To China" {
			t.Error("Unexcepted result:" + s)
		}
	}
	fmt.Printf("echo2: %fs\n", time.Since(start).Seconds())
}

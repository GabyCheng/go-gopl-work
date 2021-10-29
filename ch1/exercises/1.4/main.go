package main

import (
	"bufio"
	"fmt"
	"os"
)

type LnFile struct {
	Count     int
	Filenames []string
}

func main() {
	counts := make(map[string]*LnFile)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}

	for line, n := range counts {
		if n.Count > 1 {
			fmt.Printf("%d\t%s\t%s\n", n.Count, n.Filenames, line)
		}
	}

}

func countLines(f *os.File, counts map[string]*LnFile) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		if text == "end" {
			break
		}

		_, ok := counts[text]
		if ok {
			counts[text].Count++
			counts[text].Filenames = append(counts[text].Filenames, f.Name())
		} else {
			counts[text] = new(LnFile)
			counts[text].Count = 1
			counts[text].Filenames = append(counts[text].Filenames, f.Name())
		}

	}
}

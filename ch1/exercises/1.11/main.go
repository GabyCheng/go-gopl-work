// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, filename := range os.Args[1:] {
		result := getUrls(filename)
		for _, url := range result {
			go fetch(url, ch) // start a goroutine
		}
		for range result {
			fmt.Println(<-ch) // receive from channel ch
		}
	}
	close(ch)
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func getUrls(filename string) []string {

	result := make([]string, 0, 100)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Sprint("read file error", filename, err)
	}

	for _, url := range strings.Split(string(data), "\r\n") {
		result = append(result, url)
	}

	return result
}

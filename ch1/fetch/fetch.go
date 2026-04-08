package fetch

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

// Prints the content found at a URL.
func Fetch() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}

// Exercise 1.7: Use io.Copy to copy the response body to os.Stdout without
// requiring a buffer large enough to hold the entire stream.
func FetchWithoutBuffer() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying from %s: %v\n", url, err)
		}
	}
}

// Exercise 1.8: Add the prefix http:// to each argument URL if it is missing.
func FetchWithMissingPrefix() {
	for _, url := range os.Args[1:] {
		const httpPrefix = "http://"
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying from %s: %v\n", url, err)
		}
	}
}

// Exercise 1.9: Also print the HTTP status code, found in resp.Status
func FetchWithStatusCode() {
	for _, url := range os.Args[1:] {
		const httpPrefix = "http://"
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("HTTP status code: %s\n", resp.Status)
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying from %s: %v\n", url, err)
		}
	}
}

// fetchAll fetches URLs in parallel and reports their times and sizes
func FetchAll() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go func(url string) {
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				ch <- fmt.Sprint(err) // send error message to channel ch
				return
			}
			nbytes, err := io.Copy(io.Discard, resp.Body)
			resp.Body.Close() // don't leak resources
			if err != nil {
				ch <- fmt.Sprintf("while reading %s: %v", url, err)
				return
			}
			secs := time.Since(start).Seconds()
			ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
		}(url)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// Exercise 1.10: print the output to a file to see if two consecutive requests
// return the same result
func FetchToFile() {
	start := time.Now()
	ch := make(chan string)
	for _, rawUrl := range os.Args[1:] {
		go func(rawUrl string) {
			start := time.Now()
			resp, err := http.Get(rawUrl)
			if err != nil {
				ch <- fmt.Sprint(err) // send error message to channel ch
				return
			}
			file, err := os.Create(url.PathEscape(rawUrl))
			if err != nil {
				ch <- fmt.Sprint(err)
				return
			}
			nbytes, err := io.Copy(file, resp.Body)
			resp.Body.Close() // don't leak resources
			if err != nil {
				ch <- fmt.Sprintf("while reading %s: %v", rawUrl, err)
				return
			}
			secs := time.Since(start).Seconds()
			ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, rawUrl)
		}(rawUrl)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

package examples

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func returnType(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("error: %s\n", err)
		}
	}(resp.Body)
	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s \n", url, ctype)
}

func siteSerial(urls []string) {
	for _, url := range urls {
		returnType(url)
	}
}

func sitesConcurrent(urls []string) {
	var wg sync.WaitGroup
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			returnType(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func returnTypeByChannel(url string, out chan string) {
	resp, err := http.Get(url)
	if err != nil {
		out <- fmt.Sprintf("error: %s", err)
		return
	}
	defer resp.Body.Close()
	ctype := resp.Header.Get("content-type")
	out <- fmt.Sprintf("%s -> %s", url, ctype)
}

func sitesConcurrentUsingChannel(urls []string) {
	ch := make(chan string)
	for _, url := range urls {
		go returnTypeByChannel(url, ch)
	}

	for range urls {
		out := <-ch
		fmt.Println(out)
	}
}

func MainGoroutines() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/ip",
	}

	start := time.Now()
	siteSerial(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrent(urls)
	fmt.Println(time.Since(start))

	start = time.Now()
	sitesConcurrentUsingChannel(urls)
	fmt.Println(time.Since(start))
}

package examples

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

var (
	urlTemplates = "http://s3.amazonaws.com/nyc-tlc/trip+data/%s_tripdata_2020-%02d.csv"
	colors       = []string{"yellow", "green"}
)

func donwloadSize(url string) (int, error) {
	req, err := http.NewRequest(http.MethodHead, url, nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf(resp.Status)
	}
	return strconv.Atoi(resp.Header.Get("Context-Length"))
}

func MainDownload() {
	start := time.Now()
	size := 0
	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplates, color, month)
			fmt.Println(url)
			n, err := donwloadSize(url)
			if err != nil {
				log.Fatal(err)
			}
			size += n
		}
	}

	duration := time.Since(start)
	fmt.Println(size, duration)
}

type result struct {
	url  string
	size int
	err  error
}

func sizeWorker(url string, ch chan result) {
	fmt.Println(url)
	res := result{url: url}
	res.size, res.err = donwloadSize(url)
	ch <- res
}

func MainDonwloadCurrency() {
	start := time.Now()
	ch := make(chan result)
	for month := 1; month <= 12; month++ {
		for _, color := range colors {
			url := fmt.Sprintf(urlTemplates, color, month)
			go sizeWorker(url, ch)
		}
	}

	size := 0
	for i := 0; i < len(colors)*12; i++ {
		r := <-ch
		if r.err != nil {
			log.Fatal(r.err)
		}
		size += r.size
	}

	duration := time.Since(start)
	fmt.Println(size, duration)
}

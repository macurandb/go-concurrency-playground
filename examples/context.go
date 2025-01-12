package examples

import (
	"context"
	"fmt"
	"time"
)

type Bid struct {
	AddURL string
	Price  float64
}

func bestBid(url string) Bid {
	time.Sleep(20 * time.Millisecond)
	return Bid{
		AddURL: url,
		Price:  0.05,
	}
}

var defaultBid = Bid{
	AddURL: "https://adsRus.com/default",
	Price:  0.01,
}

func findBind(ctx context.Context, url string) Bid {
	ch := make(chan Bid, 1)
	go func() {
		ch <- bestBid(url)
	}()

	select {
	case <-ctx.Done():
		return defaultBid
	case bid := <-ch:
		return bid
	}
}

func MainFindBind() {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	bid := findBind(ctx, "https://http.cat/418")

	fmt.Println(bid)

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	bid = findBind(ctx, "https://http.cat/404")

	fmt.Println(bid)
}

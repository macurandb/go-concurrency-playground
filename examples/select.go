package examples

import (
	"fmt"
	"time"
)

func MainSelect1() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 42
	}()
	select {
	case val := <-ch1:
		fmt.Printf("got %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("got %d from ch2\n", val)
	}
}

func MainSelect2() {
	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case val := <-out:
		fmt.Printf("got %f \n", val)
	case <-time.After(200 * time.Millisecond):
		fmt.Println("timeout")
	}
}

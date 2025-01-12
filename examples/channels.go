package examples

import (
	"fmt"
	"time"
)

func MainChannels1() {
	ch := make(chan int)

	go func() {
		ch <- 353
	}()

	val := <-ch
	fmt.Printf("got %d\n", val)

}

func MainChannels2() {
	ch := make(chan int)

	go func() {
		ch <- 353
	}()

	val, ok := <-ch
	fmt.Printf("got %d, %t\n", val, ok)

	close(ch)

	go func() {
		ch <- 353
	}()

	val, ok = <-ch
	fmt.Printf("got %d, %t\n", val, ok)

}

func MainChannels3() {

	ch := make(chan int)
	const count = 3

	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n")
		}
	}()

}

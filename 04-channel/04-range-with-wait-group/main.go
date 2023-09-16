package main

import (
	"fmt"
	"sync"
)

const total = 100

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(total)
	go publish(ch)
	go subscribe(ch, &wg)
	wg.Wait()
}

func publish(ch chan int) {
	for i := 0; i < total; i++ {
		ch <- i
	}
	close(ch)
}

func subscribe(ch chan int, wg *sync.WaitGroup) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
		wg.Done()
	}
}

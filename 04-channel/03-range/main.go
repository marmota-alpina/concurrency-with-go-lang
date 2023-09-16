package main

import "fmt"

func main() {
	ch := make(chan int)
	go publish(ch)
	subscribe(ch)
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func subscribe(ch chan int) {
	for x := range ch {
		fmt.Printf("Received: %d\n", x)
	}
}

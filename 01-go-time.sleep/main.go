package main

import (
	"fmt"
	"time"
)

func main() {
	go Echo("A")
	go Echo("B")
	time.Sleep(5 * time.Second)
}

func Echo(label string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Echo: %s sequence number %d \n", label, i)
		time.Sleep(1 * time.Second)
	}
}

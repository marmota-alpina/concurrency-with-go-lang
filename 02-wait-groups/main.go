package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//waitGroup2 := new(sync.WaitGroup)
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)
	go Echo("A", &waitGroup)
	go Echo("B", &waitGroup)
	waitGroup.Wait()
}

func Echo(label string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Echo: %s sequence number %d \n", label, i)
		time.Sleep(1 * time.Second)
	}
	wg.Done()
}

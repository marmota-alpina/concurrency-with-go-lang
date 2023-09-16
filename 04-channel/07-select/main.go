package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	Body string
	Id   int64
}

var counter int64

func main() {
	ch1 := make(chan Message)
	ch2 := make(chan Message)

	go func() {
		fmt.Println("Waiting for messages from RabbetMQ...")
		time.Sleep(time.Second * 1)
		for {
			atomic.AddInt64(&counter, 1)
			ch1 <- Message{Body: "Message from RabbetMQ!", Id: counter}
		}
	}()

	go func() {
		fmt.Println("Waiting for messages from Kafka...")
		time.Sleep(time.Second * 2)
		for {
			atomic.AddInt64(&counter, 1)
			ch2 <- Message{Body: "Message from Kafka!", Id: counter}
		}
	}()

	for {
		select {
		case msg := <-ch1:
			fmt.Printf("Received message with ID %d from RabbitMQ: %s \n", msg.Id, msg.Body)
		case msg := <-ch2:
			fmt.Printf("Received message with ID %d from Kafka: %s \n", msg.Id, msg.Body)
		}
	}
}

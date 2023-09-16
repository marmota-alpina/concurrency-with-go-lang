package main

import "fmt"

func main() {
	ch := make(chan string)

	go sendOnly("Hello, World!", ch)

	fmt.Println(receiveOnly(ch))
}

func sendOnly(name string, ch chan<- string) {
	ch <- name
}

func receiveOnly(ch <-chan string) string {
	return <-ch
}

package main

// Main thread
func main() {
	channel := make(chan string) // channel is empty
	go func() {
		channel <- "Hello World!" // channel has a message
	}()

	message := <-channel // channel is empty again

	println(message)
}

package main

func main() {
	ch := make(chan string, 2)

	ch <- "Hello, Earth!"
	ch <- "Hello, Moon!"

	// ch <- "Hello, World!" // This will block the execution of the program

	println(<-ch)
	println(<-ch)
}

package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "buffer"
	messages <- "channel"
	messages <- "finish"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

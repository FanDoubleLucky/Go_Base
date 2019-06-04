package main

import "fmt"

func main() {
	messages := make(chan string)
	go func() {
		messages <- "1"
		messages <- "1"

	}()
	go func() {
		fmt.Println(<-messages)
		fmt.Println("Finish11")
		fmt.Println(<-messages)
		fmt.Println("Finish12")
	}()

	go func() {
		fmt.Println(<-messages)
		fmt.Println("Finish21")
		fmt.Println(<-messages)
		fmt.Println("Finish22")
	}()

	fmt.Scanln()
	messages <- "F"

	fmt.Scanln()
}

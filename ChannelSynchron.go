package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 1)
	go worker(done)

	<-done
	fmt.Println("All Work Done")
}

func worker(done chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("Done")
	done <- true
}

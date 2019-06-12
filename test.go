package main

import (
	"fmt"
	"time"
)

func main() {
	ms := make(chan int, 5)

	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Println(<-ms)
		}
	}()

	var count int = 0
	for i := 1; i <= 6; i++ {
		count += 1
		ms <- i
	}

	time.Sleep(time.Second)
	fmt.Println(count)
}

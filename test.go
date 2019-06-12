package main

import (
	"fmt"
)

func showNumber(num int) {
	for i := 1; i <= 9000000; i++ {
		fmt.Println(num)
	}
}

func main() {
	go showNumber(10)
	//runtime.Gosched()
	for {
		fmt.Println("i")
	}
}

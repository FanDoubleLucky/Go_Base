package main

import "fmt"

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func ping(pings chan<- string, msg string) { //定义了pings只能收不能发，pings发会报编译错误
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { //pings只能发不能收，pongs只能收不能发
	msg := <-pings
	pongs <- msg
}

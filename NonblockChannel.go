package main

import "fmt"

//用default子句来实现非阻塞的收发
func main() {
	messages := make(chan string)
	signals := make(chan bool)
	select {
	case msg := <-messages: //如果messages收到信息
		fmt.Println("messages chan receives ", msg)
	default:
		fmt.Println("no message, free")
	}

	msg := "hi"
	select {
	case messages <- msg: //如果messages发送信息
		fmt.Println("messages chan sent ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages: //如果messages收到信息
		fmt.Println("received message", msg)
	case sig := <-signals: //如果signals收到信息
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}

package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "ping"
		messages <- "www"
		messages <- "baidu"
		fmt.Println("Over") //chan没加缓存情况下，如果下面main协程里接受少于发送，这个over将不会被打印，因为上面的发送没完成
	}()

	msg := <-messages
	msg = <-messages

	fmt.Println(msg)
}

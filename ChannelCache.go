package main

import "fmt"

func main() {
	messages := make(chan string, 3)

	messages <- "buffer"
	messages <- "channel"
	messages <- "finish" //如果上面不加缓存3，这句会报错，因为message没有缓存，而这句发消息没有对应的收消息

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

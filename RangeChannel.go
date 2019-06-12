package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) //如果没有关闭queue，下面的循环体在读取两遍后还会尝试继续读第三个

	for elem := range queue {
		fmt.Println(elem) //elem每遍历一次queue通道其实就等同于<-queue
	}
}

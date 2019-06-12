package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Millisecond * 500)

	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t) //每隔500ms打印一次
		}
	}()

	time.Sleep(time.Millisecond * 6600)
	ticker.Stop() //打点器也可以手动停止
	fmt.Println("Ticker stopped")
}

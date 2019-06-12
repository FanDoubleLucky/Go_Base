package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(time.Second * 2)
	<-timer1.C //timer计时结束之前一直阻塞
	fmt.Println("Timer1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 expired")
	}()
	stop2 := timer2.Stop() //如果timer在stop前就已经到期，返回false，定时器比单纯的time.Sleep好的就是可以在时间到之前手动取消
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
}

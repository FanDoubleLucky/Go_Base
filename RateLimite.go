package main

import (
	"fmt"
	"time"
)

//限速需要协程，通道和打点器组合使用来实现
func main() {
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	limiter := time.Tick(time.Millisecond * 2000)
	for req := range requests {
		<-limiter                               //此处为限流器的作用处,每次执行这一句的时候打点器才开始跑时间
		fmt.Println("request", req, time.Now()) //这是我们可以看到每隔200ms打印一次（处理一次request），这种是使用阻塞的方法进行限制速度
	}

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	burstyLimiter := make(chan time.Time, 3) //有时候我们不想影响整体的进程而进行限速，这时候就需要一个通道缓冲来实现
	for i := 3; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		time.NewTicker(time.Millisecond * 200)
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()
	time.Sleep(time.Second * 2)
	for req := range burstyRequests {
		<-burstyLimiter //此处为限流器的作用处
		fmt.Println("request", req, time.Now())
	}

	// 综合比较两种limiter，区别在于第一种是在主程里开打点器，每次时间到打了点才处理request，
	// 第二种通道限流器是将打点器放在了单独的协程中，每次打点器到了时间打了点就向限流通道中push一个消息，最终效果就是主程里每次要等有消息了才能执行request，每次消息之间的间隔就是副协程里的打点器时间间隔

}

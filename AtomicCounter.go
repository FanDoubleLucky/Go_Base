package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

//原子计数器，算是一种Golang管理协程状态的方式，应该就是CAS
func main() {
	var ops uint64 = 0
	var count uint64 = 0
	for i := 0; i < 5000; i++ {
		go func() {
			count += 1
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
	fmt.Println("count:", count)
}

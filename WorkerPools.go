package main

import (
	"fmt"
	"time"
)

func workerUnit(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { //main函数中没有执行close(jobs)之前不会这个range会一直执行等待，所以不会错过main里9个jobs<-j
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	for w := 1; w <= 3; w++ {
		go workerUnit(w, jobs, results)
	}
	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}
}

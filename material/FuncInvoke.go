package main

import "fmt"

func main(){
	myfunc(1,2,3,4,5,6)
}


func myfunc(args ...int){
	for _, arg := range args {
		fmt.Println(arg)
	}
}
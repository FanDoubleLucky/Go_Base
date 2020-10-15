package main

import "fmt"

func main(){
	var myArray [10]int = [10]int{1,2,3,4,5,6,7,8,9,10}
	var mySlice []int = myArray[:5]
	fmt.Println(myArray)
	fmt.Println(mySlice)
	SliceFunc()
}

func ArrayFunc(){}

func SliceFunc(){
	myslice1 := make([]int, 5)
	myslice2 := make([]int, 5, 10)
	myslice3 := []int{1,2,3,4,5}
	fmt.Println(myslice1, myslice2, myslice3)
	fmt.Println("myslice2 len: ", len(myslice2))
	fmt.Println("myslice2 cap: ", cap(myslice2))

	slice1 := []int{1,2,3,4,5}
	slice2 := []int{5,4,3}
	fmt.Println(copy(slice1,slice2)) // copy函数返回被copy的元素数量，当dst和src长度不一致时，按照短的那个从第0位开始进行copy
	fmt.Println(slice1,slice2)
}

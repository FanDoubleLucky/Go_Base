package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (this ByLength) Len() int {
	return len(this)
}

func (this ByLength) Swap(i int, j int) {
	temp := this[i]
	this[i] = this[j]
	this[j] = temp
}

func (s ByLength) Less(i int, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits)) //ByLength(fruits)是一步强制类型转换，类似与java的(int)1.2之类的操作
	fmt.Println(fruits)
}

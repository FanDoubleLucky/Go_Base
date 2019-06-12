package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("strs:", strs)

	ints := []int{23, 124, 35, 124, 64, 76, 12}
	sort.Ints(ints)
	fmt.Println(ints)

	s := sort.IntsAreSorted(ints) //检查一个序列是否已经是有序的，返回一个bool结果
	fmt.Println(s)
}

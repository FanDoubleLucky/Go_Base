package main

import (
	"fmt"
	"strings"
)

// 组合函数，一个某元素的切片，这种元素有独有对应的某种操作的函数，组合函数就是一个对整个切片进行的操作
// 如string有一个对应的函数，判断它是否包含非法字符，那么strings切片对应的组合函数就是判断是不是整个切片的元素都不包含非法字符

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

//如果切片中的字符串有一个满足了条件 f 则返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

//如果切片中的所有字符串都满足了条件 f 则返回true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

//返回所有切片中满足条件f的所有元素组成的新切片
func Filter(vs []string, f func(string) bool) []string {
	res := []string{}
	for _, v := range vs {
		if f(v) {
			res = append(res, v)
		}
	}
	return res
}

//将切片中所有元素按照f规则map成新的切片
func Map(vs []string, f func(string) string) []string {
	//第一种写法 myself
	res := []string{}
	for _, v := range vs {
		res = append(res, f(v))
	}

	//第二种写法 go by example
	res = make([]string, len(vs))
	for i, v := range vs {
		res[i] = f(v)
	}
	return res
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(All(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(Filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))
}

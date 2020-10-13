package main

import "fmt"

const Pi float64 = 3.1415926
const ZERO       = 0  // const变量类型可以省略
const (
	size int64 = 1024
	eof = -1
)

const (
	c0 = iota // 0 内部常量
	c1 = iota // 1 内部常量
	ExternalConst = iota // 2 外部常量 可被外部包访问
)

const (
	c2 = iota // 0 iota会重置回0
	c3 = iota // 1
)

func main(){
	VarValueSwap()
	fmt.Println(c0, c1, ExternalConst)
}

func VarValueSwap() {
	i := 1
	j := 2
	fmt.Printf("Before Swap -- i: %v j: %v\n", i, j)
	i,j = j,i // TODO 多重赋值，这使得在交换变量值时不需要第三个变量
	fmt.Printf("After Swap  -- i: %v j: %v\n", i, j)
}


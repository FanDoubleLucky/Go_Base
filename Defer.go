package main

import (
	"fmt"
	"os"
)

//defer类似于finally，做一些善后清理的操作，defer的操作是在所在函数结束的时候才运行
func main() {
	f := createFile("defer.txt")
	defer closeFile(f) //此处声明了defer但是不会执行，直到main结束的时候才执行，可以看到defer? writing都先于close打印
	fmt.Println("defer?")
	writeFile(f)
}

func createFile(path string) *os.File {
	fmt.Println("Creating File")
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	return f
}

func closeFile(f *os.File) {
	fmt.Println("Closing File")
	f.Close()
}

func writeFile(f *os.File) {
	fmt.Println("Writing File")
	fmt.Fprintln(f, "First message")  //输入第一行内容
	fmt.Fprintln(f, "Second message") //输入第二行内容
}

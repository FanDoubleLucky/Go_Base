package main

import "os"

//程序出现来出乎意料的错误，我们不知道该如何处理或者根本不想处理的时候就打印panic,panic不同于error，error一定程度上是可以预测的
func main() {
	panic("a problem")

	_, error := os.Create("/tmp/file")

	if error != nil {
		panic(error)
	}
}

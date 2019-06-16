package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("defer.txt")
	//dat是读出来的文件内容，dat类型是[]uint8
	check(err)
	fmt.Println(len(dat))
	fmt.Println(string(dat))
	//这种使用ioutil read是将整个文件全部读取到内存中

	//os.Open得到句柄 然后利用句柄可以进行更加细化的读取
	f, err := os.Open("defer.txt") //f不是文件内容，只是文件句柄,f类型是*os.File
	check(err)

	defer f.Close()

	//设置好read的容器的大小，来读取定量的数据
	b1 := make([]byte, 10)
	n1, err := f.Read(b1) //read 返回读取到的内容长度
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))

	o2, err := f.Seek(6, 0) //从文件开头0处数6个offset开始读
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

	o3, err := f.Seek(16, 0)
	check(err)
	b3 := make([]byte, 5)
	n3, err := io.ReadAtLeast(f, b3, 3) //从f句柄中读内容到b3中，至少要读3bytes，所以最终读到多少还是要看b3设置了多大
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
}

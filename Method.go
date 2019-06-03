package main

import "fmt"

type person struct {
	name    string
	project string
	age     int
}

//通过这个例子可以看出直接将结构体赋值给别的变量是copy了一份新的，用指针的话就是指向同一个
func main() {
	s := person{project: "s", age: 20}
	fyz := person{"fyz", "EE", 20}
	lsb := fyz
	lsb.name = "lsb"
	fmt.Println(lsb.name) //lsb
	fmt.Println(fyz.name) //fyz
	wyf := &fyz
	wyf.name = "wyf"

	fmt.Print(fyz.name) //wyf
	wyf.learn()
}

func (this *person) learn() {
	fmt.Println(this.name, "learning")
}

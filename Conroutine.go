package main

import "fmt"

func f1(input string) {
	for i := 0; i < 30; i++ {
		fmt.Println(input, ":", i)
	}
}

func main() {
	f1("R1")
	go f1("R2")
	go f1("R3")

	go func(name string) {
		for i := 0; i < 10; i++ {
			fmt.Println(name, ":", i)
		}
	}("Inner func")

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}

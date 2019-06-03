package main

import (
	"errors"
	"fmt"
)

func divide(arg0 int, arg1 int) (int, error) {
	if arg1 == 0 {
		return -1, errors.New("arg1 can't be zero")
	}

	return arg0 / arg1, nil
}

func main() {
	result, err := divide(4, 0)
	fmt.Println(result, err)
	result, err = f2(0)
	if errT, ok := err.(*argError); ok {
		fmt.Println(errT.arg)
		fmt.Println(errT.prob)
	}

}

type argError struct {
	arg  int
	prob string
}

func (this *argError) Error() string {
	return fmt.Sprintf("%d - %s", this.arg, this.prob)
}

func f2(arg int) (int, *argError) {
	if arg == 0 {
		return -1, &argError{arg, " can't work"}
	}
	return arg, nil
}

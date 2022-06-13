package main

import (
	"errors"
	"fmt"
)

//defer  与java中的finally类似
func main() {
	c, err := add(-1, -2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(c)
	}
	defer myPrint()
	defer func() {
		fmt.Println("aaa")
	}()

}

func myPrint() {
	fmt.Println("兜底执行")
}
func add(a, b int) (c int, err error) {
	if a < 0 || b < 0 {
		err = errors.New("param can below zero")
		return
	}
	c = a + b
	err = nil
	return
}

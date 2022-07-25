package chapter5

import (
	"fmt"
	"math"
	"strings"
)

func LeFunc1() {
	fmt.Println(G(0.3, 0.4))
}
func G(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func Func_Sign() {
	fmt.Printf("%T\n", Add)
	fmt.Printf("%T\n", Sub)
	fmt.Printf("%T\n", First)
	fmt.Printf("%T\n", Zero)
}

func Add2(rune2 rune) rune {
	return rune2 + 12
}
func Func_Learn() {
	f := Add
	fmt.Println(f(1, 2))
	f = Sub
	s := strings.Map(Add2, "Admix")
	fmt.Printf(s)
}

// 函数的类型称为函数签名 当两个函数用用相同的形参列表以及返回类型时 则认位两个函数的签名一致
func Add(x, y int) int {
	return x + y
}
func Sub(x, y int) int {
	return x - y
}
func First(x, y int) int {
	return x
}
func Zero(x, y int) int {
	return 0
}

func LearnGetFunc() {
	f := GetFunc()
	g := f(1, 2)
	fmt.Println(g)
}

func GetFunc() func(int, int) int {
	return Add
}

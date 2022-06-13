package design_pattern

import (
	"fmt"
	"time"
)

//设置别名

type mutliPlyFunc func(int, int) int

func multiply(a, b int) int {
	return a * b
}
func multiply2(a, b int) int {
	return a << b
}

//返回值是一个函数
func exeTime(f mutliPlyFunc) mutliPlyFunc {
	return func(i int, i2 int) int {
		start := time.Now()
		r := f(i, i2)
		end := time.Since(start)
		fmt.Println("执行耗时：--", end)
		return r
	}
}
func test() {
	b := multiply(1, 2)
	fmt.Println(b)

	//返回一个装饰器
	plyFunc := exeTime(multiply)
	c := plyFunc(1, 2)
	fmt.Println(c)

	m := exeTime(multiply2)
	fmt.Println(m(1, 2))
}

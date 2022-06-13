package learn

import "fmt"

/*
匿名函数 使用场景
1：保证局部变量的安全性
2：将匿名函数作为函数参数
*/
func learn2() {
	//匿名函数  java中的 lamda
	add := func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
}

/**
通过匿名函数定义不同实现
*/
func learn3() {

	//匿名函数
	add1 := func(a, b int) int {
		return a + b
	}
	add(1, 2, add1)

	add2 := func(a, b int) int {
		return a*2 + b
	}
	add(1, 2, add2)
}

func add(a, b int, call func(a, b int) int) {
	fmt.Println(call(a, b))
}

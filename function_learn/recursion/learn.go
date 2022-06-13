package recursion

import (
	"fmt"
	"time"
)

type myFunc func(int) int

//通过序号确定斐波那契数列值
func fibo(n int) int {

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	return fibo(n-1) + fibo(n-2)
}

//装饰器模式
func exeTime(myFunc2 myFunc) myFunc {

	return func(i int) int {
		s := time.Now()
		value := myFunc2(i)
		e := time.Since(s)
		fmt.Println("执行耗时： ", e)
		return value
	}
}
func test() {

	m := exeTime(fibo)
	v := m(50)     //执行耗时：  54.9858205s
	fmt.Println(v) //7778742049
	//贼垃圾
}

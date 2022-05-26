package main

import "fmt"

const size int = 20

//go中变量的作用域 与变量声明的位置有关 函数内则为局部变量  函数外则为全局变量
var age int = 20

func getName() (userName, nickName string) {
	return "Young", "Little"
}

//以大写字母开头的常量在包外可见 小写字母开头的常量只能在包内访问
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	numberOfDays
)

func main() {
	var v1 int //声明一个变量
	fmt.Println("v1 =", v1)
	var (
		v2 int
		v3 string
	)
	fmt.Println("v2 =", v2)
	fmt.Println("v3 =", v3)
	var v11 int
	fmt.Println("v11 =", v11)
	var v22 string
	fmt.Println("v22 =", v22)
	var v33 bool
	fmt.Println("v33 =", v33)
	var v44 [10]int
	fmt.Println("v44 =", v44)
	var v55 struct {
		f float64
	}
	fmt.Println("v55 =", v55)
	var v66 *int
	fmt.Println("v66 =", v66)
	var v77 map[string]int
	fmt.Println("v77 =", v77)
	var v88 func(a int) int
	fmt.Println("v88 =", v88)

	var c11 int = 10 //自己声明类型
	fmt.Println("c11 =", c11)
	var c22 = 10 // 编译器自己推导
	fmt.Println("c22 =", c22)
	c33 := 10 // 省略var 这种变量初始化方法 变量应该是未声明过的
	fmt.Println("c33 =", c33)

	_, nickName := getName()
	fmt.Println("nickName =", nickName)

	//多重赋值
	var i int = 10
	var a int = 20
	i, a = a, i
	fmt.Println("i=", i)
	fmt.Println("a = ", a)

	fmt.Println("const size is:", size)
}

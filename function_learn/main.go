package main

import (
	"fmt"
	"function_learn/mymath"
)

func main() {
	a := mymath.Add(1, 2) //调用其他包函数时首字母必须大写 属于值传递
	fmt.Println(a)
	//内置函数可以随时引用
	//len 返回入参的长度  cap返回容量
	str := "golang"
	fmt.Println(len(str))

	arr := [5]string{"a", "b", "c", "d", "e"}
	s := arr[:1]
	fmt.Println(len(s), cap(s))

	dict := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println(len(dict))
	//new make
	p1 := new(int)
	fmt.Println(p1) //指针

	type student struct {
		id   int
		name string
		age  int
	}
	p4 := new(student)
	fmt.Println(p4) //返回指针

	intS := make([]int, 2)
	fmt.Println(intS)

	m1 := make(map[string]int, 2)
	m1["a"] = 100
	fmt.Println(m1)

	var a1 int = 1
	var b int = 2
	c := add2(&b, &a1)
	fmt.Println(a1, b)
	fmt.Println(c)

	//可变参数
	myFunc(1, 2, 3, 4)

	slice := []int{1, 2, 3, 4, 5}
	myFunc(slice...)
}

//引用传递
func add2(a, b *int) int {
	*a *= 2
	*b *= 2
	return *a + *b
}

func myFunc(nums ...int) {
	for _, i2 := range nums {
		fmt.Println(i2)
	}

}

package details

import "fmt"

// go 中 数组的长度初始化后是不可变的  需要一个可变集合 接下来学习下

func sliceLearn() {
	var slice []string = []string{"a", "b"} //初始化时没有指定长度
	fmt.Println(slice)
	//定义一个数组
	months := [...]string{"January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	q2 := months[3:6]
	fmt.Println(q2)
	summer := months[5:8] //左闭右开原则
	fmt.Println(summer)

	year := months[:]
	fmt.Println(year)

	first := months[:6]
	second := months[6:]
	fmt.Println(first)
	fmt.Println(second)
	//上面方式为通过数组生产切片

	//同样也可以基于切片 创造切片

	//直接创建
	mySlice := make([]int, 5) //第二个参数代表初始长度
	fmt.Println(mySlice)

	mySlice2 := []int{1, 2, 3}
	fmt.Println(mySlice2)

	//遍历切片
	for i2 := range mySlice2 {
		fmt.Println(i2)
	}
	//动态添加元素
	var oldSlice = make([]int, 5, 10)
	fmt.Println("len is ", len(oldSlice), " cap is ", cap(oldSlice))
	fmt.Println(oldSlice)

	newSlice := append(oldSlice, 6, 6, 6)
	fmt.Println(newSlice)

	//元素复制  copy 函数有坑  慎用
	//大坑 切片出来的新切片 会和旧切片 指向共同的引用 属于浅拷贝
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := slice1[1:3]
	fmt.Println(slice2)
	slice2[1] = 6
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)
	//slice1: [1 2 6 4 5]
	//slice2: [2 6]
}

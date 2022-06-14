package main

import (
	"fmt"
	"math/rand"
)

func selectLearn() {
	//初始化一个chan数组 没加缓冲
	chs := [3]chan int{make(chan int, 1), make(chan int, 1), make(chan int, 1)}
	index := rand.Intn(3)
	fmt.Printf("随即索引/数值 %d \n", index)
	chs[index] <- index

	select {
	case num := <-chs[0]:
		fmt.Println("第一个分支被选中 num ", num)

	case num2 := <-chs[1]:
		fmt.Println("第2个分支被选中 num2 ", num2)
	case num3 := <-chs[2]:
		fmt.Println("第3个分支被选中 num2 ", num3)
	default: //没有default 会阻塞
		fmt.Println("没有分支被选中")
	}
}

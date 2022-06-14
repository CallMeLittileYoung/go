package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
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

func select2() {
	s := time.Now()
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)

	chans := make([]chan int, cpus)

	for i := 0; i < cpus; i++ {
		ch := make(chan int)
		chans[i] = ch
		go sum4(i, ch) //开启协程
	}
	sum := 0
	for _, ch := range chans {
		sum += <-ch
	}
	//total sum is 399999960000000  cost: 11.1755ms
	//--- PASS: Test_select_learn (0.01s)
	fmt.Println("total sum is", sum, " cost:", time.Since(s))
}

//新增

func sum4(seq int, ch chan int) {
	defer close(ch) //保证channel正常关闭
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}
	fmt.Println("子协程 ", seq, " 运算结果", sum)
	ch <- sum
}

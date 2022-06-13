package main

import (
	"fmt"
	"time"
)

func chanLearn() {

	s := time.Now()
	chs := make([]chan int, 100)

	for i := 0; i < 100; i++ {
		chs[i] = make(chan int)
		go Add3(i, 1, chs[i])
	}

	for _, ch := range chs {
		//fmt.Println(ch)
		<-ch
		//我们再通过 <-ch 语句从通道切片 chs 中依次接收数据（不对结果做任何处理，相当于写入通道的数据只是个标识而已，表示这个通道所属的协程逻辑执行完毕）
	}
	fmt.Println("cost", time.Since(s))
}

func Add3(a, b int, ch chan int) {

	fmt.Println("a ", a, " + b ", b, " sum is ", a+b)
	ch <- 1
}

func channelLearn2() {
	s := time.Now()
	ch := make(chan int, 20) //增加20个缓冲区,可以提高性能
	go test2(ch)

	for i2 := range ch {
		fmt.Println(i2)
	}
	end := time.Now()
	consume := end.Sub(s).Seconds()
	fmt.Println("程序执行耗时(s)：", consume)
}
func test2(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

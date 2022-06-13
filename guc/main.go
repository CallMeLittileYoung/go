package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 100; i++ {
		go add(i, 1) //开启协程
	}
	time.Sleep(1000 * time.Millisecond) //time sleep
}

func add(a, b int) {
	fmt.Println("a+b = ", a+b)
}

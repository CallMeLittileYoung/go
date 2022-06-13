package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var counter int = 0

func lockLearn() {
	s := time.Now()
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		go Add2(i, 1, &lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
	fmt.Println("cost :", time.Since(s))
}
func Add2(a, b int, lock *sync.Mutex) {
	lock.Lock()
	counter++
	fmt.Println("a+b is", a+b, " counter is ", counter)
	lock.Unlock()
}

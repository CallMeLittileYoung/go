package chapter8

import (
	"fmt"
	"time"
)

func spinner(delay time.Duration) {
	for {
		for _, r := range `abcde` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func test() {
	go spinner(10 * time.Microsecond)
	const x = 45
	fib := fib(x)
	fmt.Printf("\rx is %d value is %d", x, fib)
}

package main

import "fmt"

func main() {
	var i Integer = 1
	var m Math = i //将Integer 赋值给Math
	v := m.Add(2)
	fmt.Println(v)

	fmt.Println(m.Multiply(4))
}

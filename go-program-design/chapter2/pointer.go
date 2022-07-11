package chapter2

import "fmt"

// Pointer 指针
func Pointer() {
	p := 10
	changeValue(&p)
	fmt.Println(p)

	changeValue2(p)
	fmt.Println(p)
}

//传递一个指针
func changeValue(p *int) {
	*p++
}

//传递一个指针
func changeValue2(p int) {
	p++
}

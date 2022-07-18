package chapter4

import "fmt"

// 调用函数时候 每次传递的都是 一个副本  所以传递 比如大数组时 变会存在效率问题，并且函数内部的修改都只影响副本
func Arr() {
	//长度为4的数组 值为类型的默认值
	var a [4]int
	fmt.Println(a)
	var b [4]int = [4]int{1, 2, 3, 4}
	fmt.Println(b)
}

func Zero(b *[32]byte) {
	for i, _ := range b {
		b[i] = 0
	}
}

func Zero2(b *[32]byte) {
	*b = [32]byte{}
}

package chapter3

import (
	"fmt"
)

func Op() {
	//0000 0010  |  0010 0000 = 0010 0010
	var x uint8 = 1<<1 | 1<<5
	//0000 0010 |  0000 0100 =0000 0110
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b \n", x)
	fmt.Printf("%08b \n", y)

	//交集 &    0000 0010
	fmt.Printf("a ∩ b is %08b \n", x&y)
	//并集 |    0010 0110
	fmt.Printf("a ∪ b is %08b \n", x|y)
	//对称差 0010 0100
	fmt.Printf("a ^ b is %08b \n", x^y)
	//差集  0010 0000 元素存在于a中 且不存在于b
	fmt.Printf("a -b is  %08b \n", x&^y)

}

func SVG() {

}

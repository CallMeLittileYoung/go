package details

import "fmt"

/**
  go整型学习
*/
func intLearn() {
	var a int8 = 127 //int8 类型 1字节
	fmt.Println("a=", a)
	var b uint8 = 255 //无符号8位整形
	fmt.Println("b=", b)
	var c int16 = 32767 //带符号16位整形
	fmt.Println("c=", c)
	var d uint16 = 65535 //无符号16位整形
	fmt.Println("d=", d)
	var e int32 = 2147483647 //带符号32位整形 4字节
	fmt.Println("e=", e)
	var f uint32 = 2147483647 //无符号32位整形 4字节
	fmt.Println("=", f)
	var g int = 19 //与平台有关 32位或者64位 uint同理
	fmt.Println("g=", g)

	//var h int8 = 10; go无法做不同类型转换 这种和java有区别
	//var i int16 = h;
}

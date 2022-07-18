package chapter4

import "fmt"

func Sl() {
	months := []string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months)
	fmt.Println("A ", months[0])
	//
	q2 := months[4:7] //左闭右开
	fmt.Println(q2)   //[April May June]
	months[4] = "Aprils"
	fmt.Println(q2) //[Aprils May June]
	//新切片与旧切片共享底层数组 所以对原始切片的修改 会影响到新的切片
	//q22 := q2[:20]  如果超过了原始切片的容量 会触发程序崩溃
	//fmt.Println(q22)
	strings := q2[:5] //[Aprils May June July August] 直接从原始切片中扩容了  切片的切片要使用时候要谨慎
	fmt.Println(strings)

	//创建一个数组的slice 等于为这个数组起了个别名
}

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

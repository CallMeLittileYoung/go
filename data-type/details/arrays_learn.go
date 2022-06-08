package details

import "fmt"

func arrayLearn() {
	var a [8]byte //长度为8的byte数组
	fmt.Println(a)

	var b [3][3]int //二维数组
	fmt.Println(b)

	var c [3][3][3]float64 //三维数组
	fmt.Println(c)

	var d = [3]int{1, 2, 3} //声明时初始化
	fmt.Println(d)

	var e = new([3]string)
	fmt.Println(e)

	f := [5]int{1, 2, 3, 4, 5}
	fmt.Println(f)
	// 数组声明方式 [cap]type{element}
	g := [5]int{1, 2, 3}
	fmt.Println(g) //空位由默认值填充

	//初始化指定下标元素
	h := [5]int{0: 1, 3: 5}
	fmt.Println(h) //[1 0 0 5 0]
	//数组在声明长度后不可更改

	i := [5]int{1, 2, 3, 4, 5}
	a1, a2 := i[0], i[len(i)-1]
	fmt.Println(a1, a2)
	i[4] = 100 //元素赋值
	fmt.Println(i[4])
	//遍历数组
	for index := 0; index < len(i); index++ {
		fmt.Println("Ele", index, "of arr is ", i[index])
	}
	//java中的foreach
	for i2, i3 := range i {
		fmt.Println("Ele", i2, "of arr is ", i3)
	}
	//go中的语法糖
	for _, i3 := range i {
		fmt.Println("ele value is", i3)
	}
	//九九乘法表
	var multi [9][9]string

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			n1 := i + 1
			n2 := j + 1
			if n1 < n2 {
				continue
			}
			multi[i][j] = fmt.Sprintf("%dx%d=%d", n1, n2, n1*n2)
		}
	}
	for _, i3 := range multi {
		for _, i4 := range i3 {
			fmt.Printf("%-8s", i4)
		}
		fmt.Println()
	}

}

package chapter2

import "fmt"

func NewLearn() {
	//声名一个变量 注意 new返回的是地址 ,可以多多用new来初始化变量
	i := new(int)
	fmt.Println(*i)
}

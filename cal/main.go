package main

import (
	"cal/simplemath"
	"fmt"
	"os"
	"strconv"
)

var Usage = func() {
	fmt.Println("aaa")
	fmt.Println("bbb")
}

func main() {
	/*
		获取命令行参数
	*/
	args := os.Args
	if args == nil || len(args) < 3 {
		Usage()
		return
	}

	switch args[1] {

	case "add":
		if len(args) != 4 {
			fmt.Println("wrong args")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("wrong args")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("result", ret)

	case "sqrt":
		if len(args) != 3 {
			fmt.Println("wrong args")
			return
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("wrong args")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("result", ret)
	}
}

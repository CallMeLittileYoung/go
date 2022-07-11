package chapter1

import (
	"bufio"
	"fmt"
	"os"
)

func FindRepeatLine() {
	//初始化一个map
	m := make(map[string]int)
	//NewScanner returns a new Scanner to read from r   os.Stdin 标准输入
	for s, i := range m {
		if i > 1 {
			fmt.Printf("%d\t%s\n", i, s)
			//fmt.Println("%d\t%s", i, s) why warning?
		}
	}
}

//打开一个文件并读取
func FindRepeatLine2() {
	//初始化一个map
	m := make(map[string]int) // make返回的是指针
	file, err := os.Open("C:\\Users\\Administrator\\Desktop\\go-test.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
		return
	}
	CountLines(file, m)

	for s, i := range m {
		fmt.Printf("%d\t%s\n", i, s)
	}
}

func CountLines(f *os.File, count map[string]int) {
	// os.File
	input := bufio.NewScanner(f)
	//
	for input.Scan() {
		count[input.Text()]++
	}
}

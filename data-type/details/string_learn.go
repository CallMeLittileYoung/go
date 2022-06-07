package details

import fmt "fmt"

func stringLearn() {
	var str string
	str = "Hello World"
	str2 := "叫我小年轻"
	fmt.Printf("The Length of %s is %d \n", str, len(str))
	fmt.Printf("str2 is %s \n", str2)
	//多行字符串
	results := `Search results for "Golang":
				- Go
				- Golang
                Golang Programming
                `
	fmt.Printf("%s", results)

	str4 := "叫我小年轻"
	n := len(str4)
	fmt.Println(n)
	for i, ch := range str4 {
		fmt.Println(i, string(ch))
	}
	str5 := str4[:3] //叫
	fmt.Println(str5)
}

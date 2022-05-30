package details

import "fmt"

func floatLearn() {
	var floatValue float32
	floatValue = 10
	fmt.Println("value is ", floatValue)

	floatValue2 := 10.0 //自动推导 会推导为float64
	fmt.Println("value2 is ", floatValue2)

	//floatValue = floatValue2  golang中不允许大类型接小类型
	floatValue = float32(floatValue2) //必须强转

	floatValue3 := 0.1
	floatValue4 := 0.7
	floatValue5 := floatValue3 + floatValue4
	fmt.Println("value5 is ", floatValue5)
	//value5 is  0.7999999999999999
}

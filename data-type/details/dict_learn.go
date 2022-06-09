package details

import "fmt"

func dictLearn() {
	var testMap = map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
	}
	k := "1"
	i, ok := testMap[k] //查找元素
	if ok {
		fmt.Println(i)
	}
	//go map是无序的
	fmt.Println(testMap)
	fmt.Println(testMap)

	testMap["4"] = 4
	fmt.Println(testMap)

	//通过make函数创建
	m := make(map[string]int)
	m["1"] = 1
	m["2"] = 2
	fmt.Println(m)
	//删除元素
	delete(m, "1")
	fmt.Println(m)
	//遍历元素
	for s, i2 := range m {
		fmt.Println("key is ", s, "values is ", i2)
	}
	//go 字典的排序 需要先依靠对 key进行排序
}

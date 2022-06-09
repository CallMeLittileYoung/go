package flow_control

import "fmt"

func learn() {

	var scores int = 100
	switch {
	case scores >= 90:
		fmt.Println("Grades is A")
	case scores >= 80 && scores < 80:
		fmt.Println("Grades is B")
	case scores >= 70 && scores < 80:
		fmt.Println("Grade: C")
	case scores >= 60 && scores < 70:
		fmt.Println("Grade: D")
	default:
		fmt.Println("Grade: F")
	}

	score := 80
	switch score { //只有分支值判等时才
	case 90, 100:
		fmt.Println("Grades: A")
		fallthrough //傻子才会在代码里面这么写吧 会继续执行下个分支
	case 80:
		fmt.Println("Grades: B")
		fallthrough
	case 70:
		fmt.Println("Grades: C")
	case 60:
	case 65:
		fmt.Println("Grades: D")
	default:
		fmt.Println("Grades: F")
	}
}

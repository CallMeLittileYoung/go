package chapter5

import "fmt"

type Rocket struct {
}

func (receiver *Rocket) Launch() {
	fmt.Println("起飞了")
}

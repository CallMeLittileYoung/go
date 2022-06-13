package main

import (
	"fmt"
	. "object-oriented/entity" //. 无需使用entity
)

/**
类属性与成员方法的可见性 以及组合 和多态的实现
*/
func main() {

	ani := NewAnimal("中华田园犬")
	fmt.Println(ani)
	pet := NewPet("宠物狗")
	dog := NewDog(ani, pet)

	fmt.Println(dog.GetName())
	fmt.Println(dog.FavorFood())
}

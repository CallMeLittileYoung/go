package main

import (
	"fmt"
	"object-oriented/entity"
)

func main() {
	animal := entity.NewAnimal("中华田园犬")
	dog := entity.NewDog(*animal)
	fmt.Println(dog.GetName())
	fmt.Println(dog.Call()) //汪汪汪
	fmt.Println(dog.FavorFood())
}

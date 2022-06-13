package test

import (
	"fmt"
	entity "object-oriented/entity"
	"testing"
)

func Test_all(t *testing.T) {

	student := entity.NewStudent(1, "Young", 12, true, 100)
	fmt.Println(student)
	id := student.Id()
	fmt.Println("id is ", id)
	name := student.Name()
	fmt.Println("name is", name)
	student.SetName("Little Young")
	fmt.Println(student.Name())

	animal := entity.NewAnimal("中华田园犬")
	fmt.Println(animal)

	dog := entity.NewDog(*animal)
	fmt.Println(dog)

}

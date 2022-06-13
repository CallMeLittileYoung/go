package entity

//通过组合实现继承

type Dog struct {
	Animal
}

func NewDog(animal Animal) *Dog {
	return &Dog{Animal: animal}
}

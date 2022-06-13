package entity

//通过组合实现继承

type Dog struct {
	animal *Animal
	pet    Pet
}

func NewDog(animal *Animal, pet Pet) *Dog {
	return &Dog{animal: animal, pet: pet}
}

// Call 重写方法实现覆盖
func (a Dog) Call() string {
	return a.animal.Call() + "汪汪汪"
}

func (a Dog) FavorFood() string {
	return a.animal.FavorFood() + "骨头"
}

func (a Dog) GetName() string {
	return a.pet.GetName()
}

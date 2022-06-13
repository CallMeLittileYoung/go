package entity

//通过组合实现继承

type Dog struct {
	Animal
}

func NewDog(animal Animal) *Dog {
	return &Dog{Animal: animal}
}

// Call 重写方法实现覆盖
func (a Dog) Call() string {
	return "汪汪汪"
}

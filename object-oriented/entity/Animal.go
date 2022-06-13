package entity

type Animal struct {
	Name string
}

func NewAnimal(name string) *Animal {
	return &Animal{Name: name}
}

func (a Animal) Call() string {
	return "动物的叫声"
}

func (a Animal) FavorFood() string {
	return "动物爱吃食物"
}
func (a Animal) GetName() string {
	return a.Name
}

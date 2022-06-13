package main

// Integer 定义一个integer类型
type Integer int

func (a Integer) Add(b Integer) Integer {
	return a + b
}
func (a Integer) Multiply(b Integer) Integer {
	return a * b
}

type (
	Math interface {
		Add(i Integer) Integer
		Multiply(integer Integer) Integer
	}
)

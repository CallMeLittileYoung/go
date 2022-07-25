package chapter5

type IntList struct {
	value int
	tail  *IntList
}

func (l *IntList) Sum() int {
	if nil == l {
		return 0
	}
	return l.value + l.tail.Sum()
}

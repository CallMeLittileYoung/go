package chapter5

import (
	"fmt"
	"testing"
)

func Test_IntList(t *testing.T) {
	i := &IntList{
		value: 1,
		tail: &IntList{
			value: 2,
			tail: &IntList{
				value: 3,
				tail:  nil,
			},
		},
	}
	fmt.Println(i.Sum())
}

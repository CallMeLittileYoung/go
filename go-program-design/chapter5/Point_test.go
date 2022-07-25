package chapter5

import (
	"fmt"
	"testing"
)

func Test_Path(t *testing.T) {
	path := Path{
		{x: 1, y: 1},
		{x: 5, y: 1},
		{x: 5, y: 4},
		{x: 1, y: 1},
	}
	distance := path.Distance()
	fmt.Println(distance)
}

func Test_Point(t *testing.T) {
	fmt.Printf("%T\n", (*Point).ScaleBy) //func(*chapter5.Point, float64)
}

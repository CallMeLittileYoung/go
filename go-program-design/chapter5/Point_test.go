package chapter5

import (
	"fmt"
	"image/color"
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

func Test_ColoredPoint(t *testing.T) {

	point := &ColoredPoint{
		Point: Point{1, 2},
		Color: color.RGBA{R: 3, G: 1, B: 2, A: 2},
	}
	p := Point{
		x: 2,
		y: 3,
	}
	//可以理解为编译器 根据组合类型 生成了新的 Distance方法
	distance := point.Distance(p)
	fmt.Println(distance)
}

func Test_Method(t *testing.T) {
	point1 := Point{
		x: 1,
		y: 2,
	}
	point2 := Point{
		x: 2,
		y: 3,
	}
	f := Point.Distance
	f2 := f(point2, point1)
	fmt.Println(f2)
}

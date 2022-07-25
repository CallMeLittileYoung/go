package chapter5

import (
	"image/color"
	"math"
)

type Point struct {
	x, y float64
}

// ColoredPoint point 方法全部被纳入ColoredPoint中
type ColoredPoint struct {
	Point //匿名成员
	Color color.RGBA
}

// Distance p *Point 被称为接收器  可以看作时Distance 属于 对象p
func (p *Point) Distance(q Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

func (p *Point) ScaleBy(factor float64) {
	p.x *= factor
	p.y *= factor
}

// Distance 包级别函数
func Distance(q, p Point) float64 {
	return math.Hypot(q.x-p.x, q.y-p.y)
}

type Path []Point

func (p Path) Distance() float64 {
	sum := 0.0
	for i, point := range p {
		if i > 0 {
			// := &point  编译器会隐式的转为指针
			sum += point.Distance(p[i-1])
		}
	}
	return sum
}

package test

import (
	"go-program-design/chapter4"
	"testing"
)

func Test_Struct(t *testing.T) {

	var wheel chapter4.Wheel
	wheel.Spoke = 8
	wheel.Circle.Radius = 4
	wheel.Circle.Center.Y = 1
	wheel.Circle.Center.X = 2
	// 上述通过组合方式的struct 访问起来很麻烦 go提供了匿名成员
	var wheel2 chapter4.Circle2
	wheel2.X = 1
	wheel2.Y = 2
	//看起来简洁了很多

}

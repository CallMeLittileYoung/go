package chapter4

type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spoke  int
}

type Point2 struct {
	X, Y int
}

type Circle2 struct {
	Point2
	Radius int
}

type Wheel2 struct {
	Circle2
	Spoke int
}

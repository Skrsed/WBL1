package task

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x float64, y float64) Point {
	return Point{x, y}
}

func (p1 Point) distanceTo(p2 Point) float64 {
	dxSqr := math.Pow(p1.x-p2.x, 2)
	dySqr := math.Pow(p1.y-p2.y, 2)

	return math.Sqrt(dxSqr + dySqr)
}

// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами
// x,y и конструктором.
func FindDistance() {
	point1 := NewPoint(6, 7)
	point2 := NewPoint(10, 10)

	fmt.Println(point1.distanceTo(point2))
}

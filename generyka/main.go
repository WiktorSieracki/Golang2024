package main

import (
	"fmt"
	"math"
)

type Point[T types] struct {
	X, Y T
}

type types interface {
	int | int16 | int32 | int64 | float32 | float64 | uint | uint16 | uint32 | uint64
}

// func distance[T types](x, y T) T {
// 	return T(math.Sqrt(float64(x*x + y*y)))
// }

func (p Point[T]) distance() T {
	return T(math.Sqrt(float64(p.X*p.X + p.Y*p.Y)))
}

func main() {
	p := Point[float64]{1, 2}
	fmt.Println(p.distance())

}
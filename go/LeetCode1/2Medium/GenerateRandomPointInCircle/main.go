package main

import (
	"crypto/rand"
	"math"
)

type Solution struct {
	radius   float64
	x_center float64
	y_center float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {

}

func (this *Solution) RandPoint() []float64 {
	//
	d := this.radius * math.Sqrt(rand.Int())
	tht := rand.Int() * 2 * math.Pi
	return []float64{math.Cos(tht)*d + this.x_center, math.Sin(tht)*d + this.y_center}

}

func main() {
	obj := Constructor(2.0, 1.1, 0.7)
	param_1 := obj.RandPoint()
}

package main

import (
	"fmt"
	"math"
)

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	times := minutesToTest / minutesToDie
	times = times + 1
	result := math.Log(float64(buckets)) / math.Log(float64(times))
	return int(math.Ceil(result))
}

func main() {
	b := 1000
	mD := 15
	mT := 60
	fmt.Printf("%v\n", poorPigs(b, mD, mT))
}

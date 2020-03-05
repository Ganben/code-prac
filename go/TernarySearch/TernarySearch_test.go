package TernarySearch

import (
	"testing"
	"math/rand"
	"time"
)

func SortArray(array []int) {
	for itemIndex, itemValue := range array {
		for itemIndex != 0 && array[itemIndex-1] > itemValue {
			array[itemIndex] = array[itemIndex-1]
			itemIndex -= 1
		}
		array[itemIndex] = itemValue
	}
}

func TestTernarySearch(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, random.Intn(100-10)+10)
	for i := range array {
		array[i] = random.Intn(1000)
	}

	SortArray(array)
	for i, value := range array {
		result := TernarySearch(array, value)
		if result != i {
			t.Fail()
		}
	}
}
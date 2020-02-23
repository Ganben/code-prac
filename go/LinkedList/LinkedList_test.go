package LinkedList

import (
	"math/rand"
	"testing"
	"time"
)

func TestingGetItems(t *testing.T) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	array := make([]int, random.Intn(100-10)+10)
	linkedlist := &LinkedList{}

	for i:= range array {
		array[i] = random.Intn(100)
		linkedlist.InsertLast(array[i])
	}

	array2 := linkedlist.GetItems()
	for i := range array {
		if array[i] != array2[i] {
			t.Fail()
		}

		result := linkedlist.SearchValue(array[i])
		if result == false {
			t.Fail()
		}
	}

}
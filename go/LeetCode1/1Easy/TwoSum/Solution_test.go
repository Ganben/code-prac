package Solution

import (
	"math/rand"
	"reflect"
	"testing"
)

var cases = []struct {
	name string
	inputs [][]int
	expect []int
}{
	{"1 test 1", [][]int{{2,7,11,15}, {9}}, []int{0,1}},
	{"2 test 2", [][]int{{3,2,4}, {6}}, []int{1,2}},
	{"3 test 3", [][]int{{2,7,11,15}, {9}}, []int{0,1}},
	{"4 test 4", [][]int{{7,6,7,3,2,1,4,9,10}, {17}}, []int{0,8}},
}

func TestTowSum1(t *testing.T) {
	//


	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := TowSum1(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected: %v, got %v", c.expect, ret)
			}
		})
	}

}

func TestTowSum2(t *testing.T) {
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ret := TowSum3(c.inputs[0], c.inputs[1][0])
			if !reflect.DeepEqual(ret, c.expect) {
				t.Fatalf("expected %v, got %v", c.expect, ret)
			}
		})
	}
}

const N = 100

func BenchmarkTwoSum1(b *testing.B) {
	nums := []int{}
	for i := 0; i< N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TowSum1(nums, 9)
	}
}

func BenchmarkTwoSum2(b *testing.B) {
	nums := []int{}
	for i := 0; i < N; i++ {
		nums = append(nums, rand.Int())
	}
	nums = append(nums, 7, 2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		TowSum2(nums, 9)
	}
}
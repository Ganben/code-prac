package RangeSumQueryImmutable

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	ret := NumArray{}
	ret.arr = nums
	return ret
}

func (this *NumArray) SumRange(i int, j int) int {
	if j > len(this.arr) {
		return -1
	}
	ans := 0
	for xi := i; xi <= j; xi++ {
		ans += this.arr[xi]
	}
	return ans
}

type NumArray struct {
	arr []int
}

func Constructor(nums []int) NumArray {
	n := NumArray{}
	n.arr = make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		n.arr[i] = nums[i]
	}
	return n
}

func (this *NumArray) Update(i int, val int) {
	this.arr[i] = val
}

func (this *NumArray) SumRange(i int, j int) int {
	ans := 0
	for idx := i; idx <= j; idx++ {
		ans += this.arr[idx]
	}
	return ans
}
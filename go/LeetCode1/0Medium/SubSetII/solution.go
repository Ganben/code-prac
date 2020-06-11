package solution

var n int
var numsArray []int
var row []int
var ans [][]int

func subsetsWithDup(nums []int) [][]int {
	numsArray = nums
	n = len(nums)
	ans = [][]int{}
	ans = append(ans, []int{})
	row = []int{}
	// sort
	bubbleSort(n)
	backtrack(0)
	return ans
}

func backtrack(i int) {
	if i == n {
		return
	}
	for j := i; j < n; j++ {
		if j > i && numsArray[j] == numsArray[j-1] {
			continue
		}
		row = append(row, numsArray[j])
		ans = append(ans, row)
		backtrack(j + 1)
		row = row[:len(row)-1]
	}

}

func bubbleSort(n int) {
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			if numsArray[j] < numsArray[j-1] {
				numsArray[j], numsArray[j-1] = numsArray[j-1], numsArray[j]
			}
		}
	}
}

package main

func thirdMax(nums []int) int {
sort.Ints(nums)
return nums[len(nums)-3]
}

func main() {
	input := []int{2,2,3,1}
	fmt.Printf("%v\n", thirdMax(input))
}
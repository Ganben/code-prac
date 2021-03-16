package main

func numberOfArithmeticSlices(nums []int) int {
dp := make([]int, len(nums))
sum := 0
for i:=2; i<len(dp); i++ {
	if nums[i]-nums[i-1] == nums[i-1] - nums[i-2] {
		dp[i] = 1 + dp[i-1]
		sum += dp[i]
	}
}
return sum
}

func main() {
	a := []int{1,2,3,4}
	fmt.Printf("%v\n", numberOfArithmeticSlices(a))
}
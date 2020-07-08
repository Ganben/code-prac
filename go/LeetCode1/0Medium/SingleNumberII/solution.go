package SingleNumberII

func singleNumber(nums []int) int {
	seeOnce, seeTwice := 0, 0
	for _, n := range nums {
		seeOnce = ^seeTwice & (seeOnce ^ n)
		seeTwice = ^seeOnce & (seeTwice ^ n)
	}
	return seeOnce

}

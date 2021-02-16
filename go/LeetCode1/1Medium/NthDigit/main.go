package main

import "fmt"

func findNthDigit(n int) int {
	theN := 1
	count := 0
	for i := 1; count < n; i++ {
		count += lenN(i)
		theN = i
	}
	//
	fmt.Printf("N:%v,c:%v\n", theN, count)
	diff := count - n
	if diff == 0 {
		return theN % 10
	}
	res := 0
	narr := make([]int, lenN(theN))
	for i := 0; i < len(narr); i++ {
		res = theN % 10
		narr[len(narr)-1-i] = res
		theN /= 10

	}
	fmt.Printf("%v\n", narr)
	return narr[len(narr)-diff-1]
}

func lenN(n int) int {
	ans := 0
	// nn := n
	for n > 0 {
		n = n / 10
		ans++
	}
	// fmt.Printf("N=%v,l=%v\n", nn, ans)
	return ans
}
func main() {
	fmt.Printf("%v\n", findNthDigit(1000))
}

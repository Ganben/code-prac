package main

import (
	"fmt"
	"strconv"
)

func originalDigits(s string) string {
	count := map[rune]int{}
	for _, c := range s {
		count[rune(c)]++
	}
	//
	out := make([]int, 10)
	//
	out[0] = count['z'] //zero
	out[2] = count['w'] // two
	out[4] = count['u'] // four
	out[6] = count['x'] // six
	out[8] = count['g'] // eight
	//
	out[3] = count['h'] - out[8]                   // three - eight
	out[5] = count['f'] - out[4]                   // five - four
	out[7] = count['s'] - out[6]                   // seven - six
	out[9] = count['i'] - out[5] - out[6] - out[8] // nine, five, six, eight
	out[1] = count['n'] - out[7] - 2*out[9]        // one, seven, nine
	ans := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < out[i]; j++ {
			ans += strconv.Itoa(i)
		}
	}
	return ans
}

func main() {
	si := "owoztneoer"
	fmt.Printf("%v\n", originalDigits(si))
}

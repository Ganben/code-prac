package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := []int{}
	if len(p) > len(s) {
		return res
	}
	left, right := 0, 0
	countObj := map[byte]int{}
	count, missingType := 0, 0
	for i := 0; i < len(p); i++ {
		if _, exist := countObj[p[i]]; !exist {
			countObj[p[i]] = 1
			missingType++
		} else {
			countObj[p[i]]++
		}
	}
	//
	for right < len(s) {
		currValue := s[right]
		right++
		countObj[currValue]--
		if countObj[currValue] == 0 {
			count++
		}
		//
		for count == missingType {
			if right-left == len(p) {
				res = append([]int{left}, res...)
			}
			leftValue := s[left]
			left++
			countObj[leftValue]++
			if countObj[leftValue] > 0 {
				count--
			}
		}
	}
	return res

}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Printf("%v\n", findAnagrams(s, p))
}

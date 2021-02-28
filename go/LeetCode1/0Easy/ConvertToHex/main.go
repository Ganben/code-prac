package main

import "fmt"

func toHex(num int) string {
	bnum := int32(num)
	if bnum == 0 {
		return "0"
	}
	hexArr := "0123456789abcdef"
	ans := []string{}
	for i := 0; i < 8; i++ {
		if bnum == 0 {
			break
		}
		temp := bnum & 0xf
		ans = append(ans, string(hexArr[temp]))
		fmt.Printf("%v\n:remains:%v\n", ans, bnum)

		bnum >>= 4
	}
	sans := ""
	for i := len(ans) - 1; i >= 0; i-- {
		sans += ans[i]
	}
	return sans
}

func main() {
	fmt.Printf("result:%v\n", toHex(26))
	fmt.Printf("result:%v\n", toHex(-65535))
}

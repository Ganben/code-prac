package FracToRecurringDecimal

import (
	"fmt"
	"strconv"
)

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	res := ""
	if (numerator < 0) != (denominator < 0) {
		res += "-"
	}
	// overflow check
	var dividend int64
	var divisor int64
	dividend = abs(numerator)
	divisor = abs(denominator)
	res += strconv.Itoa(int(dividend / divisor))
	remainder := dividend % divisor
	if remainder == 0 {
		return res
	}
	res += "."
	nmap := map[int64]int{}
	for remainder != 0 {
		v, exist := nmap[remainder]

		if exist {
			//search and add ()
			f := ""
			res, f = nremove(res, v)
			res += "(" + f + ")"
			break
		}
		//
		nmap[remainder] = len(res)
		remainder *= 10
		n := int(remainder / divisor)
		res += strconv.Itoa(n)
		fmt.Printf("%v", res)
		remainder = remainder % divisor

	}
	return res

}

func abs(x int) int64 {
	if x < 0 {
		return int64(-1 * x)
	}
	return int64(x)
}

func nremove(res string, rem int) (string, string) {
	f := res[rem:]
	res = res[:rem]
	return res, f
}

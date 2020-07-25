package ExcelSheetColNumber

func titleToNumber(s string) int {
	iarr := []int{}
	for i := 0; i < len(s); i++ {
		b := int(s[i]) - int('A') + 1
		iarr = append(iarr, b)
	}
	n := 0
	d := 1
	for i := len(iarr) - 1; i >= 0; i-- {
		n += d * iarr[i]
		d *= 26
	}
	return n
}

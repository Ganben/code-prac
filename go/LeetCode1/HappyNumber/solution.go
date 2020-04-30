package solution

func isHappy(n int) bool {
	tmp := int2arr(n)
	for i:=0; i< 10000; i++ {
		ss := 0
		for _, k := range tmp {
			ss += k * k
		}
		if ss == 1 {
			return true
		}
		tmp = int2arr(ss)

	}
	return false
}

func int2arr(n int) []int {
	arr := []int{}
	for n > 10 {
		arr = append(arr, n % 10)
		// fmt.Printf("arr %v\n", arr)
		n = int(n / 10)
	}
	return arr
}
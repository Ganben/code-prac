package PascalTriangleII

func getRow(rowIndex int) []int {
	res := make([]int, rowIndex+1)
	for i, _ := range res {
		res[i] = 1
	}

	for i := 0; i < rowIndex+1; i++ {
		for j := i - 1; j > 0; j-- {
			res[j] = res[j] + res[j-1]
		}
	}
	return res
}

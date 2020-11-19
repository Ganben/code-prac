package RangeSumQuery2DImmutable

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	len1 := len(matrix)
	if len1 == 0 || len(matrix[0]) == 0 {
		var a [][]int
		var addarr NumMatrix
		addarr.matrix = a
		return addarr
	}
	len2 := len(matrix[0])
	all_num := make([][]int, len1)
	for i := 0; i < len(all_num); i++ {
		all_num[i] = make([]int, len2)
	}
	for i, _ := range matrix {
		for j, _ := range matrix[i] {
			if i == 0 && j == 0 {
				all_num[i][j] = matrix[i][j]
			} else if i == 0 {
				all_num[i][j] = matrix[i][j] + all_num[i][j-1]
			} else if j == 0 {
				all_num[i][j] = matrix[i][j] + all_num[i-1][j]
			} else {
				all_num[i][j] = all_num[i-1][j] + all_num[i][j-1] - all_num[i-1][j-1] + matrix[i][j]
			}

		}
	}
	var addarr = NumMatrix{all_num}
	return addarr
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	ans := 0
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			ans += this.matrix[i][j]
		}
	}
	return ans
}

func (this *NumMatrix) SumRegion(i int, j int, row2 int, col2 int) int {
	if len(this.array) == 0 {
		return 0
	}
	res := 0
	if i == 0 && j == 0 {
		res = this.array[row2][col2]
	} else if i == 0 {
		res = this.array[row2][col2] - this.array[row2][j-1]
	} else if j == 0 {
		res = this.array[row2][col2] - this.array[i-1][col2]
	} else {
		res = this.array[row2][col2] - this.array[row2][j-1] - this.array[i-1][col2] + this.array[i-1][j-1]
	}
	return res
}

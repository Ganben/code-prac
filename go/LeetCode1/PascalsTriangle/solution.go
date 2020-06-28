package PascalsTriangle

func generate(numRows int) [][]int {
	var res [][]int
	for i := 0; i < numRows; i++ {
		rows := []int{}
		tmp := make([]int, i+1)
		if i == 0 {
			rows = []int{1}
		} else {
			for j := 0; j <= i+1; j++ {
				if j == 0 {
					rows = append(rows, res[i][0])
				} else if j == i+1 {
					rows = append(rows, res[i][j-1])
				} else {
					rows = append(rows, res[i][j-1]+res[i][j])
				}
			}
		}
		copy(rows, tmp)
		res = append(res, tmp)
	}
}

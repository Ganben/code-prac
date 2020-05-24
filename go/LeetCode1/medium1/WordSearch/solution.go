package search

var visited [][]bool
var last []int
var ans bool
var directions [][]int

func exist(board [][]byte, word string) bool {
	ans = false
	directions = [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	for _, row := range board {
		tmp := make([]bool, len(row))
		visited = append(visited, tmp)
	}
	for i, row := range board {
		for j, c := range row {
			if c == word[0] {
				visited[i][j] = true
				last = []int{i, j}
				search(board, i, j, word[1:])
				visited[i][j] = false
			}
		}
	}
	return ans
}

func search(board [][]byte, i int, j int, word string) {
	if len(word) == 0 {
		ans = true
		return
	}
	// case 4 directions
	for _, c := range directions {
		ii := i + c[0]
		jj := j + c[1]
		if ii >= 0 && ii < len(board) && jj >= 0 && jj < len(board[0]) {
			if board[ii][jj] == word[0] {
				visited[ii][jj] = true
				search(board, ii, jj, word[1:])
				visited[ii][jj] = false
			}
		}

	}

}

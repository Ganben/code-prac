package search

var visited [][]bool
var ans bool

func exist(board [][]byte, word string) bool {
	ans = false
	for _, row := range board {
		tmp := make([]bool, len(row))
		visited = append(visited, tmp)
	}
	for i, row := range board {
		for j, c := range row {
			if c == word[0] {
				visited[i][j] = true
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

}

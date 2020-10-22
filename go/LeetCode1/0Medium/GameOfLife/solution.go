package GameOfLife

func gameOfLife(board [][]int) {
	rows := len(board)
	cols := len(board[0])
	neighbors := [][]int{
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
	}

	for i, row := range board {
		for j, _ := range row {
			liveNeighbors := 0
			for _, v := range neighbors {
				r := i + v[0]
				c := j + v[1]
				if r < rows && r >= 0 && c < cols && c >= 0 && abs(board[r][c]) == 1 {
					liveNeighbors++
				}
				//

			}
			if board[i][j] == 1 && (liveNeighbors < 2 || liveNeighbors > 3) {
				// lived but turn dead
				board[i][j] = -1

			}
			if board[i][j] == 0 && liveNeighbors == 3 {
				// dead but turn live
				board[i][j] = 2
			}

		}
	}
	// update
	for i, row := range board {
		for j, v := range row {
			if v > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}

}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

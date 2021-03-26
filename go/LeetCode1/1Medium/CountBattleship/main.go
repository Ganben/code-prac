package main

func countBattleship(board [][]byte) int {
	lenA := len(board)
	lenB := len(board[0])
	count := 0
	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' || j > 0 && board[i][j-1] == 'X' {
					continue

				}
				count++
			}
		}
	}
	return count
}

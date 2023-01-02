func solveSudoku(board [][]byte) {
	// O(9^n) n: number of '.'
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				var b byte
				for b = '1'; b <= '9'; b++ {
					if valid(board, i, j, b) {
						board[i][j] = b
						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func valid(board [][]byte, i, j int, b byte) bool {
	for k := 0; k < 9; k++ {
		if board[i][k] == b || board[k][j] == b || board[3*(i/3)+k/3][3*(j/3)+k%3] == b {
			return false
		}
	}
	return true
}
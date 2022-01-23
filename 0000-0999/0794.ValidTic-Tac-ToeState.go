func validTicTacToe(board []string) bool {
	turn := 0
	rows, cols := make([]int, 3), make([]int, 3)
	diag, antiDiag := 0, 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 'X' {
				turn++
				rows[i]++
				cols[j]++
				if i == j {
					diag++
				}
				if i+j == 2 {
					antiDiag++
				}
			} else if board[i][j] == 'O' {
				turn--
				rows[i]--
				cols[j]--
				if i == j {
					diag--
				}
				if i+j == 2 {
					antiDiag--
				}
			}
		}
	}
	xWin := rows[0] == 3 || rows[1] == 3 || rows[2] == 3 ||
		cols[0] == 3 || cols[1] == 3 || cols[2] == 3 ||
		diag == 3 || antiDiag == 3
	oWin := rows[0] == -3 || rows[1] == -3 || rows[2] == -3 ||
		cols[0] == -3 || cols[1] == -3 || cols[2] == -3 ||
		diag == -3 || antiDiag == -3
	if (turn != 0 && turn != 1) || (xWin && turn != 1) || (oWin && turn != 0) || (xWin && oWin) {
		return false
	}
	return true
}

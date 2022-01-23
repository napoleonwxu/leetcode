func countBattleships(board [][]byte) int {
	cnt := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				if (i > 0 && board[i-1][j] == 'X') || (j > 0 && board[i][j-1] == 'X') {
					continue
				}
				cnt++
			}
		}
	}
	return cnt
}
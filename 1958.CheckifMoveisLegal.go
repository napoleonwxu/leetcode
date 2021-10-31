func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	dr := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	dc := []int{-1, 0, 1, -1, 1, -1, 0, 1}
	for i := 0; i < len(dr); i++ {
		cnt := 1
		r, c := rMove+dr[i], cMove+dc[i]
		for r >= 0 && r < len(board) && c >= 0 && c < len(board[0]) {
			if board[r][c] == '.' {
				break
			}
			cnt++
			if board[r][c] == color {
				if cnt >= 3 {
					return true
				} else {
					break
				}
			}
			r, c = r+dr[i], c+dc[i]
		}
	}
	return false
}

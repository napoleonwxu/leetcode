func solveSudoku(board [][]byte)  {
    // O(9^m) m: number of '.'
    solve(board, 0, 0)
}

func solve(board [][]byte, row, col int) bool {
    for i := row; i < len(board); i, col = i+1, 0 {
        for j := col; j < len(board[0]); j++ {
            if board[i][j] == '.' {
                var n byte
                for n = '1'; n <= '9'; n++ {
                    if valid(board, i, j, n) {
                        board[i][j] = n
                        if solve(board, i, j+1) {
                            return true
                        } else {
                            board[i][j] = '.'
                        }
                    }
                }
                return false
            }
        }
    }
    return true
}

func valid(board [][]byte, i, j int, n byte) bool {
    for k := 0; k < 9; k++ {
        if board[i][k] == n || board[k][j] == n || board[3*(i/3)+k/3][3*(j/3)+k%3] == n {
            return false
        }
    }
    return true
}
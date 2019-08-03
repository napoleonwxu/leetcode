func gameOfLife(board [][]int)  {
    if len(board) == 0 || len(board[0]) == 0 {
        return
    }
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            life := -board[i][j]
            for x := i-1; x <= i+1; x++ {
                for y := j-1; y <= j+1; y++ {
                    if x >= 0 && x < m && y >= 0 && y < n {
                        life += board[x][y] & 1
                    }
                }
            }
            if life == 3 || (board[i][j] == 1 && life == 2) {
                board[i][j] += 2
            }
        }
    }
    for i := range board {
        for j := range board[i] {
            board[i][j] >>= 1
        }
    }
}
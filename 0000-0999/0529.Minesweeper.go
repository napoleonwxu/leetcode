func updateBoard(board [][]byte, click []int) [][]byte {
    m, n := len(board), len(board[0])
    if board[click[0]][click[1]] == 'M' {
        board[click[0]][click[1]] = 'X'
        return board
    }
    queue := [][]int{click}
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        queue = queue[1:]
        cnt := byte(0)
        for i := -1; i <= 1; i++ {
            for j := -1; j <= 1; j++ {
                if i == 0 && j == 0 {
                    continue
                }
                xi, yj := x+i, y+j
                if xi >= 0 && xi < m && yj >= 0 && yj < n && (board[xi][yj] == 'M' || board[xi][yj] == 'X') {
                    cnt++
                }
            }
        }
        if cnt > 0 {
            board[x][y] = '0' + cnt
            continue
        }
        board[x][y] = 'B'
        for i := -1; i <= 1; i++ {
            for j := -1; j <= 1; j++ {
                if i == 0 && j == 0 {
                    continue
                }
                xi, yj := x+i, y+j
                if xi >= 0 && xi < m && yj >= 0 && yj < n && board[xi][yj] == 'E' {
                    queue = append(queue, []int{xi, yj})
                    board[xi][yj] = 'B'  // Avoid to be added again.
                }
            }
        }
    }
    return board
}
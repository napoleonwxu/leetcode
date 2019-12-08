func tictactoe(moves [][]int) string {
    grid := make([][]byte, 3)
    for i := range grid {
        grid[i] = make([]byte, 3)
    }
    flag := true
    for _, m := range moves {
        if grid[m[0]][m[1]] != 0 {
            continue
        }
        if flag {
            grid[m[0]][m[1]] = 'X'
            if win(grid, m[0], m[1]) {
                return "A"
            }
        } else {
            grid[m[0]][m[1]] = 'O'
            if win(grid, m[0], m[1]) {
                return "B"
            }
        }
        flag = !flag
    }
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if grid[i][j] == 0 {
                return "Pending"
            }
        }
    }
    return "Draw"
}

func win(grid [][]byte, x, y int) bool {
    if grid[x][y] == grid[1][1] && 
       ((grid[1][1] == grid[0][0] && grid[1][1] == grid[2][2]) || 
        (grid[1][1] == grid[0][2] && grid[1][1] == grid[2][0])) {
        return true
    }
    if grid[x][y] == grid[0][y] && grid[x][y] == grid[1][y] && grid[x][y] == grid[2][y] {
        return true
    }
    if grid[x][y] == grid[x][0] && grid[x][y] == grid[x][1] && grid[x][y] == grid[x][2] {
        return true
    }
    return false
}
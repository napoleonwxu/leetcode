func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
        return -1
    }
    grid[0][0] = 1
    queue := [][]int{[]int{0, 0}}
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        queue = queue[1:]
        for i := -1; i <= 1; i++ {
            for j := -1; j <= 1; j++ {
                if i == 0 && j == 0 {
                    continue
                }
                nx, ny := x + i, y + j
                if nx >= 0 && nx < n && ny >= 0 && ny < n && grid[nx][ny] == 0 {
                    grid[nx][ny] = grid[x][y] + 1
                    queue = append(queue, []int{nx, ny})
                    if nx == n-1 && ny == n-1 {
                        return grid[nx][ny]
                    }
                }
            }
        }
    }
    return -1
}
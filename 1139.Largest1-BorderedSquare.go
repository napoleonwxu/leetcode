func largest1BorderedSquare(grid [][]int) int {
    // O(n^3) + O(n^2)
    m, n := len(grid), len(grid[0])
    hor := make([][]int, m)
    ver := make([][]int, m)
    for i := 0; i < m; i++ {
        hor[i] = make([]int, n)
        ver[i] = make([]int, n)
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                if i == 0 {
                    ver[i][j] = 1
                } else {
                    ver[i][j] = ver[i-1][j] + 1
                }
                if j == 0 {
                    hor[i][j] = 1
                } else {
                    hor[i][j] = hor[i][j-1] + 1
                }
            }
        }
    }
    Len := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            l := min(hor[i][j], ver[i][j])
            for ; l > Len; l-- {
                if hor[i-l+1][j] >= l && ver[i][j-l+1] >= l {
                    Len = l
                }
            }
        }
    }
    return Len * Len
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
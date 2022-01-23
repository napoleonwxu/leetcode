// ByteDance-Lark
func maxDistance(grid [][]int) int {
    n := len(grid)

    land := [][2]int{}
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 1 {
                land = append(land, [2]int{i, j})
            }
        }
    }
    if len(land) == 0 || len(land) == n*n {
        return -1
    }
    dis := 0
    // BFS, O(n^2) + O(n^2)
    dir := [][2]int{[2]int{-1, 0}, [2]int{0, -1}, [2]int{0, 1}, [2]int{1, 0}}
    for len(land) > 0 {
        nxt := [][2]int{}
        for _, pos := range land {
            for _, d := range dir {
                i, j := pos[0]+d[0], pos[1]+d[1]
                if i >= 0 && i < n && j >= 0 && j < n && grid[i][j] == 0 {
                    nxt = append(nxt, [2]int{i, j})
                    grid[i][j] = 1
                }
            }
        }
        land = nxt
        dis++
    }
    return dis-1
    /* Brute force: O(n^2*cnt(1)) + O(n^2), 1200+ms
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == 0 {
                tmp := n + n
                for _, pos := range land {
                    tmp = min(tmp, abs(pos[0]-i) + abs(pos[1]-j))
                }
                if tmp > dis {
                    dis = tmp
                }
            }
        }
    }
    return dis
    */
    /* DFS, O(n^2*cnt(1)) + O(n^2), TLE
    dis := make([][]int, n)
    for i := range dis {
        dis[i] = make([]int, n)
        for j := range dis[i] {
            dis[i][j] = n << 1
        }
    }
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                dfs(dis, i, j, 0)
            }
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            ans = max(ans, dis[i][j])
        }
    }
    if ans == 0 || ans == n<<1 {
        return -1
    }
    return ans
    */
}

func dfs(dis [][]int, i, j, d int) {
    if i < 0 || i >= len(dis) || j < 0 || j >= len(dis) || dis[i][j] == -1 {
        return
    }
    tmp := min(dis[i][j], d)
    dis[i][j] = -1
    dfs(dis, i-1, j, d+1)
    dfs(dis, i+1, j, d+1)
    dfs(dis, i, j-1, d+1)
    dfs(dis, i, j+1, d+1)
    dis[i][j] = tmp
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
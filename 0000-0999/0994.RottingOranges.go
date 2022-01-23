func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    fresh := 0
    rottens := [][2]int{}
    for i, vs := range grid {
        for j, v := range vs {
            if v == 2 {
                rottens = append(rottens, [2]int{i, j})
            } else if v == 1 {
                fresh++
            }
        }
    }
    if fresh == 0 {
        return 0
    }
    dir := [][2]int{{-1, 0}, {0, -1}, {0, 1}, {1, 0}}
    minute := 0
    for len(rottens) > 0 {
        next := [][2]int{}
        for _, r := range rottens {
            for _, d := range dir {
                i, j := r[0] + d[0], r[1] + d[1]
                if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 1 {
                    grid[i][j] = 2
                    fresh--
                    next = append(next, [2]int{i, j})
                }
            }
        }
        minute++
        rottens = next
    }
    if fresh > 0 {
        return -1
    }
    return minute - 1
}
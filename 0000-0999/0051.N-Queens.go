func solveNQueens(n int) [][]string {
	// O(n!) + O(n^2)
    ans := [][]string{}
    queue := make([][]string, n)
    for i := range queue {
        queue[i] = make([]string, n)
        for j := range queue[i] {
            queue[i][j] = "."
        }
    }
    dfs(queue, 0, &ans)
    return ans
}

func dfs(queue [][]string, row int, ans *[][]string) {
    if row == len(queue) {
        tmp := make([]string, len(queue))
        for i := range queue {
            tmp[i] = strings.Join(queue[i], "")
        }
        *ans = append(*ans, tmp)
        return
    }
    for col := 0; col < len(queue[row]); col++ {
        if valid(queue, row, col) {
            queue[row][col] = "Q"
            dfs(queue, row+1, ans)
            queue[row][col] = "."
        }
    }
}

func valid(queue [][]string, row, col int) bool {
    for i := row-1; i >= 0; i-- {
        if queue[i][col] == "Q" {
            return false
        }
    }
    for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
        if queue[i][j] == "Q" {
            return false
        }
    }
    for i, j := row-1, col+1; i >= 0 && j < len(queue[i]); i, j = i-1, j+1 {
        if queue[i][j] == "Q" {
            return false
        }
    }
    return true
}
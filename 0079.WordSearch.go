func exist(board [][]byte, word string) bool {
    if len(board) == 0 || len(board[0]) == 0 {
        return false
    }
    m, n := len(board), len(board[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dfs(board, i, j, word) {
                return true
            }
        }
    }
    return false
}

func dfs(board [][]byte, i, j int, word string) bool {
    if len(word) == 0 {
        return true
    }
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[0] {
        return false
    }
    board[i][j] ^= 1<<7
    found := dfs(board, i-1, j, word[1:]) || dfs(board, i+1, j, word[1:]) || dfs(board, i, j-1, word[1:]) || dfs(board, i, j+1, word[1:])
    board[i][j] ^= 1<<7
    return found
}
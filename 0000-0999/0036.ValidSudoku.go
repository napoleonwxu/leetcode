import "strconv"

func isValidSudoku(board [][]byte) bool {
    n := len(board)
    m := make(map[string]bool, n*n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if board[i][j] != '.' {
                str := "(" + string(board[i][j]) + ")"
                str1 := strconv.Itoa(i) + str
                str2 := str + strconv.Itoa(j)
                str3 := strconv.Itoa(i/3) + str + strconv.Itoa(j/3)
                if m[str1] || m[str2] || m[str3] {
                    return false
                }
                m[str1] = true
                m[str2] = true
                m[str3] = true
            }
        }
    }
    return true
}
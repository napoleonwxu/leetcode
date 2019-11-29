func oddCells(n int, m int, indices [][]int) int {
    // O(L) + O(m+n)
    row, col := make([]int, n), make([]int, m)
    cntRow, cntCol := 0, 0
    for _, indice := range indices {
        row[indice[0]] ^= 1
        col[indice[1]] ^= 1
        if row[indice[0]] == 1 {
            cntRow++
        } else {
            cntRow--
        }
        if col[indice[1]] == 1 {
            cntCol++
        } else {
            cntCol--
        }
    }
    return cntRow*m + cntCol*n - 2*cntRow*cntCol
    /* O(L*(m+n)+mn) + O(mn)
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, m)
    }
    for _, indice := range indices {
        for i := 0; i < n; i++ {
            matrix[i][indice[1]]++
        }
        for j := 0; j < m; j++ {
            matrix[indice[0]][j]++
        }
    }
    cnt := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if matrix[i][j]&1 == 1 {
                cnt++
            }
        }
    }
    return cnt
    */
}
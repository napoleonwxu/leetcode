func generateMatrix(n int) [][]int {
    matrix := make([][]int, n)
    for i := range matrix {
        matrix[i] = make([]int, n)
    }
    num := 1
    i1, i2 := 0, n-1
    j1, j2 := 0, n-1
    for i1 <= i2 && j1 <= j2 {
        for j := j1; j <= j2; j++ {
            matrix[i1][j] = num
            num++
        }
        for i := i1+1; i <= i2; i++ {
            matrix[i][j2] = num
            num++
        }
        if i1 < i2 && j1 < j2 {
            for j := j2-1; j >= j1; j-- {
                matrix[i2][j] = num
                num++
            }
            for i := i2-1; i > i1; i-- {
                matrix[i][j1] = num
                num++
            }
        }
        i1, i2 = i1+1, i2-1
        j1, j2 = j1+1, j2-1
    }
    return matrix
}
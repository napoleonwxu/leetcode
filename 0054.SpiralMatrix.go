func spiralOrder(matrix [][]int) []int {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return []int{}
    }
    i1, i2 := 0, len(matrix)-1
    j1, j2 := 0, len(matrix[0])-1
    spiral := []int{}
    for i1 <= i2 && j1 <= j2 {
        for j := j1; j <= j2; j++ {
            spiral = append(spiral, matrix[i1][j])
        }
        for i := i1+1; i <= i2; i++ {
            spiral = append(spiral, matrix[i][j2])
        }
        if i1 < i2 && j1 < j2 {
            for j := j2-1; j >= j1; j-- {
                spiral = append(spiral, matrix[i2][j])
            }
            for i := i2-1; i > i1; i-- {
                spiral = append(spiral, matrix[i][j1])
            }
        }
        i1, i2 = i1+1, i2-1
        j1, j2 = j1+1, j2-1
    }
    return spiral
}
// array: 560_SubarraySumEqualsK
func numSubmatrixSumTarget(matrix [][]int, target int) int {
    m, n := len(matrix), len(matrix[0])
    for _, row := range matrix {
        for j := 1; j < n; j++ {
            row[j] += row[j-1]
        }
    }
    cnt := 0
    for j := 0; j < n; j++ {
        for k := j; k < n; k++ {
            Map := make(map[int]int)
            Map[0] = 1
            sum := 0
            for i := 0; i < m; i++ {
                if j == 0 {
                    sum += matrix[i][k]
                } else {
                    sum += matrix[i][k] - matrix[i][j-1]
                }
                cnt += Map[sum-target]
                Map[sum]++
            }
        }
    }
    return cnt
}
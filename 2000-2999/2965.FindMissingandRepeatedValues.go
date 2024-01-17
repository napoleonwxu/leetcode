func findMissingAndRepeatedValues(grid [][]int) []int {
    n, sum := len(grid), 0
    m := make(map[int]bool, n)
    ans := make([]int, 2)
    for _, row := range grid {
        for _, num := range row {
            if m[num] {
                ans[0] = num
            } else {
                m[num] = true
                sum += num
            }
        }
    }
    ans[1] = (n*n+1)*(n*n)/2 - sum
    return ans
}
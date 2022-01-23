func matrixReshape(nums [][]int, r int, c int) [][]int {
    m, n := len(nums), len(nums[0])
    if m*n != r*c {
        return nums
    }
    ans := make([][]int, r)
    t := 0
    for i := range ans {
        ans[i] = make([]int, c)
        for j := range ans[i] {
            ans[i][j] = nums[t/n][t%n]
            t++
        }
    }
    return ans
}
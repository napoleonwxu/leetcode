func isArraySpecial(nums []int, queries [][]int) []bool {
    n := len(nums)
    parity := make([]int, n)
    for i := 1; i < n; i++ {
        parity[i] = parity[i-1]
        if nums[i]%2 == nums[i-1]%2 {
            parity[i]++
        }
    }
    ans := make([]bool, len(queries))
    for i, q := range queries {
        ans[i] = parity[q[0]] == parity[q[1]]
    }
    return ans
}
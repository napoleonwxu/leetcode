func findKDistantIndices(nums []int, key int, k int) []int {
    n := len(nums)
    ans := make([]int, 0, n)
    cur := -1
    for i, num := range nums {
        if num != key {
            continue
        }
        for j := i-k; j <= i+k; j++ {
            if j >= 0 && j < n && j > cur {
                ans = append(ans, j)
                cur = j
            }
        }
    }
    return ans
}
func findMatrix(nums []int) [][]int {
    cnt := make(map[int]int)
    ans := [][]int{}
    for _, num := range nums {
        cnt[num]++
        if cnt[num] > len(ans) {
            ans = append(ans, []int{})
        }
        idx := cnt[num] - 1
        ans[idx] = append(ans[idx], num)
    }
    return ans
}
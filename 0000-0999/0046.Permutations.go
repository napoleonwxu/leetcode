func permute(nums []int) [][]int {
    n := len(nums)
    perm := make([]int, n)
    used := make([]bool, n)
    ans := [][]int{}
    dfs(nums, perm, used, n, 0, &ans)
    return ans
}

func dfs(nums, perm []int, used []bool, n, cur int, ans *[][]int) {
    if cur == n {
        tmp := make([]int, n)
        copy(tmp, perm)
        *ans = append(*ans, tmp)
        return
    }
    for i := 0; i < n; i++ {
        if used[i] {
            continue
        }
        perm[cur] = nums[i]
        used[i] = true
        dfs(nums, perm, used, n, cur+1, ans)
        used[i] = false
    }
}
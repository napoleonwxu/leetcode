func permuteUnique(nums []int) [][]int {
    sort.Ints(nums)
    n := len(nums)
    perm := make([]int, n)
    used := make([]bool, n)
    ans := make([][]int, 0)
    dfs(nums, perm, used, 0, &ans)
    return ans
}

func dfs(nums, perm []int, used []bool, cur int, ans *[][]int) {
    if cur >= len(nums) {
        tmp := make([]int, len(perm))
        copy(tmp, perm)
        *ans = append(*ans, tmp)
        return
    }
    for i, num := range nums {
        if used[i] || (i > 0 && nums[i] == nums[i-1] && used[i-1]) {
            continue
        }
        perm[cur] = num
        used[i] = true
        dfs(nums, perm, used, cur+1, ans)
        used[i] = false
    }
}
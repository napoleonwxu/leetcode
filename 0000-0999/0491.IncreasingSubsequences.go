func findSubsequences(nums []int) [][]int {
    ans := [][]int{}
    dfs(nums, 0, &ans, []int{})
    return ans
}

func dfs(nums []int, idx int, ans *[][]int, holder []int) {
    n := len(holder)
    if len(holder) >= 2 {
        // *ans = append(*ans, holder) is wrong !!! Use deep copy.
        tmp := make([]int, n)
        copy(tmp, holder)
        *ans = append(*ans, tmp)
    }
    used := make(map[int]bool)
    for i := idx; i < len(nums); i++ {
        if used[nums[i]] {
            continue
        }
        if len(holder) == 0 || nums[i] >= holder[n-1] {
            used[nums[i]] = true
            holder = append(holder, nums[i])
            dfs(nums, i+1, ans, holder)
            holder = holder[:n]
        }
    }
}
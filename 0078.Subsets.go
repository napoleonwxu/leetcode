func subsets(nums []int) [][]int {
    p := 1 << uint(len(nums))
    ans := make([][]int, p)
    for i := 0; i < p; i++ {
        sub := []int{}
        for j := 0; j < len(nums); j++ {
            if i&(1<<uint(j)) != 0 {
                sub = append(sub, nums[j])
            }
        }
        ans[i] = sub
    }
    /*
    ans := [][]int{}
    dfs(nums, []int{}, 0, &ans)
    */
    return ans
}

func dfs(nums, path []int, i int, ans *[][]int) {
    tmp := make([]int, len(path))
    copy(tmp, path)
    *ans = append(*ans, tmp)
    for j := i; j < len(nums); j++ {
        dfs(nums, append(path, nums[j]), j+1, ans)
    }
}
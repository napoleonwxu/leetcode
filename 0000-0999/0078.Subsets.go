func subsets(nums []int) [][]int {
    n := len(nums)
    ans := make([][]int, 1<<n)
    for i := 0; i < 1<<n; i++ {
        for j, num := range nums {
            if i&(1<<j) != 0 {
                ans[i] = append(ans[i], num)
            }
        }
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
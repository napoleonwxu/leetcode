func numTilePossibilities(tiles string) int {
    nums := make([]int, 26)
    for _, ch := range tiles {
        nums[ch-'A']++
    }
    return dfs(nums)
}

func dfs(nums []int) int {
    cnt := 0
    for i := 0; i < 26; i++ {
        if nums[i] == 0 {
            continue
        }
        nums[i]--
        cnt++
        cnt += dfs(nums)
        nums[i]++
    }
    return cnt
}
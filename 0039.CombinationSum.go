func combinationSum(candidates []int, target int) [][]int {
    ans := [][]int{}
    dfs(candidates, []int{}, target, 0, &ans)
    return ans
}

func dfs(cand, path []int, target, idx int, ans *[][]int) {
    if target < 0 {
        return
    }
    if target == 0 {
        tmp := make([]int, len(path))
        copy(tmp, path)
        *ans = append(*ans, tmp)
        return
    }
    for i := idx; i < len(cand); i++ {
        dfs(cand, append(path, cand[i]), target-cand[i], i, ans)
    }
}

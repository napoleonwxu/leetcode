// 159. Longest Substring with At Most Two Distinct Characters
func totalFruit(tree []int) int {
    // O(n) + O(3)
    cnt := make(map[int]int)
    i, ans := 0, 0
    for j := 0; j < len(tree); j++ {
        cnt[tree[j]]++
        for ; len(cnt) > 2; i++ {
            cnt[tree[i]]--
            if cnt[tree[i]] == 0 {
                delete(cnt, tree[i])
            }
        }
        ans = max(ans, j-i+1)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func longestSubsequence(arr []int, difference int) int {
    cnt := make(map[int]int)
    ans := 0
    for _, num := range arr {
        cnt[num] = cnt[num-difference] + 1
        ans = max(ans, cnt[num])
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
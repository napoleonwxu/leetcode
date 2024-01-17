const MAX = 1e7

func largestCombination(candidates []int) int {
    ans := 0
    for b := 1; b <= MAX; b <<= 1 {
        cnt := 0
        for _, num := range candidates {
            if num&b > 0 {
                cnt++
            }
        }
        ans = max(ans, cnt)
    }
    return ans
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
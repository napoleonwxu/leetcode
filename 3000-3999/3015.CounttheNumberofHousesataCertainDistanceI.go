func countOfPairs(n int, x int, y int) []int {
    if x > y {
        x, y = y, x
    }
    ans := make([]int, n)
    for i := 1; i <= n; i++ {
        for j := 1; j < i; j++ {
            idx := min(i-j, abs(x-j)+abs(y-i)+1)
            if idx > 0 {
                ans[idx-1] += 2
            }
        }
    }
    return ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
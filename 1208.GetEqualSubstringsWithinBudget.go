func equalSubstring(s string, t string, maxCost int) int {
    n := len(s)
    // O(n) + O(1)
    i, j := 0, 0
    for ; j < n; j++ {
        maxCost -= abs(int(s[j]) - int(t[j]))
        if maxCost < 0 {
            maxCost += abs(int(s[i]) - int(t[i]))
            i++
        }
    }
    return j-i
    /* O(n) + O(n)
    cost := make([]int, n)
    for i := range s {
        cost[i] = abs(int(s[i]) - int(t[i]))
    }
    ans, sum := 0, 0
    i := 0
    for j := 0; j < n; j++ {
        sum += cost[j]
        for ; sum > maxCost; i++ {
            sum -= cost[i]
        }
        ans = max(ans, j-i+1)
    }
    return ans
    */
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
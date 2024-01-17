func closetTarget(words []string, target string, startIndex int) int {
    n := len(words)
    ans := n
    for i, word := range words {
        if word == target {
            ans = min(ans, abs(i-startIndex))
            ans = min(ans, n-abs(i-startIndex))
        }
    }
    if ans == n {
        return -1
    }
    return ans
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
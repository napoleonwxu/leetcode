func minHeightShelves(books [][]int, shelf_width int) int {
    n := len(books)
    dp := make([]int, n+1)
    dp[0] = 0
    for i := 1; i <= n; i++ {
        width, height := books[i-1][0], books[i-1][1]
        dp[i] = dp[i-1] + height
        for j := i-1; j > 0; j-- {
            width += books[j-1][0]
            height = max(height, books[j-1][1])
            if width > shelf_width {
                break
            }
            dp[i] = min(dp[i], dp[j-1]+height)
        }
    }
    return dp[n]
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
func stoneGameII(piles []int) int {
    n := len(piles)
    sum := make([]int, n)
    sum[n-1] = piles[n-1]
    for i := n-2; i >= 0; i-- {
        sum[i] = sum[i+1] + piles[i]
    }
    memo := make([][]int, n)
    for i := range memo {
        memo[i] = make([]int, n)
    }
    return dfs(memo, sum, 0, 1)
}

func dfs(memo [][]int, sum []int, i, M int) int {
    if i == len(sum) {
        return 0
    }
    if 2*M >= len(sum)-i {
        return sum[i]
    }
    if memo[i][M] != 0 {
        return memo[i][M]
    }
    nxt := math.MaxInt32
    for x := 1; x <= 2*M; x++ {
        // the min value the next player can get
        nxt = min(nxt, dfs(memo, sum, i+x, max(M, x)))
    }
    memo[i][M] = sum[i] - nxt
    return memo[i][M]
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
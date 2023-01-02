func maximumCandies(candies []int, k int64) int {
    left, right := 0, 0
    for _, candy := range candies {
        right = max(right, candy)
    }
    for left < right {
        mid := (left + right + 1) / 2
        if canSplit(candies, k, mid) {
            left = mid
        } else {
            right = mid - 1
        }
    }
    return left
}

func canSplit(candies []int, k int64, per int) bool {
    var cnt int64
    for _, candy := range candies {
        cnt += int64(candy/per)
    }
    return cnt >= k
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
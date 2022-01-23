func minEatingSpeed(piles []int, H int) int {
    // O(NlogM), N: len(piles), M: max(piles)
    left, right := 1, 1
    for _, pile := range piles {
        right = max(right, pile)
    }
    for left < right {
        mid := (left+right) >> 1
        if costHour(piles, mid) <= H {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func costHour(piles []int, k int) int {
    hour := 0
    for _, pile := range piles {
        hour += (pile+k-1)/k
    }
    return hour
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func shipWithinDays(weights []int, D int) int {
    // O(nlogn)
    left, right := 0, 0
    for _, w := range weights {
        left = max(left, w)
        right += w
    }
    for left < right {
        mid := (left+right) >> 1
        if shipDay(weights, mid) <= D {
            right = mid
        } else {
            left = mid + 1
        }
    }
    return left
}

func shipDay(weights []int, cap int) int {
    day, sum := 0, 0
    for _, w := range weights {
        if sum+w > cap {
            day++
            sum = 0
        }
        sum += w
    }
    return day+1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
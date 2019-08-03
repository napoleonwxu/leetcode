func maxDistToClosest(seats []int) int {
    n := len(seats)
    left, right := 0, n-1
    for ; left < n && seats[left] == 0; left++ {}
    for ; right >= 0 && seats[right] == 0; right-- {}
    empty, tmp := 0, 0
    for i := left+1; i < right; i++ {
        if seats[i] == 0 {
            tmp++
            empty = max(empty, tmp)
        } else {
            tmp = 0
        }
    }
    return max(max(left, n-right-1), (empty+1)>>1)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
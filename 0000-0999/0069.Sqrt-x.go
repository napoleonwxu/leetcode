func mySqrt(x int) int {
    if x <= 1 {
        return x
    }
    left, right := 1, x
    for left <= right {
        mid := (left + right) >> 1
        pro := mid * mid
        if pro <= x && (mid+1)*(mid+1) > x {
            return mid
        } else if pro < x {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
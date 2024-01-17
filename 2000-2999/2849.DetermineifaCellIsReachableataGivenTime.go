func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
    x, y := abs(fx-sx), abs(fy-sy)
    if x == 0 && y == 0 {
        return t != 1
    }
    return t >= max(x, y)
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
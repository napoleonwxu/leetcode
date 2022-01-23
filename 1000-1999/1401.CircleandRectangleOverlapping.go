func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
    x_close := clamp(x_center, x1, x2)
    y_close := clamp(y_center, y1, y2)
    x_dis := x_close - x_center
    y_dis := y_close - y_center
    return x_dis*x_dis + y_dis*y_dis <= radius*radius
}

func clamp(val, mi, ma int) int {
    return max(mi, min(ma, val))
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
func minMovesToCaptureTheQueen(a int, b int, c int, d int, e int, f int) int {
    if a == e || b == f {
        if a == e && a == c && (d-b)*(d-f) < 0 {
            return 2
        }
        if b == f && b == d && (c-a)*(c-e) < 0 {
            return 2
        }
        return 1
    }
    if abs(c-e) == abs(d-f) {
        if abs(c-a) == abs(d-b) && (b-d)*(b-f) < 0 {
            return 2
        }
        return 1
    }
    return 2
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
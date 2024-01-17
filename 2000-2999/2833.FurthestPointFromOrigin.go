func furthestDistanceFromOrigin(moves string) int {
    l, or := 0, 0
    for _, move := range moves {
        if move == 'L' {
            l++
        } else if move == 'R' {
            l--
        } else {
            or++
        }
    }
    return abs(l) + or
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
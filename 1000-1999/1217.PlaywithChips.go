func minCostToMoveChips(chips []int) int {
    pos := [2]int{}
    for _, chip := range chips {
        pos[chip&1]++
    }
    return min(pos[0], pos[1])
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
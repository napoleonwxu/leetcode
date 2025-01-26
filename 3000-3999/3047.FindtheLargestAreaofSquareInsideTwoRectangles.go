func largestSquareArea(bottomLeft [][]int, topRight [][]int) int64 {
    n := len(bottomLeft)
    maxLen := 0
    for i := 0; i < n; i++ {
        for j := i+1; j < n; j++ {
            maxLen = max(maxLen, getMaxLen(bottomLeft[i], topRight[i], bottomLeft[j], topRight[j]))
        }
    }
    return int64(maxLen) * int64(maxLen)
}

func getMaxLen(p11, p12, p21, p22 []int) int {
    x1, y1 := max(p11[0], p21[0]), max(p11[1], p21[1])
    x2, y2 := min(p12[0], p22[0]), min(p12[1], p22[1])
    return min(max(0, x2-x1), max(0, y2-y1))    
}

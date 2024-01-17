func areaOfMaxDiagonal(dimensions [][]int) int {
    ans := 0
    sq := int64(0)
    for _, d := range dimensions {
        tmp := int64(d[0])*int64(d[0]) + int64(d[1])*int64(d[1])
        area := d[0] * d[1]
        if tmp > sq || (tmp == sq && area > ans) {
            sq = tmp
            ans = area
        }
    }
    return ans
}

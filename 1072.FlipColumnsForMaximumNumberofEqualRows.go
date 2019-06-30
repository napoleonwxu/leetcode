func maxEqualRowsAfterFlips(matrix [][]int) int {
    Map := make(map[int]int)
    for _, row := range(matrix) {
        key := 0
        for _, c := range row {
            key = (key << 1) | (row[0] ^ c)
        }
        Map[key]++
    }
    ans := 0
    for _, v := range Map {
        if v > ans {
            ans = v
        }
    }
    return ans
}
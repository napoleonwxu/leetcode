func maxHeightOfTriangle(red int, blue int) int {
    r1, r2 := 0, 0
    h := 0
    for (r1 <= red && r2 <= blue) || (r1 <= blue && r2 <= red) {
        h++
        if h%2 == 1 {
            r1 += h
        } else {
            r2 += h
        }
    }
    return h-1
}
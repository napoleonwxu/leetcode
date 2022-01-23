func brokenCalc(X int, Y int) int {
    cnt := 0
    for Y > X {
        if Y&1 == 0 {
            Y >>= 1
        } else {
            Y += 1
        }
        cnt++
    }
    cnt += X - Y
    return cnt
}

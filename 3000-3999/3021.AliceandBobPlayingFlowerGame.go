func flowerGame(n int, m int) int64 {
    /*
    Alice wins when x+y is odd:
    x is odd: n/2 + n%2; even: n/2
    y is odd: m/2 + m%2; even: m/2
    (n/2 + n%2) * m/2 + n/2 * (m/2 + m%2) = n*m/2
    */
    return int64(n) * int64(m) / 2
}
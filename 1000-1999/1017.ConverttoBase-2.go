func baseNeg2(N int) string {
    if N == 0 {
        return "0"
    }
    ans := ""
    /* https://www.geeksforgeeks.org/convert-number-negative-base-representation/
    for N != 0 {
        remainder := N % -2
        N /= -2
        if remainder < 0 {
            remainder += 2
            N += 1
        }
        ans = strconv.Itoa(remainder) + ans
    }
    */
    for N != 0 {
        ans = strconv.Itoa(N & 1) + ans
        N = -(N>>1)
    }
    return ans
}
func divide(dividend int, divisor int) int {
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }
    sign := 1
    if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
        sign = -1
    }
    dividend, divisor = abs(dividend), abs(divisor)
    ans := 0
    for dividend >= divisor {
        tmp, m := divisor, 1
        for tmp << 1 <= dividend {
            tmp <<= 1
            m <<= 1
        }
        dividend -= tmp
        ans += m
    }
    return sign * ans
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}
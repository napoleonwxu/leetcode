func myAtoi(s string) int {
    sign, num := 1, 0
    i, n :=0, len(s)
    for ; i < n && s[i] == ' '; i++ {}
    if i < n && (s[i] == '-' || s[i] == '+') {
        if s[i] == '-' {
            sign = -1
        }
        i++
    }
    for ; i < n && s[i] >= '0' && s[i] <= '9'; i++ {
        d := int(s[i]-'0')
        if num > math.MaxInt32/10 || (num == math.MaxInt32/10 && d > math.MaxInt32%10) {
            if sign == 1 {
                return math.MaxInt32
            }
            return math.MinInt32
        }
        num = 10*num + d
    }
    return sign*num
}
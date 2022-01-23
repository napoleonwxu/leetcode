func myAtoi(str string) int {
    n := len(str)
    i := 0
    for ; i < n && str[i] == ' '; i++ {}
    if i >= n {
        return 0
    }
    sign := 1
    if str[i] == '-' {
        sign = -1
        i++
    } else if str[i] == '+' {
        i++
    }
    num := 0
    for ; i < n && str[i] >= '0' && str[i] <= '9'; i++ {
        d := int(str[i]-'0')
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
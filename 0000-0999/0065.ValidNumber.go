func isNumber(s string) bool {
    s = strings.TrimSpace(s)
    num := false
    point := false
    e := false
    for i, ch := range s {
        if ch >= '0' && ch <= '9' {
            num = true
        } else if ch == '.' {
            if point || e {
                return false
            }
            point = true
        } else if ch == 'e' {
            if e || !num {
                return false
            }
            e = true
            num = false
        } else if ch == '+' || ch == '-' {
            if i != 0 && s[i-1] != 'e' {
                return false
            }
        } else {
            return false
        }
    }
    return num
}
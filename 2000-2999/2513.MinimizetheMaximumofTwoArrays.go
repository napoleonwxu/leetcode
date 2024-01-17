func minimizeSet(divisor1 int, divisor2 int, uniqueCnt1 int, uniqueCnt2 int) int {
    divisor3 := lcm(divisor1, divisor2)
    l, r := 2, int(1e10)
    for l < r {
        m := (l + r) / 2
        if check(m, divisor1, divisor2, divisor3, uniqueCnt1, uniqueCnt2) {
            r = m
        } else {
            l = m + 1
        }
    }
    return l
}

func check(num, d1, d2, d3, c1, c2 int) bool {
    if num-num/d1 < c1 {
        return false
    }
    if num-num/d2 < c2 {
        return false
    }
    if num-num/d3 < c1+c2 {
        return false
    }
    return true
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }
    return x
}

func lcm(x, y int) int {
    return x * y / gcd(x, y)
}
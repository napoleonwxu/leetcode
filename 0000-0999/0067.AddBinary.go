func addBinary(a string, b string) string {
    ans := ""
    i, j := len(a)-1, len(b)-1
    add := byte(0)
    for i >= 0 || j >= 0 || add > 0 {
        if i >= 0 {
            add += a[i] - '0'
            i--
        }
        if j >= 0 {
            add += b[j] - '0'
            j--
        }
        ans = strconv.Itoa(int(add&1)) + ans
        add >>= 1
    }
    return ans
    /*
    ans := make([]byte, max(len(a), len(b))+1)
    i, j := len(a)-1, len(b)-1
    k := len(ans) - 1
    c := byte(0)
    for i >= 0 || j >= 0 || c > 0 {
        if i >= 0 {
            c += a[i] - '0'
            i--
        }
        if j >= 0 {
            c += b[j] - '0'
            j--
        }
        ans[k] = c&1 + '0'
        k--
        c >>= 1
    }
    return string(ans[k+1:])
    */
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
func addBinary(a string, b string) string {
    ans := ""
    c := byte(0)
    i, j := len(a)-1, len(b)-1
    for i >= 0 || j >= 0 {
        if i >= 0 {
            c += a[i] - '0'
            i--
        }
        if j >= 0 {
            c += b[j] - '0'
            j--
        }
        if c & 1 == 1 {
            ans = "1" + ans
        } else {
            ans = "0" + ans
        }
        c >>= 1
    }
    if c == 1 {
        ans = "1" + ans
    }
    return ans
}
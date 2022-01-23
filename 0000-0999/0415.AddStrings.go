func addStrings(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    ans := make([]byte, max(m, n)+1)
    i, j, k := m-1, n-1, len(ans)-1
    c := byte(0)
    for i >= 0 || j >= 0 || c > 0 {
        if i >= 0 {
            c += num1[i] - '0'
            i--
        }
        if j >= 0 {
            c += num2[j] - '0'
            j--
        }
        ans[k] = c%10 + '0'
        c /= 10
        k--
    }
    return string(ans[k+1:])
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
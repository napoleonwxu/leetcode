func addStrings(num1 string, num2 string) string {
    ans := make([]byte, max(len(num1),len(num2))+1)
    i, j, k := len(num1)-1, len(num2)-1, len(ans)
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
        k--
        ans[k] = c%10 + '0'
        c /= 10
    }
    return string(ans[k:])
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
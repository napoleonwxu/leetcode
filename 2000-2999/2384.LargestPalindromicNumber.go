import (
    "strconv"
    "strings"
)

func largestPalindromic(num string) string {
    cnts := make([]int, 10)
    for _, ch := range num {
        cnts[ch-'0']++
    }
    pre := ""
    for i := 9; i >= 0; i-- {
        if i > 0 || len(pre) > 0 {
            pre += strings.Repeat(strconv.Itoa(i), cnts[i]/2)
            cnts[i] %= 2
        }
    }
    mid := ""
    for i := 9; i >= 0; i-- {
        if cnts[i] > 0 {
            mid = strconv.Itoa(i)
            break
        }
    }
    return pre + mid + reverse(pre)
}

func reverse(s string) string {
    n := len(s)
    r := make([]byte, n)
    for i, ch := range []byte(s) {
        r[n-i-1] = ch
    }
    return string(r)
}
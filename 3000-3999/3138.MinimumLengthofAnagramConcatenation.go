func minAnagramLength(s string) int {
    n := len(s)
    for i := 1; i < n; i++ {
        if n%i == 0 && check(s, i) {
            return i
        }
    }
    return n
}

func check(s string, k int) bool {
    cnts := make([]int, 26)
    for i := 0; i < k; i++ {
        cnts[int(s[i]-'a')]++
    }
    for i := k; i < len(s); i += k {
        tmp := make([]int, 26)
        for j := i; j < i+k; j++ {
            tmp[int(s[j]-'a')]++
        }
        for j := 0; j < 26; j++ {
            if cnts[j] != tmp[j] {
                return false
            }
        }
    }
    return true
}
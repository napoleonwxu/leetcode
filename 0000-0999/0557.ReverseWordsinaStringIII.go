func reverseWords(s string) string {
    n := len(s)
    b := []byte(s)
    i := 0
    for i < n {
        j := i
        for j < n {
            if b[j] == ' ' {
                break
            }
            j++
        }
        for l, r := i, j-1; l < r; l, r = l+1, r-1 {
            b[l], b[r] = b[r], b[l]
        }
        i = j+1
    }
    return string(b)
}
func longestDecomposition(text string) int {
    cnt := 0
    left, right := "", ""
    i, j := 0, len(text)-1
    for ; i < j; i, j = i+1, j-1 {
        left += string(text[i])
        right = string(text[j]) + right
        if left == right {
            cnt += 2
            left, right = "", ""
        }
    }
    if i == j || left != "" {
        cnt += 1
    }
    return cnt
    /*
    n := len(text)
    if n == 0 {
        return 0
    }
    for i := 1; i <= n/2; i++ {
        if text[:i] == text[n-i:] {
            return 2 + longestDecomposition(text[i:n-i])
        }
    }
    return 1
    */
}
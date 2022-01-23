func minRemoveToMakeValid(s string) string {
    n := len(s)
    chars := make([]byte, n)
    balance := 0
    left, right := -1, n
    for i := 0; i < n; i++ {
        if s[i] == '(' {
            balance++
        } else if s[i] == ')' {
            if balance == 0 {
                continue
            }
            balance--
        }
        left++
        chars[left] = s[i]
    }
    for balance > 0 {
        if chars[left] == '(' {
            balance--
        } else {
            right--
            chars[right] = chars[left]
        }
        left--
    }
    return string(append(chars[:left+1], chars[right:]...))
}
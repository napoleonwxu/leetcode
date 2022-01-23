func isPalindrome(s string) bool {
    s = strings.ToLower(s)
    for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
        for ; left < right && !isAlphanum(s[left]); left++ {}
        for ; left < right && !isAlphanum(s[right]); right-- {}
        if s[left] != s[right] {
            return false
        }
    }
    return true
}

func isAlphanum(ch byte) bool {
    if (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') {
        return true
    }
    return false
}
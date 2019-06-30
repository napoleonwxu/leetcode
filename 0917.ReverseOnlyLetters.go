func reverseOnlyLetters(S string) string {
    s := []rune(S)
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        for ; i < j && (s[i] < 'A' || (s[i] > 'Z' && s[i] < 'a') || s[i] > 'z'); i++ {}
        for ; i < j && (s[j] < 'A' || (s[j] > 'Z' && s[j] < 'a') || s[j] > 'z'); j-- {}
        s[i], s[j] = s[j], s[i]
    }
    return string(s)
}
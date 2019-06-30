func detectCapitalUse(word string) bool {
    if len(word) <= 1 {
        return true
    }
    if word[0] >= 'a' && word[1] <= 'Z' {
        return false
    }
    upper := word[1] <= 'Z'
    for i := 2; i < len(word); i++ {
        if (upper && word[i] >= 'a') || (!upper && word[i] <= 'Z') {
            return false
        }
    }
    return true
}
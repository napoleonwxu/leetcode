func repeatedStringMatch(A string, B string) int {
    times := int(math.Ceil(float64(len(B))/float64(len(A))))
    if strings.Contains(strings.Repeat(A, times), B) {
        return times
    }
    if strings.Contains(strings.Repeat(A, times+1), B) {
        return times + 1
    }
    return -1
}
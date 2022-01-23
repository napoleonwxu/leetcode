func generateTheString(n int) string {
    return strings.Repeat("a", n+n&1-1) + strings.Repeat("b", 1-n&1)
}
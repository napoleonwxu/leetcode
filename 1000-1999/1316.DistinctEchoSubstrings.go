func distinctEchoSubstrings(text string) int {
    n := len(text)
    // O(n^3)
    Map := make(map[string]bool)
    for i := 0; i <= n-2; i++ {
        for j := 1; i+2*j <= n; j++ {
            if text[i:i+j] == text[i+j:i+2*j] {
                Map[text[i:i+2*j]] = true
            }
        }
    }
    return len(Map)
}
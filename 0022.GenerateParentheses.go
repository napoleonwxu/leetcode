func generateParenthesis(n int) []string {
    if n == 0 {
        return []string{""}
    }
    ans := []string{}
    for i := 0; i < n; i++ {
        for _, left := range generateParenthesis(i) {
            for _, right := range generateParenthesis(n-i-1) {
                ans = append(ans, "(" + left + ")" + right)
            }
        }
    }
    return ans
}
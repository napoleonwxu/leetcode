func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{""}
	}
	ans := []string{}
	//dfs(n, n, "", &ans)
	for i := 0; i < n; i++ {
		for _, left := range generateParenthesis(i) {
			for _, right := range generateParenthesis(n - i - 1) {
				ans = append(ans, "("+left+")"+right)
			}
		}
	}
	return ans
}

func dfs(left, right int, path string, ans *[]string) {
	if left > right || left < 0 || right < 0 {
		return
	}
	if left == 0 && right == 0 {
		*ans = append(*ans, path)
	}
	dfs(left-1, right, path+"(", ans)
	dfs(left, right-1, path+")", ans)
}
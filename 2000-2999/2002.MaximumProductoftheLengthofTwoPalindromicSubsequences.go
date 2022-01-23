var ans int

func maxProduct(s string) int {
	ans = 0
	chs := []byte(s)
	sub1, sub2 := make([]byte, 0, len(s)), make([]byte, 0, len(s))
	dfs(chs, sub1, sub2, 0)
	return ans
}

func dfs(chs, sub1, sub2 []byte, idx int) {
	if idx >= len(chs) {
		if isPalindromic(sub1) && isPalindromic(sub2) {
			ans = max(ans, len(sub1)*len(sub2))
		}
		return
	}
	dfs(chs, sub1, sub2, idx+1)
	dfs(chs, append(sub1, chs[idx]), sub2, idx+1)
	dfs(chs, sub1, append(sub2, chs[idx]), idx+1)
}

func isPalindromic(bytes []byte) bool {
	n := len(bytes)
	for i := 0; i < n/2; i++ {
		if bytes[i] != bytes[n-i-1] {
			return false
		}
	}
	return true
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
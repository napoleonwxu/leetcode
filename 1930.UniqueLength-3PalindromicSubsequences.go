func countPalindromicSubsequence(s string) int {
	idx := make([][]int, 26)
	for i := 0; i < len(idx); i++ {
		idx[i] = []int{-1, -1}
	}
	for j := 0; j < len(s); j++ {
		i := s[j] - 'a'
		if idx[i][0] == -1 {
			idx[i][0] = j
		} else {
			idx[i][1] = j
		}
	}
	cnt := 0
	for i := 0; i < len(idx); i++ {
		m := make(map[byte]bool)
		for j := idx[i][0] + 1; j < idx[i][1]; j++ {
			m[s[j]] = true
		}
		cnt += len(m)
	}
	return cnt
}
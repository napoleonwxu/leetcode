const MOD = 1e9 + 7

func numberOfWays(startPos int, endPos int, k int) int {
	if endPos < startPos-k || endPos > startPos+k {
		return 0
	}
	diff := endPos - startPos
	startPos = k + 1
	endPos = startPos + diff
	cur := make([]int, 2*k+3)
	cur[startPos] = 1
	for step := 1; step <= k; step++ {
		nxt := make([]int, 2*k+3)
		for i := 1; i <= 2*k+1; i++ {
			nxt[i] = (cur[i-1] + cur[i+1]) % MOD
		}
		cur = nxt
	}
	return cur[endPos]
}
func evenOddBit(n int) []int {
	ans := make([]int, 2)
	for i := 0; n > 0; n, i = n/2, i+1 {
		if n&1 == 1 {
			if i&1 == 0 {
				ans[0]++
			} else {
				ans[1]++
			}
		}
	}
	return ans
}
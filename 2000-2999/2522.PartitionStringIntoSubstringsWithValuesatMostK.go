func minimumPartition(s string, k int) int {
	cnt, num := 0, 0
	for _, ch := range s {
		num = 10*num + int(ch-'0')
		if num > k {
			cnt++
			num = int(ch - '0')
		}
		if num > k {
			return -1
		}
	}
	return cnt + 1
}

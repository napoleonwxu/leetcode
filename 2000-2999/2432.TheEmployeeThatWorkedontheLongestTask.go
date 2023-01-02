func hardestWorker(n int, logs [][]int) int {
	ans := 0
	pre, time := 0, 0
	for _, log := range logs {
		tmp := log[1] - pre
		if tmp > time || (tmp == time && log[0] < ans) {
			time, ans = tmp, log[0]
		}
		pre = log[1]
	}
	return ans
}
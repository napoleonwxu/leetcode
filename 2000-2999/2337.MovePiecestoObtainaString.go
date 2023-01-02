func canChange(start string, target string) bool {
	n := len(start)
	i, j := 0, 0
	for i <= n && j <= n {
		for ; i < n && start[i] == '_'; i++ {
		}
		for ; j < n && target[j] == '_'; j++ {
		}
		if i == n || j == n {
			return i == n && j == n
		}
		if start[i] != target[j] {
			return false
		}
		if start[i] == 'L' {
			if i < j {
				return false
			}
		} else {
			if i > j {
				return false
			}
		}
		i, j = i+1, j+1
	}
	return true
}
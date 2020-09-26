func busyStudent(startTime []int, endTime []int, queryTime int) int {
	cnt := 0
	for i := range startTime {
		if queryTime >= startTime[i] && queryTime <= endTime[i] {
			cnt++
		}
	}
	return cnt
}
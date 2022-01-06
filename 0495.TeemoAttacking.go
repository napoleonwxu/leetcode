func findPoisonedDuration(timeSeries []int, duration int) int {
	ans := 0
	for i := 1; i < len(timeSeries); i++ {
		ans += min(timeSeries[i]-timeSeries[i-1], duration)
	}
	return ans + duration
	/*
	   start, end := 0, 0
	   for _, t := range timeSeries {
	       start, end = max(t, end), t+duration
	       ans += end - start
	   }
	   return ans
	*/
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func buildArray(target []int, n int) []string {
	ans := []string{}
	i := 1
	for _, t := range target {
		for ; i < t; i++ {
			ans = append(ans, "Push")
			ans = append(ans, "Pop")
		}
		i++
		ans = append(ans, "Push")
	}
	return ans
}
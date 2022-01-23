import "sort"

func numberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[i][1] > properties[j][1]
		}
		return properties[i][0] < properties[j][0]
	})
	cnt, n := 0, len(properties)
	mx := properties[n-1][1]
	for i := n - 2; i >= 0; i-- {
		if properties[i][1] < mx {
			cnt++
		}
		mx = max(mx, properties[i][1])
	}
	return cnt
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
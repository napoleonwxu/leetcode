import "sort"

const MOD = 1e9 + 7

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	m, n := len(horizontalCuts), len(verticalCuts)
	max_h := max(horizontalCuts[0], h-horizontalCuts[m-1])
	max_v := max(verticalCuts[0], w-verticalCuts[n-1])
	for i := 1; i < m; i++ {
		max_h = max(max_h, horizontalCuts[i]-horizontalCuts[i-1])
	}
	for i := 1; i < n; i++ {
		max_v = max(max_v, verticalCuts[i]-verticalCuts[i-1])
	}
	return (max_h * max_v) % MOD
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
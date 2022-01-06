import "sort"

func maximumBeauty(items [][]int, queries []int) []int {
	// O(NlogN + QlogQ)
	queryPairs := make([][]int, len(queries))
	for i, query := range queries {
		queryPairs[i] = []int{query, i}
	}
	sort.Slice(queryPairs, func(i, j int) bool {
		return queryPairs[i][0] < queryPairs[j][0]
	})
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	idx, maxBeauty := 0, 0
	ans := make([]int, len(queries))
	for _, queryPair := range queryPairs {
		for idx < len(items) && items[idx][0] <= queryPair[0] {
			maxBeauty = max(maxBeauty, items[idx][1])
			idx++
		}
		ans[queryPair[1]] = maxBeauty
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
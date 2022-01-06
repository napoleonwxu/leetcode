import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	n := len(score)
	scoreAndIdx := make([][]int, n)
	for i := 0; i < n; i++ {
		scoreAndIdx[i] = []int{score[i], i}
	}
	sort.Slice(scoreAndIdx, func(i, j int) bool {
		return scoreAndIdx[i][0] > scoreAndIdx[j][0]
	})
	ans := make([]string, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			ans[scoreAndIdx[i][1]] = "Gold Medal"
		} else if i == 1 {
			ans[scoreAndIdx[i][1]] = "Silver Medal"
		} else if i == 2 {
			ans[scoreAndIdx[i][1]] = "Bronze Medal"
		} else {
			ans[scoreAndIdx[i][1]] = strconv.Itoa(i + 1)
		}
	}
	return ans
}
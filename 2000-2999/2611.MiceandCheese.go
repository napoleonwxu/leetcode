import "sort"

func miceAndCheese(reward1 []int, reward2 []int, k int) int {
	n, ans := len(reward1), 0
	diff := make([]int, n)
	for i := 0; i < n; i++ {
		ans += reward2[i]
		diff[i] = reward1[i] - reward2[i]
	}
	sort.Ints(diff)
	for ; k > 0; k-- {
		ans += diff[n-k]
	}
	return ans
}

func countTriplets(arr []int) int {
	// O(n^2)
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] ^ arr[i]
	}
	cnt := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if prefix[i] == prefix[j] {
				cnt += j - i - 1
			}
		}
	}
	return cnt
}
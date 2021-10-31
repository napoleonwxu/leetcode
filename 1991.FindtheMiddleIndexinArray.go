func findMiddleIndex(nums []int) int {
	// O(n) + O(1)
	sum, sumLeft := 0, 0
	for _, num := range nums {
		sum += num
	}
	for i, num := range nums {
		if sumLeft == sum-sumLeft-num {
			return i
		}
		sumLeft += num
	}
	/* O(n) + O(n)
	   n := len(nums)
	   sumLeft, sumRight := make([]int, n), make([]int, n)
	   for i := 1; i < n; i++ {
	       sumLeft[i] = sumLeft[i-1] + nums[i-1]
	       sumRight[n-i-1] = sumRight[n-i] + nums[n-i]
	   }
	   for i := 0; i < n; i++ {
	       if sumLeft[i] == sumRight[i] {
	           return i
	       }
	   }
	*/
	return -1
}
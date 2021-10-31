func countQuadruplets(nums []int) int {
	cnt, n := 0, len(nums)
	// O(n^3)
	m := make(map[int]int)
	m[nums[n-1]] = 1
	for c := n - 2; c >= 2; c-- {
		for b := c - 1; b >= 1; b-- {
			for a := b - 1; a >= 0; a-- {
				cnt += m[nums[a]+nums[b]+nums[c]]
			}
		}
		m[nums[c]]++
	}
	/* O(n^4)
	   for a := 0; a < n-3; a++ {
	       for b := a+1; b < n-2; b++ {
	           for c := b+1; c < n-1; c++ {
	               sum := nums[a] + nums[b] + nums[c]
	               for d := c+1; d < n; d++ {
	                   if nums[d] == sum {
	                       cnt++
	                   }
	               }
	           }
	       }
	   }
	*/
	return cnt
}
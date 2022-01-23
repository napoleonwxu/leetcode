import "sort"

func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		if len(nums[i]) == len(nums[j]) {
			for k := 0; k < len(nums[i]); k++ {
				if nums[i][k] != nums[j][k] {
					return nums[i][k] > nums[j][k]
				}
			}
		}
		return len(nums[i]) > len(nums[j])
	})
	return nums[k-1]
	/*
	   diffLenNums := make([][]string, 101)
	   for _, num := range nums {
	       n := len(num)
	       diffLenNums[n] = append(diffLenNums[n], num)
	   }
	   n := len(diffLenNums) - 1
	   for ; n > 0 && len(diffLenNums[n]) < k; n-- {
	       k -= len(diffLenNums[n])
	   }
	   leftNums := diffLenNums[n]

	   for i := 0; i < n; i++ {
	       idxNums := make([][]string, 10)
	       for _, num := range leftNums {
	           idx := int(num[i] - '0')
	           idxNums[idx] = append(idxNums[idx], num)
	       }
	       idx := len(idxNums) - 1
	       for ; idx > 0 && len(idxNums[idx]) < k; idx-- {
	           k -= len(idxNums[idx])
	       }
	       leftNums = idxNums[idx]
	   }
	   return leftNums[k-1]
	*/
}
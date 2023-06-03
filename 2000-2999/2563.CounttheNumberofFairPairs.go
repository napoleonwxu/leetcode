import "sort"

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	// O(nlgn) + O(2*n) => O(nlgn)
	return countLess(nums, upper) - countLess(nums, lower-1)
	/* O(nlgn) + O(2*n*lgn) => O(nlgn)
	   cnt := int64(0)
	   for i := 0; i < len(nums); i++ {
	       left := binSearchLeft(nums, lower-nums[i])
	       right := binSearchRight(nums, upper-nums[i])
	       if i >= left && i < right {
	           cnt += int64(right-left-1)
	       } else {
	           cnt += int64(right-left)
	       }
	   }
	   return cnt/2
	*/
}

func countLess(nums []int, target int) int64 {
	cnt := int64(0)
	for i, j := 0, len(nums)-1; i < j; i++ {
		for ; i < j && nums[i]+nums[j] > target; j-- {
		}
		cnt += int64(j - i)
	}
	return cnt
}

func binSearchLeft(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

func binSearchRight(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

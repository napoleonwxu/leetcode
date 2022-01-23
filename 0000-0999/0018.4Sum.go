func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    ans := [][]int{}
    kSum(nums, []int{}, target, 4, 0, &ans)
    return ans
}

func kSum(nums, path []int, target, k, lo int, ans *[][]int) {
    // O(n^(k-1))
    n := len(nums)
    if n-lo < k || k < 2 || target < k*nums[lo] || target > k*nums[n-1] {
        return
    }
    if k == 2 {
        hi := n - 1
        for lo < hi {
            if nums[lo]+nums[hi] == target {
                *ans = append(*ans, append(path, nums[lo], nums[hi]))
                for ; lo < hi && nums[lo] == nums[lo+1]; lo++ {}
                for ; lo < hi && nums[hi] == nums[hi-1]; hi-- {}
                lo, hi = lo+1, hi-1
            } else if nums[lo]+nums[hi] < target {
                lo++
            } else {
                hi--
            }
        }
    } else {
        for i := lo; i <= n-k; i++ {
            if i == lo || (i > lo && nums[i] != nums[i-1]) {
                kSum(nums, append(path, nums[i]), target-nums[i], k-1, i+1, ans)
            }
        }
    }
}
func threeSum(nums []int) [][]int {
    n := len(nums)
    sort.Ints(nums)
    ans := [][]int{}
    for i := 0; i < n-2; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        lo, hi := i+1, n-1
        for lo < hi {
            if nums[lo]+nums[hi] == -nums[i] {
                ans = append(ans, []int{nums[i], nums[lo], nums[hi]})
                for ; lo < hi && nums[lo] == nums[lo+1]; lo++ {}
                for ; lo < hi && nums[hi] == nums[hi-1]; hi-- {}
                lo++
                hi--
            } else if nums[lo]+nums[hi] < -nums[i] {
                lo++
            } else {
                hi--
            }
        }
    }
    return ans
}
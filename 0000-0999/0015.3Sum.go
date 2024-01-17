func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    ans := [][]int{}
    n := len(nums)
    for i := 0; i < n-2; i++ {
        if nums[i] > 0 {
            break
        }
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i+1, n-1
        for j < k {
            if nums[i]+nums[j]+nums[k] == 0 {
                ans = append(ans, []int{nums[i], nums[j], nums[k]})
                for ; j < k && nums[j] == nums[j+1]; j++ {
                }
                for ; j < k && nums[k] == nums[k-1]; k-- {
                }
                j, k = j+1, k-1
            } else if nums[i]+nums[j]+nums[k] < 0 {
                j++
            } else {
                k--
            }
        }
    }
    return ans
}
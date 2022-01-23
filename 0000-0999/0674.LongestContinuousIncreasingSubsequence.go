func findLengthOfLCIS(nums []int) int {
    ans, tmp := 0, 0
    for i := 0; i < len(nums); i++ {
        if i == 0 || nums[i] > nums[i-1] {
            tmp++
            if tmp > ans {
                ans = tmp
            }
        } else {
            tmp = 1
        }
    }
    return ans
}
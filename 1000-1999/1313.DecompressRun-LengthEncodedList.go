func decompressRLElist(nums []int) []int {
    ans := []int{}
    for i := 0; i < len(nums); i += 2 {
        for j := 0; j < nums[i]; j++ {
            ans = append(ans, nums[i+1])
        }
    }
    return ans
}
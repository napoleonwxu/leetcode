func minDeletion(nums []int) int {
    del, n := 0, len(nums)
    for i := 0; i < n-1; i++ {
        if nums[i] == nums[i+1] && (i-del)%2 == 0 {
            del++
        }
    }
    return del + (n-del)%2
}